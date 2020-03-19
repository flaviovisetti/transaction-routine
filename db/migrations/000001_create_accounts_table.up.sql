BEGIN;

CREATE TABLE IF NOT EXISTS accounts (
  id serial PRIMARY KEY,
  document_number VARCHAR (255) NOT NULL
);

COMMIT;