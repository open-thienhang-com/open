1. AUTHO GEN DOCCUMENT
Install swag to generate swagger automatically
https://github.com/swaggo/swag/releases

MACOS: 
./swag init -g pkg/controller/\*

UBUNTU: 
./cmd/dev/swagu init -g pkg/controller/*
  Link: http://localhost:8000/api/v1/index.html
./swagu init -g ../../pkg/controller/*
./swagu init -g pkg/controller/*
2. RUN PROJECT:
go run cmd/dev/main.go 
./swag init -g pkg/controller/*


./swagu init --parseDependency --parseInternal --parseDepth 1 -g cmd/authen.thienhang.com/main.go

./swagu init --parseInternal --parseDepth 1 -g ../../pkg/controller/*


./swagu init --parseDependency --parseInternal --parseDepth 1 -g main.go

./swagu init -g ../../pkg/controller/*



./swagu init --parseInternal --parseDepth  100 -g main.go

./swagu init --parseInternal --parseDepth 1 -g cmd/authen.thienhang.com/main.go