CREATE TABLE "Account" (
  "email" character varying NOT NULL,
  "password" character varying NOT NULL,
  "deleted" BOOLEAN NOT NULL,
  "confirmed" BOOLEAN NOT NULL,
  "confirm_code" character varying NOT NULL,
  CONSTRAINT "Account_pk" PRIMARY KEY ("email")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "Session" (
      "session" character varying NOT NULL,
      "email" character varying NOT NULL,
      CONSTRAINT "Session_pk" PRIMARY KEY ("session")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "Scheme"(
    "id" bigserial NOT NULL,
    "name" character varying NOT NULL,
    "description" character varying NOT NULL,
    "author" character varying NOT NULL,
    "creation_date" timestamp NOT NULL,
    CONSTRAINT "Scheme_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

ALTER TABLE "Session" ADD CONSTRAINT "Session_fk0" FOREIGN KEY ("email") REFERENCES "Account"("email");
ALTER TABLE "Scheme" ADD CONSTRAINT "Scheme_fk0" FOREIGN KEY ("author") REFERENCES "Account"("email");
