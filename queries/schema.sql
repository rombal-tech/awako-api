CREATE TABLE "Account" (
              "email" character varying NOT NULL,
              "password" character varying NOT NULL,
              "deleted" BOOLEAN NOT NULL,
              CONSTRAINT "Account_pk" PRIMARY KEY ("email")
) WITH (
OIDS=FALSE
);



CREATE TABLE "ConfirmedAccount" (
            "email" character varying NOT NULL,
            "code" character varying NOT NULL,
            "confirmed" BOOLEAN NOT NULL,
            CONSTRAINT "ConfirmedAccount_pk" PRIMARY KEY ("email")
) WITH (
OIDS=FALSE
);




ALTER TABLE "ConfirmedAccount" ADD CONSTRAINT "ConfirmedAccount_fk0" FOREIGN KEY ("email") REFERENCES "Account"("email");