-- Modify "users" table
ALTER TABLE "public"."users" ADD COLUMN "firebase_uid" text NOT NULL, ADD CONSTRAINT "users_firebase_uid_key" UNIQUE ("firebase_uid");
