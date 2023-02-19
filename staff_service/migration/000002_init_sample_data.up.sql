INSERT INTO `account` (`username`, `staff_id`) VALUES
('VH', '2');

INSERT INTO `attendance` (`staff_id`, `attendance_date`, `checkin_time`, `checkout_time`) VALUES
('12345', '2023-02-01', '2023-02-17 00:00:00', '2023-02-01 07:00:30'),
('12345', '2023-02-02', '2023-02-17 00:00:00', '2023-02-01 11:00:30'),
('12345', '2023-02-03', '2023-02-17 00:00:00', '2023-02-01 11:00:30');

INSERT INTO `requests` (`id`, `request_date`, `request_type`, `staff_id`, `status`) VALUES
('1', '2023-02-09', 'ADD', '1', 'PENDING'),
('2', '2023-02-09', 'ADD', '3', 'PENDING');
('3', '2023-02-09', 'DELETE', '12345', 'PENDING');

INSERT INTO `staff` (`staff_id`, `staff_name`, `province`, `district`, `ward`, `street`, `birthdate`, `hometown`, `citizen_id`, `staff_position`, `start_date`, `salary`, `gender`, `phone`, `email`, `status`, `branch_id`) VALUES
('1', 'aaaa', 'a', 'b', 'x', 'd', '2023-02-01', 'DL', '038746453', 'nhân viên', '2023-02-17 14:53:25', 1200000, 'FEMALE', '0294756342', 'aaaa@gmail.com', 'PENDING', '1'),
('12345', 'Bùi Hiển', 'Bình dương', 'Dĩ An', 'Đông Hòa', 'Lương Định Của', '2000-02-22', 'Daklak', '0192756559', 'nhân viên', '2023-01-31 17:00:00', 1000000, 'MALE', '0294756342', 'a@gmail.com', 'APPROVED' , '1'),
('2', 'VH', 'a', 'b', 'c', 'd', '2001-02-22', 'DL', '092929', 'Quản lý', '2023-02-06 14:48:51', 3000000, 'MALE', '0294756333', 'vh@gmail.com', 'APPROVED' , '1'),
('3', 'sdfg', 'a', 'b', 'x', 'd', '2023-02-01', 'DL', '038746873', 'nhân viên', '2023-02-17 14:54:03', 1200000, 'FEMALE', '0294756342', '3@gmail.com', 'PENDING' , '1');
