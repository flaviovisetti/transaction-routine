BEGIN;

CREATE TABLE IF NOT EXISTS transactions (
  id serial PRIMARY KEY,
  account_id INTEGER REFERENCES accounts(id) NOT NULL,
  operation_type_id INTEGER REFERENCES operation_types(id) NOT NULL,
  amount DECIMAL NOT NULL,
  event_date TIMESTAMP NOT NULL
);

COMMIT;