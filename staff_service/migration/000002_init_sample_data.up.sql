INSERT INTO `account` (`username`, `staff_id`) VALUES
('admin1', 'admin1'),
('branch1', 'branch1'),
('goods1', 'goods1'),
('lead_branch1', 'lead_branch1'),
('warehouse1', 'warehouse1');

INSERT INTO `attendance` (`staff_id`, `attendance_date`, `checkin_time`, `checkout_time`) VALUES
('12345', '2023-02-01', '2023-02-17 00:00:00', '2023-02-01 07:00:30'),
('12345', '2023-02-02', '2023-02-17 00:00:00', '2023-02-01 11:00:30'),
('12345', '2023-02-03', '2023-02-17 00:00:00', '2023-02-01 11:00:30');

INSERT INTO `requests` (`id`, `request_date`, `request_type`, `staff_id`, `status`) VALUES
('1', '2023-02-09', 'ADD', '1', 'PENDING'),
('3', '2023-02-09', 'DELETE', '12345', 'PENDING');

INSERT INTO `staff` (`staff_id`, `staff_name`, `province`, `district`, `ward`, `street`, `birthdate`, `hometown`, `citizen_id`, `staff_position`, `start_date`, `salary`, `gender`, `phone`, `email`, `branch_id`, `status`) VALUES
('1', 'aaaa', 'a', 'b', 'x', 'd', '2023-02-01', 'DL', '038746453', '7', '2023-02-17 07:53:25', 1200000, 'FEMALE', '0294756342', 'aaaa@gmail.com', '1', 'PENDING'),
('12345', 'Bùi Hiển', 'Bình dương', 'Dĩ An', 'Đông Hòa', 'Lương Định Của', '2000-02-22', 'Daklak', '0192756559', '7', '2023-01-31 10:00:00', 1000000, 'MALE', '0294756342', 'a@gmail.com', '1', 'APPROVED'),
('3', 'sdfg', 'a', 'b', 'x', 'd', '2023-02-01', 'DL', '038746873', '7', '2023-02-17 07:54:03', 1200000, 'FEMALE', '0294756342', '3@gmail.com', '1', 'PENDING'),
('admin1', 'admin', 'Lương Định Của', 'Đông Hòa', 'Dĩ An', 'Bình Dương', '2001-07-19', 'DakLak', '0192756559343', '2', '2023-03-21 08:09:18', 123123344, 'UNDEFINED', '04536821', 'admin@gmail.com', '2', 'APPROVED'),
('branch1', 'branch1 Bùi', 'Võ Văn Ngân', '11', 'Thủ Đức', 'HCM', '1995-03-09', 'Đà Lạt', '3772487934512', '3', '2023-03-27 04:28:40', 12000000, 'MALE', '0345288743', 'branch1@gmail.com', NULL, 'APPROVED'),
('goods1', 'goods1 Trần', 'Võ Văn Ngân', '14', 'Dĩ An', 'Bình Dương', '1995-03-20', 'Đà Lạt', '3772487934334', '5', '2023-03-27 04:32:04', 13000000, 'MALE', '03452846743', 'goods1@gmail.com', NULL, 'APPROVED'),
('lead_branch1', 'lead_branch1 Nguyễn', 'Tô Ngọc Vân', '11', 'Thủ Đức', 'HCM', '1998-09-09', 'Bình Phước', '37722357934512', '6', '2023-03-27 04:32:04', 10000000, 'FEMALE', '0345288743', 'lead_branch1@gmail.com', '2', 'APPROVED'),
('warehouse1', 'warehouse1 Lê', 'Võ Văn Ngân', '11', 'Thủ Đức', 'HCM', '1995-03-09', 'Đà Lạt', '3776487934512', '4', '2023-03-27 04:32:05', 11000000, 'UNDEFINED', '0345288743', 'warehouse1@gmail.com', NULL, 'APPROVED');
