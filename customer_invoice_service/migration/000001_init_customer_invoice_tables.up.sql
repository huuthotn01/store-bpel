create table if not exists goods (
    goods_code varchar(100) not null,
    invoice_code varchar(100) not null,
    goods_quantity int not null ,
    goods_size varchar(5) not null ,
    goods_color varchar(10) not null,
    unit_price int not null ,
    total_price int not null ,
    tax int not null ,
    goods_img varchar(100),
    constraint PK_goods primary key (goods_code, goods_size, goods_color, invoice_code)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;

create table if not exists invoice (
    invoice_code varchar(100) not null,
    transaction_date date not null,
    total_price int not null,
    customer_code varchar(100),
    expected varchar(50),
    shipping_fee int,
    payment_method varchar(100),
    constraint PK_invoice primary key (invoice_code)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;

create table if not exists invoice_state (
    invoice_code varchar(100) not null,
    state varchar(250) not null ,
    time timestamp not null default current_timestamp,
    constraint PK_invoice_state primary key (invoice_code, state, time)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;