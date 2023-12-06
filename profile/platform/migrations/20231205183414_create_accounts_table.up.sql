create table accounts
(
    id         varchar(36) not null primary key,
    user_id    varchar(36) foreign key references users(id),
    balance    float,
    agency     nvarchar(max),
    bank       nvarchar(max),
    created_at datetimeoffset,
    updated_at datetimeoffset,
    deleted_at datetimeoffset
    )
go

