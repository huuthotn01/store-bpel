INSERT INTO `warehouse` (`warehouse_code`, `warehouse_name`, `capacity`, `created_at`, `street`, `ward`, `district`, `province`) VALUES
('warehouse-1', 'Kho chính', 2000, '2023-03-31 17:00:00', 'Lý Thường Kiệt', 'Phường 11', 'Quận 10', 'Hồ Chí minh'),
('warehouse-2', 'Kho Quận 10', 1000, '2022-07-31 17:00:00', 'Tô Hiến Thành', 'Phường 11', 'Quận 10', 'Hồ Chí minh'),
('warehouse-3', 'Kho Thủ Đức', 1000, '2022-12-09 17:00:00', 'Võ Văn Ngân', 'Phường 2', 'Thủ Đức', 'Hồ Chí minh');


INSERT INTO `staff_in_wh` (`staff_code`, `warehouse_code`, `started_date`, `end_date`, `role`) VALUES
('warehouse1', 'warehouse-1', '2018-01-11 17:01:00', NULL, 'MANAGER'),
('staff-w1', 'warehouse-1', '2018-04-11 17:01:00', NULL, 'STAFF'),
('staff-w2', 'warehouse-2', '2020-04-11 17:01:00', NULL, 'STAFF'),
('staff-w3', 'warehouse-3', '2022-04-11 17:01:00', NULL, 'STAFF');