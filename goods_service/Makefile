start:
	cd main && go run .

build:
	cd main && go build

migrate:
	migrate create -ext sql -dir migration -seq $(name)

migrate-up:
	migrate -path migration -database "mysql://bpel:bpel@tcp(localhost:3306)/goods_service" -verbose up

migrate-down:
	migrate -path migration -database "mysql://bpel:bpel@tcp(localhost:3306)/goods_service" -verbose down