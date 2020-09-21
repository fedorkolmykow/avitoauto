\connect postgres

CREATE DATABASE avitoauto;

\connect avitoauto

CREATE TABLE Users (
	url_id serial NOT NULL,
	url VARCHAR(255) NOT NULL,
	key VARCHAR(255) NOT NULL,
	CONSTRAINT Url_pk PRIMARY KEY (url_id)
) WITH (
  OIDS=FALSE
);

