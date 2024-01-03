CREATE TABLE "Permissions"
(
    id            uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name          VARCHAR(255),
    description   VARCHAR(255),
    internal_name VARCHAR(255),
    created_at    TIMESTAMP WITH TIME ZONE        NOT NULL,
    updated_at    TIMESTAMP WITH TIME ZONE        NOT NULL,
    deleted_at    TIMESTAMP WITH TIME ZONE
);

CREATE UNIQUE INDEX lower_permissions_internal_name_unique_idx
    ON "Permissions" (LOWER(internal_name::TEXT));

CREATE INDEX idx_permissions_id
    ON "Permissions" (id);

CREATE UNIQUE INDEX permissions_internal_name_unique_idx
    ON "Permissions" (internal_name);
