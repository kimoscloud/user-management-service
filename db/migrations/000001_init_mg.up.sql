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


INSERT INTO "Users" (id, email, accept_terms_and_conditions, hash, first_name, last_name, created_at, updated_at,
                            accept_terms_and_conditions_at, photo_url, timezone, phone, email_verified_at, is_locked,
                            bad_attempts, last_login, deleted_at)
VALUES ('a6c56de9-1fc0-4215-99a0-97753bb712d8', 'seebogado@gmail.com', true,
        '$2a$10$NSSisPP6kwFV023oAxlBs.VIB8kTQYciSupvaWykoBiHB1Qd2aJva', 'Sebastian', 'Bogado',
        '2024-01-07 04:54:49.847794 +00:00', '2024-01-07 05:13:57.880722 +00:00', '2024-01-07 04:54:49.847794 +00:00',
        '', '', '', '0001-01-01 00:00:00.000000 +00:00', false, 0, '0001-01-01 00:00:00.000000 +00:00', null);