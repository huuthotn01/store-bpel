hello:
	echo "Hello"

all_start:
	cd goods_service && make start && cd ../../
	cd warehouse_service && make start && cd ../../

gen_new_service:
	go run main.go $(service_name)