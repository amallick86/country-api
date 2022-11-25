dcup:
		sudo docker-compose up
dcdb:
		sudo docker-compose up postgres -d
dcdown:
		sudo docker-compose down
drmi:
		sudo docker rmi country-api-api-1
swagger:
		swag init
migcrt:
		goose -dir ./db/migration/ create table_name sql
gooseup:
		goose -dir ./db/migration/ -v postgres "postgres://root:root123@localhost:5432/countrydb?sslmode=disable" up
goosedown:
		goose -dir ./db/migration/ -v postgres "postgres://root:root123@localhost:5432/countrydb?sslmode=disable" down


.PHONY: dcup dcdb dcdown drmi swagger migcrt gooseup goosedown