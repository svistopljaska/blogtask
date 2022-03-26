CREATE DATABASE blog
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1;

-- Table: public.posts

-- DROP TABLE IF EXISTS public.posts;

CREATE TABLE IF NOT EXISTS public.posts
(
    id integer NOT NULL DEFAULT nextval('posts_id_seq'::regclass),
    title character varying(255) COLLATE pg_catalog."default" NOT NULL,
    content text COLLATE pg_catalog."default" NOT NULL,
    author character varying(255) COLLATE pg_catalog."default" NOT NULL,
    createdon timestamp with time zone NOT NULL,
    CONSTRAINT posts_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.posts
    OWNER to postgres;


-- Table: public.tags

-- DROP TABLE IF EXISTS public.tags;

CREATE TABLE IF NOT EXISTS public.tags
(
    id integer NOT NULL DEFAULT nextval('tags_id_seq'::regclass),
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT tags_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.tags
    OWNER to postgres;


-- Table: public.i_post_tags

-- DROP TABLE IF EXISTS public.i_post_tags;


CREATE TABLE IF NOT EXISTS public.i_post_tags
(
    id integer NOT NULL DEFAULT nextval('i_post_tags_id_seq'::regclass),
    id_post integer NOT NULL,
    id_tag integer NOT NULL,
    CONSTRAINT i_post_tags_pkey PRIMARY KEY (id),
    CONSTRAINT posts_fk FOREIGN KEY (id_post)
        REFERENCES public.posts (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT tags_fk FOREIGN KEY (id_tag)
        REFERENCES public.tags (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.i_post_tags
    OWNER to postgres;