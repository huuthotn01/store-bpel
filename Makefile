hello:
	echo "Hello"

all_start:
	cd goods_service && make start && cd ../../
	cd warehouse_service && make start && cd ../../

gen_new_service:
	go run main.go $(service_name)

build_admin_bff_image:
	docker build -f bff/admin_bff/Dockerfile -t admin-bff:latest .

build_event_service_image:
	docker build -f event_service/Dockerfile -t event-service:latest .

build_customer_service_image:
	docker build -f customer_service/Dockerfile -t customer-service:latest .