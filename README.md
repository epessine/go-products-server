# Go Products Server
---
This is an experiment to create an gRPC server using Golang to serve a "products" listing/creation API. You can create and list products using the `application/grpc/protofiles/product.proto` protofile schema.
___
## Instructions to run
1. Clone repo
2. `docker-compose up -d` to start the container
3. `docker exec -it {app_container_name} go run .` to start the server.
4. `docker exec -it {app_container_name} evans -r repl` to run Evans and interact with API.
5. `docker exec -it {app_container_name} ENV=test go test ./...` to run tests.