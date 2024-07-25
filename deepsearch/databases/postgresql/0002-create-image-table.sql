-- +migrate Up
CREATE TABLE IF NOT EXISTS image (
    id BIGSERIAL PRIMARY KEY,
    image_id BIGSERIAL NOT NULL,
    journal_id BIGSERIAL NOT NULL,
    username VARCHAR(255) NOT NULL,
    embedding vector(768) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE INDEX index_image_created_at ON image(created_at);
CREATE INDEX index_image_updated_at ON image(updated_at);

-- +migrate Down
DROP TABLE IF EXISTS image;