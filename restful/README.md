# github.com/davidkhala/goutils/restful


## Swagger Support
To generate swagger docs
1. get binary `swag` by `go install github.com/swaggo/swag/cmd/swag@latest`
2. `swag init -g main/main.go`
    - or use the go generate banner in `main/main.go`   
3. access swagger: if route setup as `/swagger/*any`, go to `/swagger/index.html`