create table if not exists goods_img (
    goods_code varchar(100) not null,
    goods_color varchar(10) not null,
    goods_img varchar(100),
    is_default tinyint,
    constraint PK_goods_img primary key (goods_code, goods_color, goods_img)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;

create table if not exists goods_in_wh (
    goods_code varchar(100) not null,
    goods_size varchar(5) not null,
    goods_color varchar(10) not null,
    wh_code varchar(100),
    quantity int,
    created_date timestamp not null default current_timestamp,
    updated_date timestamp not null default current_timestamp on update current_timestamp,
    constraint PK_goods_in_wh primary key (goods_code, goods_size, goods_color, wh_code)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;