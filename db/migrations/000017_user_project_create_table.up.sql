CREATE TABLE IF NOT EXISTS "User_Projects"
(
    id         UUID    DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id    UUID         NOT NULL,
    project_id UUID         NOT NULL,
    is_active  BOOLEAN DEFAULT true,
    status     VARCHAR(255) NOT NULL,
    role       VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL,
    updated_at TIMESTAMPTZ  NOT NULL,
    deleted_at TIMESTAMPTZ
);

-- Crear la restricción de clave foránea (FK) para team_id
ALTER TABLE "User_Projects"
    ADD CONSTRAINT fk_user_projects_user
        FOREIGN KEY (user_id)
            REFERENCES "Users" (id)
            ON DELETE CASCADE;

ALTER TABLE "User_Projects"
    ADD CONSTRAINT fk_user_projects_project
        FOREIGN KEY (project_id)
            REFERENCES "Projects" (id)
            ON DELETE CASCADE;

-- Add index
CREATE INDEX idx_user_projects_team_id
    ON "User_Projects" (user_id);
CREATE INDEX idx_user_projects_project_id
    ON "User_Projects" (project_id);
CREATE INDEX idx_user_projects_project_id_user_id
    ON "User_Projects" (project_id, user_id);