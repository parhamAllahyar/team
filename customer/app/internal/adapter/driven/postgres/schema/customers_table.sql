CREATE TABLE
IF NOT EXISTS tracks
(
    id uuid NOT NULL,
    first_name character varying (40) COLLATE pg_catalog."default",
    last_name character varying (40) COLLATE pg_catalog."default",
    email character varying (40) COLLATE pg_catalog."default",
    phone character varying (40) COLLATE pg_catalog."default",
    status character varying (40) COLLATE pg_catalog."default",
    password character varying (40) COLLATE pg_catalog."default",
    -- created_at TIME  (40) COLLATE pg_catalog."default",
    -- deleted_at TIME  (40) COLLATE pg_catalog."default",
    -- updated_at TIME  (40) COLLATE pg_catalog."default",
)
WITH
(
    OIDS = FALSE
)
TABLESPACE pg_default;