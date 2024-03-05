create table keys
(
    id         varchar(36) not null primary key,
    account_id int foreign key references accounts(id),
    name       varchar(200),
    type       varchar(100),
    created_at datetime,
    updated_at datetime,
    deleted_at datetime
)


