create table if not exists cart (
    cart_code varchar(100) not null,
    customer_code varchar(100) not null,
    quantity int not null,
    goods_code varchar(100) not null,
    goods_size varchar(5),
    goods_color varchar(10),
    constraint PK_cart primary key (goods_code, goods_size, goods_color, cart_code, customer_code)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;

create table if not exists goods (
    goods_code varchar(100) not null,
    goods_size varchar(5) not null,
    goods_color varchar(10) not null,
    constraint PK_goods primary key (goods_code, goods_size, goods_color)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;