create table if not exists cart (
    cart_id int not null auto_increment ,
    customer_id varchar(100) not null,
    constraint PK_cart primary key (cart_id)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;

create table if not exists goods (
    cart_id varchar(100) not null,
    goods_id varchar(100) not null,
    goods_size varchar(5),
    goods_color varchar(10),
    quantity int not null,
    constraint PK_goods primary key (cart_id, goods_id, goods_size, goods_color)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;