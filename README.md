# go_http_api_template
go web api server environment template.

## Quick Start
- build docker image  
$ `docker-compose build`  
- make container of web(go), mysql, phpmyadmin  
$ `docker-compose up -d`  

access
http://localhost:8080/
access phpmyadmin
http://localhost:8080/

## Tips
- You don't need to run `go run main.go` everytime you changed go file. [Air](https://github.com/cosmtrek/air) is running in go_web container
