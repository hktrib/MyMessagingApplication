-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2023-10-01T21:07:33.390Z

CREATE TABLE "users" (
  "username" varchar,
  "hashed_password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "is_email_verified" bool NOT NULL DEFAULT false,
  PRIMARY KEY ("username")
);

CREATE TABLE "verify_emails" (
  "id" bigserial,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "secret_code" varchar NOT NULL,
  "is_used" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "expired_at" timestamptz NOT NULL DEFAULT (now() + interval '5 minutes'),
  PRIMARY KEY ("id")
);

CREATE INDEX ON "verify_emails" ("username");

CREATE INDEX ON "verify_emails" ("email");

ALTER TABLE "verify_emails" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
