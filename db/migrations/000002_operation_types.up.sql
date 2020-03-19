BEGIN;

CREATE TABLE IF NOT EXISTS operation_types (
  id serial PRIMARY KEY,
  "description" VARCHAR (255) NOT NULL
);

COMMIT;