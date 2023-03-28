insert into goods values
    ('1', 35, 'Red', 'Hàng để test', 'Hàng test', 3, 'EM BÉ', 'NSX TEST', 1, 10000, 'Hàng để test, xin nhẹ tay'),
    ('2', 30, 'Blue', 'Hàng để trưng bày', 'Hàng trưng bày', 3, 'Mọi lứa tuổi', 'NSX TEST', 1, 200000, 'Hàng để trưng');

insert into goods_in_wh (goods_code, goods_size, goods_color, wh_code, quantity) values
    ('1', 35, 'Red', 'warehouse-1', 1), ('2', 30, 'Blue', 'warehouse-1', 3), ('1', 35, 'Red', 'warehouse-3', 7), ('1', 35, 'Red', 'warehouse-2', 5);