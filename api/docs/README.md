работа со сваггер
export PATH=$(go env GOPATH)/bin:$PATH
swag init -d ./cmd/,./ -o ./api/docs --parseDependency