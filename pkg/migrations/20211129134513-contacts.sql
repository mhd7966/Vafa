-- +migrate Up
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR,
  first_name VARCHAR,
  last_name VARCHAR,
  status INT,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP
);
-- +migrate Down
DROP TABLE users;