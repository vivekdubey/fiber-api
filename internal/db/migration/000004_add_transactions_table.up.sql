CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "book_id" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "count" bigint
);

COMMENT ON COLUMN "transactions"."count" IS 'negative for borrow and positive for return';

ALTER TABLE "transactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");