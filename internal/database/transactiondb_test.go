package database

import (
	"database/sql"
	"testing"

	"github.com/marcioecom/wallet-core/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)

	s.db = db
	db.Exec("CREATE TABLE clients (id VARCHAR(255), name VARCHAR(255), email VARCHAR(255), created_at DATE, updated_at DATE);")
	db.Exec("CREATE TABLE accounts (id VARCHAR(255), client_id VARCHAR(255), balance INT, created_at DATE, updated_at DATE);")
	db.Exec("CREATE TABLE transactions (id VARCHAR(255), account_from_id VARCHAR(255), account_to_id VARCHAR(255), amount INT, created_at DATE);")

	client1, err := entity.NewClient("Client 1", "client1@mail.com")
	s.Nil(err)
	client2, err := entity.NewClient("Client 2", "client2@mail.com")
	s.Nil(err)
	s.accountFrom, err = entity.NewAccount(client1)
	s.accountFrom.Credit(1000)
	s.Nil(err)
	s.accountTo, err = entity.NewAccount(client2)
	s.accountFrom.Credit(1000)
	s.Nil(err)

	s.transactionDB = NewTransactionDB(db)
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients;")
	s.db.Exec("DROP TABLE accounts;")
	s.db.Exec("DROP TABLE transactions;")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)

	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}
