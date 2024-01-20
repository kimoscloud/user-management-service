CREATE TABLE IF NOT EXISTS "Team_Projects"
(
    id         UUID    DEFAULT uuid_generate_v4() PRIMARY KEY,
    team_id    UUID         NOT NULL,
    project_id UUID         NOT NULL,
    is_active  BOOLEAN DEFAULT true,
    status     VARCHAR(255) NOT NULL,
    role       VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL,
    updated_at TIMESTAMPTZ  NOT NULL,
    deleted_at TIMESTAMPTZ
);

-- Crear la restricción de clave foránea (FK) para team_id
ALTER TABLE "Team_Projects"
    ADD CONSTRAINT fk_team_projects_team
        FOREIGN KEY (team_id)
            REFERENCES "Teams" (id)
            ON DELETE CASCADE;

ALTER TABLE "Team_Projects"
    ADD CONSTRAINT fk_team_projects_project
        FOREIGN KEY (project_id)
            REFERENCES "Projects" (id)
            ON DELETE CASCADE;

-- Add index
CREATE INDEX idx_team_projects_team_id
    ON "Team_Projects" (team_id);
CREATE INDEX idx_team_projects_project_id
    ON "Team_Projects" (project_id);
CREATE INDEX idx_team_projects_project_id_team_id
    ON "Team_Projects" (project_id, team_id);