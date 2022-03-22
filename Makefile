deploy: build_client build_api

build_client: 
		cd client && npm run build
build_api:
		go build