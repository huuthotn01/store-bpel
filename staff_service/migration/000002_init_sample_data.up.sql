INSERT INTO `account` (`username`, `staff_id`) VALUES
('admin', 'admin'),
('anLe', 'anLe'),
('aTran', 'aTran'),
('binhMac', 'binhMac'),
('ducbinh', 'ducbinh'),
('lydinh', 'lydinh'),
('nguyen', 'nguyen'),
('quyenBui', 'quyenBui');

INSERT INTO `staff` (`staff_id`, `staff_name`, `province`, `district`, `ward`, `street`, `birthdate`, `hometown`, `citizen_id`, `staff_position`, `start_date`, `salary`, `gender`, `phone`, `email`, `branch_id`, `status`) VALUES
('admin', 'Philip Bùi', 'TP. HCM', 'Q. 10', 'P. 11', 'To Hien Thanh', '2001-06-13', 'Hà Nội', '112233445', '2', '2023-04-26 08:43:42', 15000000, 'MALE', '092618234', 'admin@hcmut.vn', NULL, 'APPROVED'),
('anLe', 'Lê Hồng Ân', 'Tp.Hồ Chí Minh', 'Q.10', 'P.14', 'Lý Thường Kiệt', '2000-02-23', 'Bình Dương', '045734564', '7', '2023-04-26 10:05:01', 7500000, 'UNDEFINED', '04737823', 'anLe@gmail.com', 'branch_1682526290', 'APPROVED'),
('aTran', 'Trần Ngọc A', 'Hà Nội', 'Ba Đình', 'Thành Công', 'Láng Hạ', '2003-06-21', 'Hà Nội', '1238858341123', '7', '2023-04-26 09:58:59', 8000000, 'FEMALE', '0387527345', 'aTran@gmail.com', 'branch_1682526367', 'APPROVED'),
('binhMac', 'Mạc Hữu Bình', 'Bình Dương', 'Dĩ An', 'Đông Hòa', 'Nguyễn Thái Học', '2002-06-05', 'Bình Thuận', '74452389363', '7', '2023-04-26 10:10:05', 6000000, 'MALE', '0648385522', 'binhMac@gmail.com', 'branch_1682526156', 'APPROVED'),
('ducbinh', 'Bùi Đức Bình', 'Bình Dương', 'Dĩ An', 'Đông Hòa', '163 Lương Định Của', '1993-08-21', 'Daklak', '32445341231', '3', '2023-04-26 09:39:54', 7000000, 'MALE', '0387246134', 'ducbinh@gmail.com', 'branch_1682526156', 'APPROVED'),
('lydinh', 'Đinh Ly', 'Tp.Hồ Chí Minh', 'Thủ Đức', 'KP.1', 'Võ văn Ngân', '2001-04-13', 'Lâm Đồng', '3948356434234', '7', '2023-04-26 10:07:42', 8000000, 'FEMALE', '038462635', 'lydinh@gmail.com', 'branch_1682526238', 'APPROVED'),
('nguyen', 'Nguyễn Hải Âu', 'Tp.Hồ Chí Minh', 'Thủ Đức', 'KP.10', 'Hoàng Diệu', '1997-02-21', 'Đồng Nai', '9435878346', '6', '2023-04-26 10:12:37', 9650000, 'UNDEFINED', '038475234', 'nguyen@gmail.com', 'branch_1682526238', 'APPROVED'),
('quyenBui', 'Bùi Vương Quyền', 'Tp.Hồ Chí Minh', 'Q.10', 'P.14', 'Tô Hiến Thánh', '1995-08-25', 'Tp. Hồ Chí Minh', '034775023405', '6', '2023-04-26 10:06:20', 12000000, 'MALE', '0387527345', 'quyenBui@gmail.com', 'branch_1682526290', 'APPROVED');
