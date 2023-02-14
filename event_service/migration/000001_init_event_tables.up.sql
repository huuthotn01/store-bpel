create table if not exists event (
    event_code varchar(100),
    event_name varchar(100),
    discount int,
    start_time timestamp,
    end_time timestamp,
    constraint PK_event primary key (event_code)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;

create table if not exists event_with_goods (
    event_code varchar(100),
    goods_code varchar(100),
    constraint PK_event_with_goods primary key (event_code, goods_code)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;