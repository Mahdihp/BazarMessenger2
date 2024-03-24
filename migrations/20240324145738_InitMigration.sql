-- Create "profile_and_settings" table
CREATE TABLE "profile_and_settings" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, PRIMARY KEY ("id"));
-- Create "stores" table
CREATE TABLE "stores" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, PRIMARY KEY ("id"));
-- Create "roles" table
CREATE TABLE "roles" ("id" smallint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "name" character varying NOT NULL, "description" character varying NULL, PRIMARY KEY ("id"));
-- Create index "role_id" to table: "roles"
CREATE UNIQUE INDEX "role_id" ON "roles" ("id");
-- Create index "roles_name_key" to table: "roles"
CREATE UNIQUE INDEX "roles_name_key" ON "roles" ("name");
-- Create "users" table
CREATE TABLE "users" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "mobile" character varying NOT NULL, "username" character varying NULL, "active" boolean NOT NULL DEFAULT true, "deleted" boolean NOT NULL DEFAULT false, "created_at" timestamp NOT NULL, "updated_at" timestamp NOT NULL, PRIMARY KEY ("id"));
-- Create index "user_id" to table: "users"
CREATE UNIQUE INDEX "user_id" ON "users" ("id");
-- Create index "users_mobile_key" to table: "users"
CREATE UNIQUE INDEX "users_mobile_key" ON "users" ("mobile");
-- Create "role_users" table
CREATE TABLE "role_users" ("role_id" smallint NOT NULL, "user_id" bigint NOT NULL, PRIMARY KEY ("role_id", "user_id"), CONSTRAINT "role_users_role_id" FOREIGN KEY ("role_id") REFERENCES "roles" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "role_users_user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);