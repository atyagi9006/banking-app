CREATE TABLE "bank_employee" (
  "id" bigserial PRIMARY KEY,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "role" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "bank_employee" ("email");