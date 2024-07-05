USE wallet;

CREATE TABLE clients (
  id VARCHAR(255),
  name VARCHAR(255),
  email VARCHAR(255),
  created_at DATE,
  updated_at DATE
);

CREATE TABLE accounts (
  id VARCHAR(255),
  client_id VARCHAR(255),
  balance INT,
  created_at DATE,
  updated_at DATE
);

CREATE TABLE transactions (
  id VARCHAR(255),
  account_from_id VARCHAR(255),
  account_to_id VARCHAR(255),
  amount INT,
  created_at DATE
);
