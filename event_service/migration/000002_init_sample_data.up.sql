
INSERT INTO `event` (`event_id`, `name`, `discount`, `start_time`, `end_time`, `created_at`, `image`) VALUES
('event_1682582427', 'Sale áo mùa hè', 0.1, '2023-04-28 10:00:00', '2023-06-01 10:00:00', '2023-04-27 08:00:27', '/store-bpel/bff/admin_bff/uploads/event_1682582427/1682582433.webp'),
('event_1682582515', 'Mùa hè năng động', 0.08, '2023-05-16 10:00:00', '2023-06-01 10:00:00', '2023-04-27 08:01:55', '/store-bpel/bff/admin_bff/uploads/event_1682582515/1682582570.webp'),
('event_1682582974', 'Sale phụ kiện nữ', 0.12, '2023-05-03 10:00:00', '2023-05-18 10:00:00', '2023-04-27 08:09:34', '/store-bpel/bff/admin_bff/uploads/event_1682582974/1682582983.webp');

INSERT INTO `goods` (`event_id`, `goods_id`) VALUES
('event_1682582427', 'goods_1682531177'),
('event_1682582427', 'goods_1682531751'),
('event_1682582427', 'goods_1682531953'),
('event_1682582427', 'goods_1682569990'),
('event_1682582427', 'goods_1682570846'),
('event_1682582515', 'goods_1682531001'),
('event_1682582515', 'goods_1682569696'),
('event_1682582515', 'goods_1682570410'),
('event_1682582822', ''),
('event_1682582974', 'goods_1682530531'),
('event_1682582974', 'goods_1682570127');
