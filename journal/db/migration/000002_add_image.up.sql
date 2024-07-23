CREATE TABLE "image" (
    "id" bigserial PRIMARY KEY,
    "journal_id" bigserial NOT NULL,
    "content" bytea NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "is_embedded" boolean NOT NULL DEFAULT false
);

ALTER TABLE "image" ADD FOREIGN KEY ("journal_id") REFERENCES "journal" ("id");