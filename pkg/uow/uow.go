package uow

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type RepositoryFactory func(tx *sql.Tx) any

type UowInterface interface {
	Register(name string, fc RepositoryFactory)
	GetRepository(ctx context.Context, name string) (any, error)
	Do(ctx context.Context, fn func(uow *Uow) error) error
	CommitOrRollback() error
	Rollback() error
	UnRegister(name string)
}

type Uow struct {
	db           *sql.DB
	tx           *sql.Tx
	repositories map[string]RepositoryFactory
}

func New(ctx context.Context, db *sql.DB) *Uow {
	return &Uow{
		db:           db,
		repositories: make(map[string]RepositoryFactory),
	}
}

func (u *Uow) Register(name string, fc RepositoryFactory) {
	u.repositories[name] = fc
}

func (u *Uow) UnRegister(name string) {
	delete(u.repositories, name)
}

func (u *Uow) GetRepository(ctx context.Context, name string) (any, error) {
	if u.tx == nil {
		tx, err := u.db.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
		u.tx = tx
	}
	repo := u.repositories[name](u.tx)
	return repo, nil
}

func (u *Uow) Do(ctx context.Context, fn func(uow *Uow) error) error {
	if u.tx != nil {
		return fmt.Errorf("transaction already started")
	}

	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	u.tx = tx

	if err := fn(u); err != nil {
		errRb := u.Rollback()
		if errRb != nil {
			return fmt.Errorf("original error: %s, rollback error: %s", err.Error(), errRb.Error())
		}
		return err
	}

	return u.CommitOrRollback()
}

func (u *Uow) Rollback() error {
	if u.tx == nil {
		return errors.New("no transaction to rollback")
	}

	if err := u.tx.Rollback(); err != nil {
		return err
	}
	u.tx = nil

	return nil
}

func (u *Uow) CommitOrRollback() error {
	if err := u.tx.Commit(); err != nil {
		errRb := u.Rollback()
		if errRb != nil {
			return fmt.Errorf("original error: %s, rollback error: %s", err.Error(), errRb.Error())
		}
		return err
	}
	u.tx = nil

	return nil
}
