-- +migrate Up
CREATE TABLE IF NOT EXISTS document (
    id BIGSERIAL PRIMARY KEY,
    journal_id INTEGER NOT NULL,
    content TEXT NOT NULL, 
    embedding vector(768) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE INDEX index_document_created_at ON document(created_at);
CREATE INDEX index_document_updated_at ON document(updated_at);

-- +migrate Down
DROP TABLE IF EXISTS document;

