-- Drop table

-- DROP TABLE public.users;

CREATE TABLE users (
	id serial NOT NULL,
	username text NOT NULL,
	email text NOT NULL,
	"password" text NOT NULL,
	islogin bool NOT NULL DEFAULT false,
	CONSTRAINT afc_sso_users_pkey PRIMARY KEY (id),
	CONSTRAINT afc_sso_users_user_code_key UNIQUE (email),
	CONSTRAINT users_un UNIQUE (username)
);
