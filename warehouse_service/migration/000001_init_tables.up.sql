create table if not exists warehouse (
    warehouse_code varchar(100) not null,
    warehouse_name varchar(100) not null,
    capacity int not null,
    created_at timestamp not null default current_timestamp,
    street varchar(100),
    ward varchar(100),
    district varchar(100),
    province varchar(100),
    constraint PK_warehouse primary key (warehouse_code)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;

-- create table if not exists warehouse_staff (
--      staff_code varchar(100) not null,
--      constraint PK_warehouse_staff primary key (staff_code)
-- ) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;

-- create table if not exists warehouse_manager (
--     staff_code varchar(100) not null,
--     warehouse_code varchar(100) not null,
--     started_date timestamp not null default current_timestamp,
--     end_date timestamp null on update current_timestamp,
--     constraint PK_warehouse_manager primary key (staff_code, warehouse_code)
-- ) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;

create table if not exists staff_in_wh (
    staff_code varchar(100) not null,
    warehouse_code varchar(100) not null,
    started_date timestamp not null default current_timestamp,
    end_date timestamp null on update current_timestamp,
    role ENUM('MANAGER','STAFF') not null,
    constraint PK_staff_in_wh primary key (staff_code, warehouse_code)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;