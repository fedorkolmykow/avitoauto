\connect postgres

CREATE DATABASE avitoauto;

\connect avitoauto

CREATE TABLE URLs (
	url_id serial NOT NULL,
	url VARCHAR(255) NOT NULL UNIQUE,
	key VARCHAR(255) NOT NULL UNIQUE,
	CONSTRAINT Url_pk PRIMARY KEY (url_id)
) WITH (
  OIDS=FALSE
);

