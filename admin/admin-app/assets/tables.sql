

CREATE TYPE ADMIN_STATUS AS ENUM ('active', 'deactive');
CREATE TYPE OTP_STATUS AS ENUM ('active', 'deactive');
CREATE TYPE OTP_TYPE AS ENUM ('update_password');
CREATE TYPE LINK_STATUS AS ENUM ('active', 'deactive');
CREATE TYPE LINK_TYPE AS ENUM ('invite_link',);



CREATE TABLE IF NOT EXISTS admins (
	id UUID UNIQUE NOT NULL,
	first_name VARCHAR NOT NULL,
	last_name VARCHAR NOT NULL,
	phone VARCHAR NOT NULL,
	email VARCHAR NOT NULL,
	password VARCHAR NOT NULL,
	status ADMIN_STATUS NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP DEFAULT NULL,
	deleted_at TIMESTAMP DEFAULT NULL,
	CONSTRAINT admins_pkey PRIMARY KEY (id)
);


CREATE TABLE IF NOT EXISTS links (
	id UUID UNIQUE NOT NULL,
    type LINK_TYPE NOT NULL,
	email VARCHAR NOT NULL,
	link VARCHAR NOT NULL,
	status LINK_TYPE  NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP DEFAULT NULL,
	expired_at TIMESTAMP DEFAULT NULL,
	CONSTRAINT links_pkey PRIMARY KEY (id)
);


CREATE TABLE IF NOT EXISTS otps (
	id UUID UNIQUE NOT NULL,
	type OTP_TYPE NOT NULL,
	new_phone VARCHAR NOT NULL,
	otp INT NOT NULL,
	status OTP_STATUS NOT NULL,
	created_at TIMESTAMP NOT NULL,
	expired_at TIMESTAMP DEFAULT NULL,
	updated_at TIMESTAMP DEFAULT NULL,	
);


ALTER TABLE links action

ALTER TABLE otps action






