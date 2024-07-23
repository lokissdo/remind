CREATE TABLE "journal" (
    "id" bigserial PRIMARY KEY,
    "username" varchar NOT NULL,
    "title" varchar NOT NULL,
    "content" text,
    "status" bool NOT NULL DEFAULT false,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    "is_embedded" bool NOT NULL DEFAULT false
)