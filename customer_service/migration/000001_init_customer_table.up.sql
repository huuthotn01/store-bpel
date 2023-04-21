create table if not exists customer (
    username varchar(100) not null,
    customer_name varchar(100) not null,
    customer_email varchar(100) not null,
    customer_phone varchar(15) not null,
    customer_age int,
    customer_gender enum('MALE', 'FEMALE', 'UNDEFINED') not null,
    street varchar(100),
    ward varchar(50),
    district varchar(50),
    province varchar(50),
    image varchar(100),
    constraint PK_customer primary key (username)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;