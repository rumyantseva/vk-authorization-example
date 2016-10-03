CREATE TABLE auth (
    id serial PRIMARY KEY,
    email character varying NOT NULL UNIQUE,
    password character varying NOT NULL,

    vk_user_id bigint NOT NULL UNIQUE,
    vk_token character varying NOT NULL,
    vk_token_lifetime int NOT NULL,

    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);
