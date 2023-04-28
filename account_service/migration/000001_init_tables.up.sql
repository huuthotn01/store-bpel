create table if not exists account (
    username varchar(100) not null,
    password varchar(120) not null,
    user_role int not null,
    email varchar(100) not null,
    otp varchar(6),
    otp_timeout timestamp null,
    is_activated tinyint not null default 1,
    created_at timestamp not null default current_timestamp,
    constraint PK_account primary key (username)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;