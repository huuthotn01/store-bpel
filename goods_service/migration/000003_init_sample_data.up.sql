INSERT INTO `goods` (`goods_code`, `goods_size`, `goods_color`, `goods_name`, `goods_type`, `goods_gender`, `goods_age`, `manufacturer`, `is_for_sale`, `unit_price`, `unit_cost`, `description`) VALUES
('1', '35', 'Red', 'Hàng để test', 'sandal', 3, 'KID', 'NSX TEST', 1, 10000, 9000, 'Hàng để test, xin nhẹ tay'),
('2', '30', 'Blue', 'Hàng để trưng bày', 'western shoes', 3, 'ALL', 'NSX TEST', 1, 200000, 150000, 'Hàng để trưng'),
('goods-1', 'M', 'Đen', 'Áo khoác gió', 'jacket', 1, 'ADULT', 'Nhà cung cấp Bình Dương', 1, 250000, 180000, 'Phù hợp với thời tiết nắng nóng'),
('goods-1', 'S', 'Đen', 'Áo khoác gió', 'jacket', 1, 'ADULT', 'Nhà cung cấp Bình Dương', 1, 250000, 190000, 'Phù hợp với thời tiết nắng nóng'),
('goods-2', 'L', 'Vàng', 'Quần tây nữ', 'Trousers', 2, 'ADULT', 'Nhà cung cấp Bình Dương', 1, 150000, 120000, 'Kiểu dáng thời thường'),
('goods-2', 'M', 'Vàng', 'Quần tây', 'jacket', 2, 'ADULT', 'Nhà cung cấp Bình Dương', 1, 150000, 120000, 'Phù hợp với thời tiết nắng nóng'),
('goods-3', 'L', 'Trắng', 'Quần kaki', 'kaki', 2, 'ADULT', 'Nhà cung cấp Bình Dương', 1, 200000, 180000, 'Kiểu dáng thời thường'),
('goods-4', 'XL', 'Trắng', 'Áo thun unisex', 'T-shirt', 3, 'ADULT', 'Nhà cung cấp Bình Dương', 1, 100000, 86000, 'Vải Cotton thoáng mát'),
('goods-4', 'XL', 'Xám', 'Áo thun unisex', 'T-shirt', 3, 'ADULT', 'Nhà cung cấp Bình Dương', 1, 100000, 90000, 'Vải Cotton thoáng mát'),
('goods-5', 'S', 'Đỏ', 'Áo thun trẻ em', 'T-shirt', 1, 'KID', 'Nhà cung cấp Bình Dương', 1, 120000, 100000, 'Vải Cotton thoáng mát');



INSERT INTO `goods_in_wh` (`goods_code`, `goods_size`, `goods_color`, `wh_code`, `quantity`, `created_date`, `updated_date`) values
('1', '35', 'Red', 'warehouse-1', 1, '2023-04-01 03:30:52', NULL),
('1', '35', 'Red', 'warehouse-2', 5, '2023-04-01 03:30:52', NULL),
('1', '35', 'Red', 'warehouse-3', 7, '2023-04-01 03:30:52', NULL),
('2', '30', 'Blue', 'warehouse-1', 3, '2023-04-01 03:30:52', NULL),
('goods-1', 'M', 'Đen', 'warehouse-2', 10, '2023-04-01 04:48:37', NULL),
('goods-1', 'S', 'Đen', 'warehouse-2', 12, '2023-04-01 04:47:35', NULL),
('goods-2', 'L', 'Vàng', 'warehouse-3', 20, '2023-04-01 04:48:37', NULL),
('goods-4', 'XL', 'Trắng', 'warehouse-3', 5, '2023-04-01 04:48:38', NULL);