create table users
(
    id         varchar(36) not null primary key,
    name       varchar(255),
    email      varchar(255),
    address    varchar(255),
    cpf        varchar(11),
    phone      varchar(20),
    birthday   date,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime
)

