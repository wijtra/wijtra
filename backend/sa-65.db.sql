BEGIN TRANSACTION;
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
	CONSTRAINT "fk_users_user_type" FOREIGN KEY("user_type_id") REFERENCES "user_types"("id"),
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "patients" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"patient_name"	text,
	"gender_id"	integer,
	"gender"	text,
	"blood_type_id"	integer,
	"blood_type"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "disease_types" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
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
	CONSTRAINT "fk_diseases_disease_type" FOREIGN KEY("disease_type_id") REFERENCES "disease_types"("id"),
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "inpantient_departments" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"inpantient_department_name"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "triage_states" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"triage_state"	text,
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
	"triage_state_id"	integer,
	"triage_comment"	text,
	"user_id"	integer,
	CONSTRAINT "fk_triages_disease" FOREIGN KEY("disease_id") REFERENCES "diseases"("id"),
	CONSTRAINT "fk_triages_user" FOREIGN KEY("user_id") REFERENCES "users"("id"),
	CONSTRAINT "fk_triages_triage_state" FOREIGN KEY("triage_state_id") REFERENCES "triage_states"("id"),
	CONSTRAINT "fk_triages_patient" FOREIGN KEY("patient_id") REFERENCES "patients"("id"),
	CONSTRAINT "fk_triages_inpantient_department" FOREIGN KEY("inpantient_department_id") REFERENCES "inpantient_departments"("id"),
	PRIMARY KEY("id")
);
CREATE INDEX IF NOT EXISTS "idx_user_types_deleted_at" ON "user_types" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_users_deleted_at" ON "users" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_patients_deleted_at" ON "patients" (
	"deleted_at"
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
CREATE INDEX IF NOT EXISTS "idx_triage_states_deleted_at" ON "triage_states" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_triages_deleted_at" ON "triages" (
	"deleted_at"
);
COMMIT;
