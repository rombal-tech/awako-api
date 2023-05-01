CREATE TABLE "public.Account" (
  "email" character varying NOT NULL,
  "password" character varying NOT NULL,
  "deleted" BOOLEAN NOT NULL,
  CONSTRAINT "Account_pk" PRIMARY KEY ("email")
) WITH (
  OIDS=FALSE
  );

CREATE TABLE "public.ConfirmedAccount" (
       "email" character varying NOT NULL,
       "code" character varying NOT NULL,
       "confirmed" BOOLEAN NOT NULL,
       CONSTRAINT "ConfirmedAccount_pk" PRIMARY KEY ("email")
) WITH (
  OIDS=FALSE
  );

CREATE TABLE "public.Session" (
      "id" bigserial NOT NULL,
      "email" character varying NOT NULL,
      "session_string" character varying NOT NULL,

      CONSTRAINT "Session_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
  );

CREATE TABLE "public.Scheme"(
    "id" bigserial NOT NULL,
    "name" character varying NOT NULL,
    "description" character varying NOT NULL,
    "author" character varying NOT NULL,
    "creation_date" timestamp NOT NULL,
    CONSTRAINT "Scheme_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
  );

ALTER TABLE "public.ConfirmedAccount" ADD CONSTRAINT "ConfirmedAccount_fk0" FOREIGN KEY ("email") REFERENCES "public.Account"("email");
ALTER TABLE "public.Session" ADD CONSTRAINT "Session_fk0" FOREIGN KEY ("email") REFERENCES "public.Account"("email");
ALTER TABLE "public.Scheme" ADD CONSTRAINT "Scheme_fk0" FOREIGN KEY ("author") REFERENCES "public.Account"("email");
