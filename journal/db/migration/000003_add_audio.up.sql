CREATE TABLE "audio" (
    "id" bigserial PRIMARY KEY,
    "journal_id" bigserial NOT NULL,
    "content" bytea NOT NULL ,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "is_embedded" boolean NOT NULL DEFAULT false
);

ALTER TABLE "audio" ADD FOREIGN KEY ("journal_id") REFERENCES "journal" ("id");