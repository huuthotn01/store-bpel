INSERT INTO `cart` (`cart_id`, `customer_id`) VALUES
(1, 'user1'),
(2, 'user2');


INSERT INTO `goods` (`cart_id`, `goods_id`, `goods_size`, `goods_color`, `quantity`) VALUES
('1', 'goods-1', 'M', 'Đen', 1),
('1', 'goods-1', 'S', 'Đen', 1),
('1', 'goods-2', 'XL', 'Xanh lá', 1),
('2', 'goods-2', 'M', 'Trắng', 1),
('2', 'goods-3', 'L', 'Xám', 3);