create table keys
(
    id         varchar(36) not null primary key,
    account_id varchar(36) foreign key references accounts(id),
    name       nvarchar(max),
    type       nvarchar(max),
    created_at datetimeoffset,
    updated_at datetimeoffset,
    deleted_at datetimeoffset
)
go

