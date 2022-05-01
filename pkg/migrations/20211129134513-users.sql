-- +migrate Up
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  national_id VARCHAR UNIQUE,
  first_name VARCHAR,
  last_name VARCHAR,
  status INT DEFAULT 1,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  deleted_at TIMESTAMP
);
-- +migrate Down
DROP TABLE users;