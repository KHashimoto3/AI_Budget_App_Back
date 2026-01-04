-- Create "users" table
CREATE TABLE "public"."users" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" character varying NOT NULL,
  "disp_name" character varying NOT NULL,
  "email" character varying NOT NULL,
  "password_hash" text NOT NULL,
  "account_type" integer NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "login_at" timestamp NULL,
  "deleted_flag" boolean NOT NULL DEFAULT false,
  PRIMARY KEY ("id"),
  CONSTRAINT "users_email_key" UNIQUE ("email"),
  CONSTRAINT "users_account_type_check" CHECK (account_type = ANY (ARRAY[0, 1]))
);
-- Create index "idx_users_email" to table: "users"
CREATE INDEX "idx_users_email" ON "public"."users" ("email");
-- Create "genres" table
CREATE TABLE "public"."genres" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "name" character varying NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "genres_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_genres_user_id" to table: "genres"
CREATE INDEX "idx_genres_user_id" ON "public"."genres" ("user_id");
-- Create "images" table
CREATE TABLE "public"."images" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "filename" character varying NOT NULL,
  "file_path" character varying NOT NULL,
  "parsed_at" timestamp NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "images_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_images_user_id" to table: "images"
CREATE INDEX "idx_images_user_id" ON "public"."images" ("user_id");
-- Create "expenses" table
CREATE TABLE "public"."expenses" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "expense_date" date NOT NULL,
  "amount" integer NOT NULL,
  "genres_id" uuid NULL,
  "shop_name" character varying NOT NULL,
  "memo" text NULL,
  "input_type" character varying NOT NULL,
  "image_id" uuid NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "expenses_genres_id_fkey" FOREIGN KEY ("genres_id") REFERENCES "public"."genres" ("id") ON UPDATE NO ACTION ON DELETE SET NULL,
  CONSTRAINT "expenses_image_id_fkey" FOREIGN KEY ("image_id") REFERENCES "public"."images" ("id") ON UPDATE NO ACTION ON DELETE SET NULL,
  CONSTRAINT "expenses_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "expenses_amount_check" CHECK (amount >= 0),
  CONSTRAINT "expenses_input_type_check" CHECK ((input_type)::text = ANY ((ARRAY['manual'::character varying, 'ocr'::character varying])::text[]))
);
-- Create index "idx_expenses_expense_date" to table: "expenses"
CREATE INDEX "idx_expenses_expense_date" ON "public"."expenses" ("expense_date");
-- Create index "idx_expenses_genres_id" to table: "expenses"
CREATE INDEX "idx_expenses_genres_id" ON "public"."expenses" ("genres_id");
-- Create index "idx_expenses_user_id" to table: "expenses"
CREATE INDEX "idx_expenses_user_id" ON "public"."expenses" ("user_id");
