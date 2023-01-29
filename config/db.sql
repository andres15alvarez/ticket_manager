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
  "help_topic_id" int,
  "department_id" int,
  "state_id" int,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

ALTER TABLE "user" ADD FOREIGN KEY ("user_type_id") REFERENCES "user_type" ("id");

ALTER TABLE "ticket" ADD FOREIGN KEY ("created_by_id") REFERENCES "user" ("id");

ALTER TABLE "ticket" ADD FOREIGN KEY ("help_topic_id") REFERENCES "help_topic" ("id");

ALTER TABLE "ticket" ADD FOREIGN KEY ("department_id") REFERENCES "department" ("id");

ALTER TABLE "ticket" ADD FOREIGN KEY ("state_id") REFERENCES "state" ("id");

INSERT INTO public.state(name)
VALUES ('abierto'),
	   ('cerrado');

INSERT INTO public.help_topic(name)
VALUES ('devolucion'),
	   ('garantia'),
	   ('ayuda');

INSERT INTO public.department(name)
	VALUES ('Garantias'),
			('Finanzas');

INSERT INTO public.user_type(name)
	VALUES ('ADMIN'),
		   ('OPERATIVO'),
		   ('CLIENTE');

INSERT INTO public."user"(
	username, password, first_name, last_name, email, user_type_id, created_at, updated_at, deleted_at)
	VALUES ('paco', 'paco', 'paco', 'juarez', 'paco@gmail.com', 3, current_timestamp, current_timestamp, NULL);

INSERT INTO public.ticket(customer_name, customer_email, customer_phone_number, invoice_number, waranty_number, detail, invoice_image, waranty_certificate, subject, created_by_id, help_topic_id, department_id, state_id, created_at, updated_at, deleted_at)
	VALUES ('Kevin', 'chengkev2000@gmail.com', '4126796098', '1273', '1923', 'producto roto', '123', '123', 'producto roto', 1, 2, 1, 2, current_timestamp,current_timestamp,NULL),
	('Gabriel', 'gabriel@gmail.com', '4126683599', '1253', '1253', 'quiero mi dinero de vuelta', '123', '123', 'quiero mi dinero de vuelta', 1, 1, 1, 1, current_timestamp,current_timestamp,NULL),
	('Juan', 'juan@gmail.com', '4145568796', '1243', '1233', 'mi sobrino se espanto', '123', '123', 'mi sobrino se espanto', 1, 1, 1, 2, current_timestamp,current_timestamp,NULL),
	('Roman', 'roman@gmail.com', '4146792233', '1123', '1223', 'me dejaron paraplejico', '123', '123', 'me dejaron paraplejico', 1, 3, 1, 1, current_timestamp,current_timestamp,NULL);

  
