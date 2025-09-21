create table users(
    id uuid primary key not null default gen_random_uuid(),
    email varchar(255) not null,
    password varchar(255) not null,
    isActive boolean not null default true,
    createdAt timestamptz not null default now() 
);