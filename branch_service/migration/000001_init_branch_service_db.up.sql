create table if not exists branch (
    branch_code int auto_increment,
    branch_name varchar(100),
    branch_province varchar(50),
    branch_district varchar(50),
    branch_ward varchar(50),
    branch_street varchar(50),
    created_at DATETIME,
    manager varchar(100),
    open_time time,
    close_time time,
    constraint PK_branch primary key (branch_code)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;

create table if not exists branch_img (
    branch_code int,
    branch_img varchar(100),
    constraint PK_branch_img primary key (branch_code, branch_img)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;

create table if not exists branch_manager (
    branch_code int,
    manager_code varchar(100),
    start_date timestamp not null default CURRENT_TIMESTAMP,
    end_date timestamp,
    constraint PK_branch_manager primary key (branch_code, manager_code)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;

create table if not exists branch_staff (
    branch_code int,
    staff_code varchar(100),
    start_date timestamp not null default CURRENT_TIMESTAMP,
    end_date timestamp,
    constraint PK_branch_staff primary key (branch_code, staff_code)
) engine = InnoDB default charset = utf8mb4 collate = utf8mb4_unicode_ci ;