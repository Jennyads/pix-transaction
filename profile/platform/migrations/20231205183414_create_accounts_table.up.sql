create table accounts
(
    id         varchar(36) not null primary key,
    user_id    varchar(36) foreign key references users(id),
    balance    decimal(10,2),
    agency     varchar(100),
    bank      varchar(100),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    )
go

