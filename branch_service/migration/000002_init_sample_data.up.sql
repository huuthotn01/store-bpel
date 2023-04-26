-- insert into branch values
--     (1, 'Branch 1', 'TP. HCM', 'Q. 10', 'P. 11', 'Ly Thuong Kiet', '2023-01-01 10:00:00', 'staff-1', '07:00:00', '18:00:00');

-- insert into branch_img values
--     (1, 'https://media.tarkett-image.com/medium/RF_SEE_7_SENSES_STORE_6.jpg');

-- insert into branch_manager values
--     (1, 'staff-1', '2023-01-01 00:00:00', null);

-- insert into branch_staff values
--     (1, 'staff-1', '2023-01-01 00:00:00', null),
--     (1, 'staff-2', '2023-01-01 00:00:00', null),
--     (1, 'staff-3', '2023-01-01 00:00:00', null);

INSERT INTO `branch` (`branch_code`, `branch_name`, `branch_province`, `branch_district`, `branch_ward`, `branch_street`, `created_at`, `manager`, `open_time`, `close_time`) VALUES
('branch_1682526156', 'Chi nhánh Bình Dương', 'Bình', 'Dĩ An', 'Đông Hòa', '163 Lương Định Của', '2023-04-26 16:22:36', '', '09:00:00', '23:00:00'),
('branch_1682526238', 'Chi nhánh Thủ Đức', 'Tp.Hồ Chí Minh', 'Thủ Đức', 'KP.1', '234 Võ Văn Ngân', '2023-04-26 16:23:58', '', '06:00:00', '22:30:00'),
('branch_1682526290', 'Chi nhánh Lý Thường Kiệt', 'Tp.Hồ Chí Minh', 'Quận 10', 'P.14', '268 Lý Thường Kiệt', '2023-04-26 16:24:50', '', '07:15:00', '22:15:00'),
('branch_1682526367', 'Chi nhánh Hà Nội', 'Hà Nội', 'Ba Đình', 'Thành Công', '22 Láng Hạ', '2023-04-26 16:26:07', '', '07:00:00', '22:00:00');


INSERT INTO `branch_staff` (`branch_code`, `staff_code`, `start_date`, `end_date`) VALUES
('branch_1682526156', 'binhMac', '2023-04-26 17:10:07', NULL),
('branch_1682526156', 'ducbinh', '2023-04-26 16:39:56', NULL),
('branch_1682526238', 'lydinh', '2023-04-26 17:07:44', NULL),
('branch_1682526238', 'nguyen', '2023-04-26 17:12:39', NULL),
('branch_1682526238', 'test', '2023-04-26 17:14:54', NULL),
('branch_1682526238', 'wer', '2023-04-26 17:20:15', NULL),
('branch_1682526290', 'anLe', '2023-04-26 17:05:03', NULL),
('branch_1682526290', 'quyenBui', '2023-04-26 17:06:22', NULL),
('branch_1682526367', 'aTran', '2023-04-26 16:59:01', NULL),
('branch_1682526367', 'test', '2023-04-26 17:19:31', NULL);