create table accounts
(
    id         int IDENTITY(1,1) PRIMARY KEY,
    user_id    varchar(36) foreign key references users(id),
    balance    decimal(10,2),
    agency     varchar(100),
    bank      varchar(100),
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
)

