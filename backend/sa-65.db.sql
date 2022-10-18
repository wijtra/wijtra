BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "disease_types" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"disease_type"	text,
	"disease_type_name"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "diseases" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"disease_name"	text,
	"disease_type_id"	integer,
	PRIMARY KEY("id"),
	CONSTRAINT "fk_disease_types_disease" FOREIGN KEY("disease_type_id") REFERENCES "disease_types"("id")
);
CREATE TABLE IF NOT EXISTS "inpantient_departments" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"inpantient_department_name"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "patients" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"patient_name"	text,
	"gender"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "triages" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"patient_id"	integer,
	"disease_id"	integer,
	"inpantient_department_id"	integer,
	"triage_comment"	text,
	"user_id"	integer,
	PRIMARY KEY("id"),
	CONSTRAINT "fk_inpantient_departments_triages" FOREIGN KEY("inpantient_department_id") REFERENCES "inpantient_departments"("id"),
	CONSTRAINT "fk_users_triages" FOREIGN KEY("user_id") REFERENCES "users"("id"),
	CONSTRAINT "fk_patients_triages" FOREIGN KEY("patient_id") REFERENCES "patients"("id"),
	CONSTRAINT "fk_diseases_triages" FOREIGN KEY("disease_id") REFERENCES "diseases"("id")
);
CREATE TABLE IF NOT EXISTS "user_types" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"user_type"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "users" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"user_name"	text,
	"user_password"	text,
	"isname"	text,
	"user_type_id"	integer,
	PRIMARY KEY("id"),
	CONSTRAINT "fk_user_types_users" FOREIGN KEY("user_type_id") REFERENCES "user_types"("id")
);
CREATE INDEX IF NOT EXISTS "idx_disease_types_deleted_at" ON "disease_types" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_diseases_deleted_at" ON "diseases" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_inpantient_departments_deleted_at" ON "inpantient_departments" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_patients_deleted_at" ON "patients" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_triages_deleted_at" ON "triages" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_user_types_deleted_at" ON "user_types" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_users_deleted_at" ON "users" (
	"deleted_at"
);
COMMIT;
