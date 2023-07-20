CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);