 

GOCMD=go
MYSQL_PASS=--mysqldb-pass
MYSQL_HOST=--mysqldb-host
MYSQL_USER=--mysqldb-user
MYSQL_PORT=--mysqldb-addr
MYSQL_DB_NAME=--mysqldb-addr
GOBUILD=$(GOCMD) build


BINARY_NAME=user-service
SERVICE_NAME=user
BUILD_PATH=./bin/$(BINARY_NAME)

build:
	$(GOBUILD) -o bin/$(BINARY_NAME) cmd/main.go
   
run:
	$(BUILD_PATH) $(MYSQL_HOST) localhost $(MYSQL_USER) root $(MYSQL_PASS) password $(MYSQL_PORT)=3306 $(MYSQL_DB_NAME)=leaves_module
            
migrations-up : 
	./migrate -path migrations/schema/ -database mysql://${MYSQL_DB_USER}:${MYSQL_DB_PASS}@/${MYSQL_DB_NAME} up