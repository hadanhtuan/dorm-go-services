STEP TO RUN PROJECT 

1. install protoc https://www.geeksforgeeks.org/how-to-install-protocol-buffers-on-windows/

2. install grpc for golang https://grpc.io/docs/languages/go/quickstart/

3. install nodemon global: npm i nodemon -g

4. swagger:
    go install github.com/swaggo/swag/cmd/swag@latest
    swag init -g ./cmd/main.go -o ./docs
    swag fmt

5. gen proto

	@protoc --go_out=. --go_opt=paths=source_relative \
                --go-grpc_out=. --go-grpc_opt=paths=source_relative \
                proto/common/*.proto

	@protoc --go_out=. --go_opt=paths=source_relative \
                --go-grpc_out=. --go-grpc_opt=paths=source_relative \
                proto/user/*.proto

6. get package
  cd cmd
  go get .

7. run 
	nodemon --exec "go run" ./cmd/main.go --signal SIGTERM
