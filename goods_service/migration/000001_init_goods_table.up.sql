create table if not exists goods (
    goods_code varchar(100) not null,
    goods_size varchar(5) not null,
    goods_color varchar(10) not null,
    goods_name varchar(100) not null,
    goods_type varchar(100) not null,
    goods_gender int not null,
    goods_age ENUM('KID', 'ADULT', 'ALL') not null,
    manufacturer varchar(100) not null,
    is_for_sale tinyint,
    unit_price int,
    unit_cost int not null,
    description varchar(100),
    constraint PK_goods primary key (goods_code, goods_size, goods_color)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;