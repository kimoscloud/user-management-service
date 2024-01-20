CREATE TABLE "Projects"
(
    id              UUID                  DEFAULT uuid_generate_v4() PRIMARY KEY,
    name            VARCHAR(255) NOT NULL,
    description     TEXT,
    slug            VARCHAR(255) NOT NULL,
    image_url       VARCHAR(255),
    created_by      UUID         NOT NULL,
    is_active       BOOLEAN               DEFAULT true,
    organization_id UUID, -- organization_id is the owner of the project
    user_id         UUID, -- user_id is the owner of the project
    created_at      TIMESTAMPTZ  NOT NULL DEFAULT current_timestamp,
    updated_at      TIMESTAMPTZ  NOT NULL DEFAULT current_timestamp,
    deleted_at      TIMESTAMPTZ
);

CREATE INDEX idx_projects_deleted_at ON "Projects" (deleted_at);
CREATE INDEX idx_projects_user ON "Projects" (user_id);
CREATE INDEX idx_projects_name ON "Projects" (name);
CREATE INDEX idx_projects_slug ON "Projects" (slug);
CREATE INDEX idx_projects_organization ON "Projects" (organization_id);

ALTER TABLE "Projects"
    ADD CONSTRAINT fk_project_organization
        FOREIGN KEY (organization_id)
            REFERENCES "Organizations" (id)
            ON DELETE CASCADE;

ALTER TABLE "Projects"
    ADD CONSTRAINT fk_project_user
        FOREIGN KEY (user_id)
            REFERENCES "Users" (id)
            ON DELETE CASCADE;
