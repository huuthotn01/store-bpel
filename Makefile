hello:
	echo "Hello"

all_start:
	cd goods_service && make start && cd ../../
	cd warehouse_service && make start && cd ../../

gen_new_service:
	go run main.go $(service_name)

build_base_docker_image:
	docker build -t base-image:latest .

build_admin_bff_image:
	docker build -f bff/admin_bff/Dockerfile -t admin-bff:latest .