create table "Roles"
(
    id              uuid                                    NOT NULL primary key,
    name            varchar(255)                            NOT NULL,
    description     varchar(255),
    editable        boolean default true,
    created_at      timestamp with time zone                not null,
    updated_at      timestamp with time zone                not null,
    created_by      uuid
        constraint roles_created_by_fkey references "Users" not null,
    organization_id uuid
        constraint roles_org_id_fkey
            references "Organizations"
);

create index roles_org_id
    on "Roles" ("organization_id");

