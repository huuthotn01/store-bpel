hello:
	echo "Hello"

all_start:
	cd goods_service && make start && cd ../../
	cd warehouse_service && make start && cd ../../

gen_new_service:
	go run main.go $(service_name)

build_admin_bff_image:
	docker build -f bff/admin_bff/Dockerfile -t admin-bff:latest .

build_customer_bff_image:
	docker build -f bff/customer_bff/Dockerfile -t customer-bff:latest .

build_shared_bff_image:
	docker build -f bff/shared_bff/Dockerfile -t shared-bff:latest .

build_account_service_image:
	docker build -f account_service/Dockerfile -t account-service:latest .

build_branch_service_image:
	docker build -f branch_service/Dockerfile -t branch-service:latest .

build_cart_service_image:
	docker build -f cart_service/Dockerfile -t cart-service:latest .

build_customer_service_image:
	docker build -f customer_service/Dockerfile -t customer-service:latest .

build_event_service_image:
	docker build -f event_service/Dockerfile -t event-service:latest .

build_goods_service_image:
	docker build -f goods_service/Dockerfile -t goods-service:latest .

build_order_service_image:
	docker build -f order_service/Dockerfile -t order-service:latest .

build_staff_service_image:
	docker build -f staff_service/Dockerfile -t staff-service:latest .

build_statistic_service_image:
	docker build -f statistic_service/Dockerfile -t statistic-service:latest .

build_warehouse_service_image:
	docker build -f warehouse_service/Dockerfile -t warehouse-service:latest .
