start:
	cd main && go run .

migrate:
	migrate create -ext sql -dir migration -seq $(name)

migrate-up:
	migrate -path migration -database "mysql://bpel:bpel@tcp(localhost:3306)/staff_service" -verbose up

migrate-down:
	migrate -path migration -database "mysql://bpel:bpel@tcp(localhost:3306)/staff_service" -verbose down