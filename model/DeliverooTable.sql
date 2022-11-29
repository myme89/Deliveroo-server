CREATE TABLE IF NOT EXISTS "group_type" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "tittle" varchar NOT NULL,
  "description" varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS "restaurant" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "id_group_type" int NOT NULL,
  "tittle" varchar NOT NULL,
  "rating" int NOT NULL,
  "genre" varchar NOT NULL,
  "address" varchar NOT NULL,
  "short_description" varchar NOT NULL,
  "long" float8 NOT NULL,
  "lat" float8 NOT NULL
);

CREATE TABLE IF NOT EXISTS "menu_restaurant" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "id_restaurant" int NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "price" int NOT NULL,
  "image" varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS "group_users" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "group_name" varchar UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS "users" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "user_name" varchar UNIQUE NOT NULL,
  "id_group_users" int NOT NULL,
  "password" varchar NOT NULL
);

