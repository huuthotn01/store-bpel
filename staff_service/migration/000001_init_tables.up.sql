create table if not exists staff (
    staff_id varchar(100) not null,
    staff_name varchar(100) not null,
    province varchar(50) not null,
    district varchar(50) not null,
    ward varchar(50) not null,
    street varchar(50) not null,
    birthdate date not null,
    hometown varchar(100) not null,
    citizen_id varchar(50) not null,
    staff_position varchar(50) not null,
    start_date timestamp not null default current_timestamp,
    salary int not null,
    gender varchar(10) not null,
    phone varchar(12),
    email varchar(50),
    constraint PK_staff primary key (staff_id)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;

create table if not exists account (
    username varchar(100) not null,
    staff_id varchar(100) not null,
    constraint PK_account primary key (username)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;

create table if not exists attendance (
    staff_id varchar(100) not null,
    attendance_date date not null,
    checkin_time timestamp not null default current_timestamp,
    checkout_time timestamp,
    constraint PK_attendance primary key (staff_id, attendance_date)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;

create table if not exists requests (
    id varchar(100) not null,
    request_date date not null,
    request_type int not null,
    staff_id varchar(100) not null,
    status varchar(50) not null,
    constraint PK_requests primary key (id)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;
