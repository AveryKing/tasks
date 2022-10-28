CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "username" varchar UNIQUE,
                         "password" varchar NOT NULL
);

CREATE TABLE "tasks" (
                         "id" bigserial PRIMARY KEY,
                         "user" bigserial NOT NULL,
                         "title" varchar NOT NULL,
                         "priority" int NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "tasks" ADD FOREIGN KEY ("user") REFERENCES "users" ("id");
