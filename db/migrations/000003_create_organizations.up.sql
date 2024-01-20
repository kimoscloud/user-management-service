CREATE TABLE "Organizations"
(
    id                       uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name                     VARCHAR(255),
    created_by               uuid                            NOT NULL REFERENCES "Users" (id) ON DELETE CASCADE,
    slug                     VARCHAR(255),
    billing_email            VARCHAR(255),
    url                      VARCHAR(255),
    about                    TEXT,
    logo_url                 VARCHAR(255),
    background_image_url     VARCHAR(255),
    plan                     VARCHAR(255),
    current_period_starts_at TIMESTAMP WITH TIME ZONE,
    current_period_ends_at   TIMESTAMP WITH TIME ZONE,
    subscription_id          VARCHAR(255),
    status                   VARCHAR(255),
    timezone                 VARCHAR(255),
    created_at               TIMESTAMP WITH TIME ZONE        NOT NULL,
    updated_at               TIMESTAMP WITH TIME ZONE        NOT NULL,
    deleted_at               TIMESTAMP WITH TIME ZONE
);

CREATE UNIQUE INDEX lower_organizations_unique_idx
    ON "Organizations" (LOWER(slug::TEXT));

CREATE INDEX idx_organizations_id
    ON "Organizations" (id);

