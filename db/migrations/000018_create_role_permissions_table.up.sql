CREATE TABLE "Role_Permissions"
(
    id              uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    "Role_id"       uuid                            NOT NULL,
    "Permission_id" uuid                         NOT NULL,
    FOREIGN KEY ("Role_id") REFERENCES "Roles" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("Permission_id") REFERENCES "Permissions" ("id") ON DELETE CASCADE
);

-- Create index for role_id and permission_id
CREATE INDEX "Role_Permissions_role_id_idx" ON "Role_Permissions" ("Role_id");
CREATE INDEX "Role_Permissions_permission_id_idx" ON "Role_Permissions" ("Permission_id");