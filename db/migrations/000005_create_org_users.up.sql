create table "Organization_Users"
(
    id                 uuid                     not null
        primary key,
    organization_id             uuid                     not null
        constraint "Organization_Users_organization_id_fkey"
            references "Organizations"
            on update cascade on delete cascade,
    user_id            uuid                  not null
        references "Users"
            on update cascade,
    role               text,
    created_at         timestamp with time zone not null,
    updated_at         timestamp with time zone not null,
    status             varchar(255),
    role_id    uuid
        constraint organization_users_permission_role_fkey
            references "Roles",
    invite_sent_at     timestamp with time zone,
    created_by_user_id uuid ,
    is_active          boolean default true
);

create index organization__users_org_id
    on "Organization_Users" (organization_id);

create unique index organization__users_org_id_user_id_unique
    on "Organization_Users" (organization_id, user_id);

create index organization__users_organization_id
    on "Organization_Users" (organization_id);

create index organization__users_user_id
    on "Organization_Users" (user_id);

