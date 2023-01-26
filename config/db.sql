CREATE TABLE "user_type" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL
);

CREATE TABLE "user" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "username" varchar(255) NOT NULL UNIQUE,
  "password" varchar(255) NOT NULL,
  "first_name" varchar(100),
  "last_name" varchar(100),
  "email" varchar(100) NOT NULL UNIQUE,
  "user_type_id" int,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "help_topic" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL
);

CREATE TABLE "department" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL
);

CREATE TABLE "state" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(255) NOT NULL
);

CREATE TABLE "ticket" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "customer_name" varchar(255) NOT NULL,
  "customer_email" varchar(100) NOT NULL,
  "customer_phone_number" varchar(21) NOT NULL,
  "invoice_number" varchar(255) NOT NULL,
  "waranty_number" varchar(255) NOT NULL,
  "detail" text,
  "invoice_image" text,
  "waranty_certificate" text,
  "subject" varchar(255),
  "created_by_id" int,
  "taken_by_id" int,
  "taken_at" timestamp,
  "help_topic_id" int,
  "department_id" int,
  "state_id" int,
  "answer" text,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

ALTER TABLE "user" ADD FOREIGN KEY ("user_type_id") REFERENCES "user_type" ("id");

ALTER TABLE "ticket" ADD FOREIGN KEY ("created_by_id") REFERENCES "user" ("id");

ALTER TABLE "ticket" ADD FOREIGN KEY ("taken_by_id") REFERENCES "user" ("id");

ALTER TABLE "ticket" ADD FOREIGN KEY ("help_topic_id") REFERENCES "help_topic" ("id");

ALTER TABLE "ticket" ADD FOREIGN KEY ("department_id") REFERENCES "department" ("id");

ALTER TABLE "ticket" ADD FOREIGN KEY ("state_id") REFERENCES "state" ("id");
