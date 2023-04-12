create table if not exists goods (
    goods_code varchar(100) not null,
    goods_size varchar(5) not null,
    goods_color varchar(10) not null,
    goods_type varchar(100),
    goods_gender int,
    goods_cost int,
    unit_price int,
    quantity int,
    order_code varchar(100),
    constraint PK_goods primary key (goods_code, goods_size, goods_color,order_code)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;

create table if not exists orders (
    order_code varchar(100),
    transaction_date timestamp default current_timestamp,
    shop_code varchar(100),
    constraint PK_orders primary key (order_code)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;