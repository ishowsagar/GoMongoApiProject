# as we gonna use .env var --> always include .env file to import those
include .env

# go build -o OutputWhere source-buildingFromWhere - entry point
build :
	go build -o ${BINARY} ./cmd/api/main.go	

# to run project serve from entry point binary --> just directly pass the path and it runs the project
run :
	./${BINARY}

reboot : build run

#! Docker compose commands alias
up :
	@echo "starting MongoDB containers⚡..."
	docker-compose up --build -d --remove-orphans
down :
	@echo "shutting down MongoDB containers❌..."
	docker-compose down 

	
# connection string
# mongodb://admin:password@localhost:27017/todos_db?authSource=admin&readPreference=primary&appname=MongDB%20Compass&directConnection=true&ssl=false


#! Git commands
add :
	git add .
commit : add
	git commit -m "Added changed and the files are updated to the latest code"
stash : commit
push : stash
	git push 

