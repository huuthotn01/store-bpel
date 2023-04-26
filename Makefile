hello:
	echo "Hello"

gen_new_service:
	go run main.go $(service_name)

build_admin_bff_image:
	docker build -f bff/admin_bff/Dockerfile -t admin-bff:$(shell date '+%Y%m%d%H%M%S') .

build_customer_bff_image:
	docker build -f bff/customer_bff/Dockerfile -t customer-bff:$(shell date '+%Y%m%d%H%M%S') .

build_shared_bff_image:
	docker build -f bff/shared_bff/Dockerfile -t shared-bff:$(shell date '+%Y%m%d%H%M%S') .

build_account_service_image:
	docker build -f account_service/Dockerfile -t account-service:$(shell date '+%Y%m%d%H%M%S') .

build_branch_service_image:
	docker build -f branch_service/Dockerfile -t branch-service:$(shell date '+%Y%m%d%H%M%S') .

build_cart_service_image:
	docker build -f cart_service/Dockerfile -t cart-service:$(shell date '+%Y%m%d%H%M%S') .

build_customer_service_image:
	docker build -f customer_service/Dockerfile -t customer-service:$(shell date '+%Y%m%d%H%M%S') .

build_event_service_image:
	docker build -f event_service/Dockerfile -t event-service:$(shell date '+%Y%m%d%H%M%S') .

build_goods_service_image:
	docker build -f goods_service/Dockerfile -t goods-service:$(shell date '+%Y%m%d%H%M%S') .

build_order_service_image:
	docker build -f order_service/Dockerfile -t order-service:$(shell date '+%Y%m%d%H%M%S') .

build_staff_service_image:
	docker build -f staff_service/Dockerfile -t staff-service:$(shell date '+%Y%m%d%H%M%S') .

build_statistic_service_image:
	docker build -f statistic_service/Dockerfile -t statistic-service:$(shell date '+%Y%m%d%H%M%S') .

build_warehouse_service_image:
	docker build -f warehouse_service/Dockerfile -t warehouse-service:$(shell date '+%Y%m%d%H%M%S') .
