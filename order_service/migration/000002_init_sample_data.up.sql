insert into order_state(order_code, state) values
    (1, 'NHÀ SẢN XUẤT CHUẨN BỊ HÀNG'),
    (2, 'NHÀ SẢN XUẤT CHUẨN BỊ HÀNG'),
    (1, 'NHẬP KHO BÌNH THẠNH') ;


INSERT INTO `goods` (`goods_code`, `goods_size`, `goods_color`, `goods_name`, `order_code`, `quantity`, `unit_price`, `total_price`, `tax`, `image`, `promotion`) VALUES
('goods_1681248800', '31', 'Đen ', '', 1, 2, 80000, 72000, 1000, '', 10),
('goods_1681248800', '31', 'Đen ', '', 2, 1, 80000, 72000, 1000, '', 10),
('goods_1681248800', '32', 'Đen ', '', 1, 1, 80000, 72000, 1000, '', 10),
('goods_1681248800', '32', 'Đen ', '', 2, 1, 80000, 72000, 1000, '', 10),
('goods_1681248884', '36', 'Đỏ xanh', '', 3, 1, 100000, 100000, 1000, '', 0),
('goods_1681248884', '36', 'Trắng xanh', '', 3, 1, 100000, 100000, 1000, '', 0),
('goods_1681249061', 'L', 'Đen', '', 4, 1, 185000, 185000, 1000, '', 0),
('goods_1681249061', 'L', 'Hồng', '', 4, 1, 185000, 185000, 1000, '', 0),
('goods_1681249061', 'M', 'Đen', '', 4, 1, 185000, 185000, 1000, '', 0),
('goods_1681249162', 'L', 'Trắng hồng', '', 5, 1, 376000, 376000, 1000, '', 0),
('goods_1681249286', '37', 'Đen', '', 6, 2, 249000, 249000, 1000, '', 0),
('goods_1681249286', '38', 'Đen', '', 7, 1, 249000, 249000, 1000, '', 0),
('goods_1681249286', '38', 'Trắng', '', 7, 1, 249000, 249000, 1000, '', 0),
('goods-1', 'M', 'Đen', '', 8, 1, 250000, 250000, 1000, '', 0),
('goods-2', 'L', 'Vàng', '', 10, 2, 150000, 150000, 1000, '', 0),
('goods-2', 'M', 'Vàng', '', 10, 2, 150000, 150000, 1000, '', 0),
('goods-3', 'L', 'Trắng', '', 11, 4, 200000, 200000, 1000, '', 0),
('goods-4', 'XL', 'Trắng', '', 12, 1, 100000, 100000, 1000, '', 0),
('goods-5', 'S', 'Đỏ', '', 9, 2, 120000, 120000, 1000, '', 0);

INSERT INTO `online_orders` (`order_code`, `expected_delivery`, `shipping_fee`, `customer_id`, `payment_method`, `street`, `ward`, `district`, `province`, `customer_name`, `customer_phone`, `customer_email`, `status`) VALUES
(1, '2023-04-06', 5000, 'user2', 'Momo', 'Nguyen Du', 'Dong Hoa', 'Di An', 'Binh Duong', 'Tran Nguyen Huu Tho', '0123456789', 'test@gmail.com', 0),
(2, '2023-04-06', 5000, 'user2', 'Momo', 'Nguyen Du', 'Dong Hoa', 'Di An', 'Binh Duong', 'Tran Nguyen Huu Tho', '0123456789', 'test@gmail.com', 0),
(3, '2023-04-06', 5000, 'user1', 'Momo', 'Nguyen Du', 'Dong Hoa', 'Di An', 'Binh Duong', 'Tran Nguyen Huu', '0123456789', 'testT@gmail.com', 0),
(4, '2023-04-07', 5000, 'user1', 'Momo', 'Nguyen Du', 'Dong Hoa', 'Di An', 'Binh Duong', 'Tran Nguyen Hu', '0123456789', 'testT1@gmail.com', 0),
(5, '2023-04-07', 5000, 'user1', 'Momo', 'Nguyen Du', 'Dong Hoa', 'Di An', 'Binh Duong', 'Tran Nguyen', '0123456789', 'tes1@gmail.com', 0),
(6, '2023-04-07', 5000, 'user1', 'Momo', 'Nguyen Du', 'Dong Hoa', 'Di An', 'Binh Duong', 'Tran Nguyen', '0123456789', 'tes1@gmail.com', 0),
(7, '2023-04-10', 5000, 'user2', 'Momo', 'Nguyen Du', 'Dong Hoa', 'Di An', 'Binh Duong', 'Tran Nguyen', '0123456789', 'tes1@gmail.com', 0),
(8, '2023-04-10', 5000, 'user1', 'Momo', 'Nguyen Du', 'Dong Hoa', 'Di An', 'Binh Duong', 'Tran N', '0123456789', 't211@gmail.com', 0);


INSERT INTO `orders` (`order_code`, `transaction_date`, `total_price`, `public_order_code`) VALUES
(1, '2023-04-01 17:00:00', 216000, 'nxfbAUZh'),
(2, '2023-03-31 17:00:00', 144000, 'fqNVgCDe'),
(3, '2023-03-31 17:00:00', 200000, 'izLuKevz'),
(4, '2023-04-02 17:00:00', 555000, 'uCYnYnZL'),
(5, '2023-04-03 17:00:00', 376000, 'jTZzdFPk'),
(6, '2023-04-04 17:00:00', 498000, 'tyzkMbbA'),
(7, '2023-04-05 17:00:00', 498000, 'hNSmlsxW'),
(8, '2023-04-06 17:00:00', 250000, 'NKiueqgi'),
(9, '2023-04-07 17:00:00', 240000, 'IxpNHEvR'),
(10, '2023-04-08 17:00:00', 600000, 'lcNKwioP'),
(11, '2023-04-09 17:00:00', 800000, 'EJNrGYpD'),
(12, '2023-04-10 17:00:00', 100000, 'cNvWGctG');

INSERT INTO `store_orders` (`order_code`, `store_code`, `staff_id`) VALUES
(9, '1', '12345'),
(10, '1', '1'),
(11, '1', '1'),
(12, '1', '3');