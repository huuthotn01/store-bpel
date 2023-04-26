-- INSERT INTO `warehouse` (`warehouse_code`, `warehouse_name`, `capacity`, `created_at`, `street`, `ward`, `district`, `province`) VALUES
-- ('warehouse-1', 'Kho chính', 2000, '2023-03-31 17:00:00', 'Lý Thường Kiệt', 'Phường 11', 'Quận 10', 'Hồ Chí minh'),
-- ('warehouse-2', 'Kho Quận 10', 1000, '2022-07-31 17:00:00', 'Tô Hiến Thành', 'Phường 11', 'Quận 10', 'Hồ Chí minh'),
-- ('warehouse-3', 'Kho Thủ Đức', 1000, '2022-12-09 17:00:00', 'Võ Văn Ngân', 'Phường 2', 'Thủ Đức', 'Hồ Chí minh');


-- INSERT INTO `staff_in_wh` (`staff_code`, `warehouse_code`, `started_date`, `end_date`, `role`) VALUES
-- ('warehouse1', 'warehouse-1', '2018-01-11 17:01:00', NULL, 'MANAGER'),
-- ('staff-w1', 'warehouse-1', '2018-04-11 17:01:00', NULL, 'STAFF'),
-- ('staff-w2', 'warehouse-2', '2020-04-11 17:01:00', NULL, 'STAFF'),
-- ('staff-w3', 'warehouse-3', '2022-04-11 17:01:00', NULL, 'STAFF');


INSERT INTO `warehouse` (`warehouse_code`, `warehouse_name`, `capacity`, `created_at`, `street`, `ward`, `district`, `province`) VALUES
('warehouse-2023-04-26 17:37:43.759527 +0000 UTC', 'Kho Bình Dương', 1000, '2023-04-26 10:37:43', 'Lương Định Của', 'Đông Hòa', 'Dĩ An', 'Bình Dương'),
('warehouse-2023-04-26 17:38:08.1758413 +0000 UTC', 'Kho Thủ Đức', 500, '2023-04-26 10:38:08', 'Võ văn Ngân', 'KP.1', 'Thủ Đức', 'Tp.Hồ Chí Minh'),
('warehouse-2023-04-26 17:38:41.9647833 +0000 UTC', 'Kho Lý Thường Kiệt', 750, '2023-04-26 10:38:41', 'Lý Thường Kiệt', 'P.14', 'Q.10', 'Tp.Hồ Chí Minh'),
('warehouse-2023-04-26 17:39:13.4814632 +0000 UTC', 'Kho Hà Nội', 1200, '2023-04-26 10:39:13', 'Láng Hạ', 'Thành Công', 'Ba Đình', 'Hà Nội'),
('warehouse-2023-04-26 17:39:50.8681359 +0000 UTC', 'Kho Miền Nam', 2000, '2023-04-26 10:39:50', 'QL 1k', 'Đông Hòa', 'Dĩ An', 'Bình Dương');
