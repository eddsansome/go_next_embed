deploy: build_client build_api

build_client: 
		cd client && npm run build
build_api:
		go build

dev: start_client start_api

start_client:
		cd client && npm run dev &
start_api:
		go run main.go
