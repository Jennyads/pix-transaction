create table users
(
    id         varchar(36) not null primary key,
    name       nvarchar(max),
    email      nvarchar(max),
    address    nvarchar(max),
    cpf        nvarchar(11),
    phone      nvarchar(max),
    birthday   datetimeoffset,
    created_at datetimeoffset,
    updated_at datetimeoffset,
    deleted_at datetimeoffset
)
go

