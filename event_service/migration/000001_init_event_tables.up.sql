create table if not exists event (
    event_id int auto_increment,
    name varchar(100),
    discount float,
    start_time timestamp default current_timestamp,
    end_time timestamp default current_timestamp,
    image varchar(200),
    constraint PK_event primary key (event_id)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;

create table if not exists goods (
    event_id int,
    goods_id varchar(100),
    constraint PK_event_with_goods primary key (event_id, goods_id)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;