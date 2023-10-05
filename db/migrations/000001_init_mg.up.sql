-- Agregando la extensión UUID
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "Users"
(
    id                             uuid    DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    email                          VARCHAR(255),
    accept_terms_and_conditions    BOOLEAN DEFAULT FALSE              NOT NULL,
    hash                           VARCHAR(255),
    first_name                     VARCHAR(255),
    last_name                      VARCHAR(255),
    created_at                     TIMESTAMP WITH TIME ZONE           NOT NULL,
    updated_at                     TIMESTAMP WITH TIME ZONE           NOT NULL,
    accept_terms_and_conditions_at TIMESTAMP WITH TIME ZONE           NOT NULL,
    photo_url                      VARCHAR(255),
    timezone                       VARCHAR(255),
    phone                          VARCHAR(255),
    email_verified_at              TIMESTAMP WITH TIME ZONE,
    is_locked                      BOOLEAN DEFAULT FALSE              NOT NULL,
    bad_attempts                   INTEGER DEFAULT 0                  NOT NULL,
    last_login                     TIMESTAMP WITH TIME ZONE,
    deleted_at                     TIMESTAMP WITH TIME ZONE
);

CREATE UNIQUE INDEX email_unique_idx
    ON "Users" (LOWER(email::TEXT));

CREATE INDEX idx_users_id
    ON "Users" (id);

CREATE UNIQUE INDEX users_email
    ON "Users" (email);
