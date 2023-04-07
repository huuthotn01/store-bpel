create table if not exists goods (
    goods_code varchar(100) not null,
    goods_size varchar(5) not null,
    goods_color varchar(10) not null,
    goods_name varchar(100) not null,
    order_code int not null,
    quantity int not null,
    unit_price int not null,
    total_price int not null,
    tax double,
    image varchar(100),
    promotion double,
    constraint PK_goods primary key (goods_code, goods_size, goods_color, order_code)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;

create table if not exists orders (
    order_code int not null auto_increment,
    transaction_date timestamp not null default current_timestamp,
    total_price int not null,
    public_order_code varchar(100) not null,
    constraint PK_orders primary key (order_code),
    constraint UK_orders unique key (public_order_code)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;

create table if not exists online_orders (
    order_code int not null,
    expected_delivery date,
    shipping_fee int not null,
    customer_id varchar(100) not null,
    payment_method varchar(50) not null,
    street varchar(100),
    ward varchar(100),
    district varchar(100),
    province varchar(100),
    customer_name varchar(100),
    customer_phone varchar(100),
    customer_email varchar(100),
    status int not null,
    constraint PK_online_orders primary key (order_code)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;

create table if not exists store_orders (
    order_code int not null,
    store_code varchar(100) not null,
    staff_id varchar(100) not null,
    constraint PK_store_orders primary key (order_code)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;

create table if not exists order_state (
    order_code int not null,
    state varchar(50) not null,
    state_time timestamp not null default current_timestamp,
    constraint PK_orders_state primary key (order_code, state, state_time)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;