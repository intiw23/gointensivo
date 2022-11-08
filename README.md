# Go Intensivo


## Steps

### Init repository
``` 
go mod init github.com/intiw23/gointensivo 
```

### internal
- internal resources
- screaming archtecture
```
 . internal/
 . . order/
 . . . entity/ # entity layer
 . . . . order.go 
 . . . infra/ 
 . . . . database/ #database layer
 . . . . . order_repository.go 
 . usecase/ #usecase layer
 . . calculate_price.go
```

### downloading external packages
```
# add you external package in your code and run
go mod tidy
```

###  run all tests
```
go test ./...
```