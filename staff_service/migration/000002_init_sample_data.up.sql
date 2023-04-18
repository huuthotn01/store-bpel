INSERT INTO `account` (`username`, `staff_id`) VALUES
('1', '1'),
('12345', '12345'),
('3', '3'),
('admin1', 'admin1'),
('branch1', 'branch1'),
('d', 'd'),
('goods1', 'goods1'),
('hoLe', 'hoLe'),
('lead_branch1', 'lead_branch1'),
('staff-w1', 'staff-w1'),
('staff-w2', 'staff-w2'),
('staff-w3', 'staff-w3'),
('warehouse1', 'warehouse1');

INSERT INTO `attendance` (`staff_id`, `attendance_date`, `checkin_time`, `checkout_time`) VALUES
('12345', '2023-02-01', '2023-02-17 00:00:00', '2023-02-01 07:00:30'),
('12345', '2023-02-02', '2023-02-17 00:00:00', '2023-02-01 11:00:30'),
('12345', '2023-02-03', '2023-02-17 00:00:00', '2023-02-01 11:00:30');

INSERT INTO `requests` (`id`, `request_date`, `request_type`, `staff_id`, `status`) VALUES
('1', '2023-02-09', 'ADD', '1', 'APPROVED'),
('3', '2023-02-09', 'DELETE', '12345', 'UNAPPROVED'),
('add_1681309271', '2023-04-12', 'ADD', 'd', 'PENDING'),
('del_1681309108', '2023-04-12', 'DELETE', 'hoLe', 'PENDING');



INSERT INTO `staff` (`staff_id`, `staff_name`, `province`, `district`, `ward`, `street`, `birthdate`, `hometown`, `citizen_id`, `staff_position`, `start_date`, `salary`, `gender`, `phone`, `email`, `branch_id`, `status`) VALUES
('1', 'A Bùi', 'Bình dương', 'Dĩ An', 'Đông Hòa', 'Lương Định Của', '2023-02-01', 'DL', '038746453', '7', '2023-02-16 17:53:25', 1200000, 'FEMALE', '0294756342', 'aaaa@gmail.com', '1', 'PENDING'),
('12345', 'Bùi Hiển', 'Bình dương', 'Dĩ An', 'Đông Hòa', 'Lương Định Của', '2000-02-22', 'Daklak', '0192756559', '7', '2023-01-30 20:00:00', 1000000, 'MALE', '0294756342', 'a@gmail.com', '1', 'APPROVED'),
('3', 'Trần S', 'Bình dương', 'Dĩ An', 'Đông Hòa', 'Lương Định Của', '2023-02-01', 'DL', '038746873', '7', '2023-02-16 17:54:03', 1200000, 'FEMALE', '0294756342', '3@gmail.com', '1', 'PENDING'),
('admin1', 'admin', 'Lương Định Của', 'Đông Hòa', 'Dĩ An', 'Bình Dương', '2001-07-19', 'DakLak', '0192756559343', '2', '2023-03-20 18:09:18', 123123344, 'UNDEFINED', '04536821', 'admin@gmail.com', '2', 'APPROVED'),
('branch1', 'branch1 Bùi', 'Võ Văn Ngân', '11', 'Thủ Đức', 'HCM', '1995-03-09', 'Đà Lạt', '3772487934512', '3', '2023-03-26 14:28:40', 12000000, 'MALE', '0345288743', 'branch1@gmail.com', NULL, 'APPROVED'),
('d', 'Võ Ngọc D', 'Hồ Chí Minh', '10', '14', 'Tô Hiến Thành', '2004-04-09', 'Bình Dương', '02948572342112', '7', '2023-04-12 07:15:38', 500000, 'FEMALE', '0475638354', 'd@gmail.com', '2', 'PENDING'),
('goods1', 'goods1 Trần', 'Võ Văn Ngân', '14', 'Dĩ An', 'Bình Dương', '1995-03-20', 'Đà Lạt', '3772487934334', '5', '2023-03-26 14:32:04', 13000000, 'MALE', '03452846743', 'goods1@gmail.com', NULL, 'APPROVED'),
('hoLe', 'Lê Ngọc Hồ', 'Hồ Chí Minh', '10', '14', 'Tô Hiến Thành', '1994-04-30', 'Bình Dương', '0294852342112', '7', '2023-04-12 07:17:15', 500000, 'FEMALE', '0475633354', 'hoLe@gmail.com', '3', 'APPROVED'),
('lead_branch1', 'lead_branch1 Nguyễn', 'Tô Ngọc Vân', '11', 'Thủ Đức', 'HCM', '1998-09-09', 'Bình Phước', '37722357934512', '6', '2023-03-26 14:32:04', 10000000, 'FEMALE', '0345288743', 'lead_branch1@gmail.com', '2', 'APPROVED'),
('staff-w1', 'Phan Anh', 'Bình dương', 'Dĩ An', 'Tân Hòa', 'Nguyễn Thái Học', '2001-02-22', 'Phú Yêu', '484638345245', '7', '2023-03-31 14:27:54', 1200000, 'FEMALE', '043456563', 'w1@gmail.com', NULL, 'APPROVED'),
('staff-w2', 'Đinh L', 'Bình dương', 'Dĩ An', 'Tân Hòa', 'Nguyễn Thái Học', '2001-02-26', 'Tây Ninh', '484638334245', '7', '2023-03-31 14:29:31', 2200000, 'FEMALE', '043455563', 'w2@gmail.com', NULL, 'APPROVED'),
('staff-w3', 'Trương M', 'Bình dương', 'Dĩ An', 'Tân Hòa', 'Nguyễn Thái Học', '2001-12-22', 'Phú Yêu', '486638345245', '7', '2023-03-31 14:29:32', 3200000, 'FEMALE', '043457563', 'w3@gmail.com', NULL, 'APPROVED'),
('warehouse1', 'warehouse1 Lê', 'Võ Văn Ngân', '11', 'Thủ Đức', 'HCM', '1995-03-09', 'Đà Lạt', '3776487934512', '4', '2023-03-26 14:32:05', 11000000, 'UNDEFINED', '0345288743', 'warehouse1@gmail.com', NULL, 'APPROVED');
