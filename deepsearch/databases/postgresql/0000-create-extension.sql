-- +migrate Up
CREATE EXTENSION IF NOT EXISTS vector;

-- +migrate Down
DROP EXTENSION vector CASCADE;
