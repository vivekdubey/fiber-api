CREATE TABLE "books" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "author" varchar NOT NULL,
  "description" varchar NOT NULL,
  "count" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "book_id" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "count" bigint
);

COMMENT ON COLUMN "entries"."count" IS 'negative for borrow and positive for return';

ALTER TABLE "entries" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");