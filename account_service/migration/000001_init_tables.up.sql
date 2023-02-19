create table if not exists account (
    username varchar(100) not null,
    password varchar(120) not null,
    user_role int not null,
    is_activated tinyint not null,
    created_at timestamp not null default current_timestamp,
    constraint PK_account primary key (username)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;