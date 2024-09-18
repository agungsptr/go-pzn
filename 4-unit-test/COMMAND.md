# Golang Unit Test

All unit test
```sh
go test ./...
```

Using verbose
```sh
go test -v ./...
```

Spesific function to test
```sh
go test -v -run FuncName
```

Spesific only one sub-test of fucntion test
```sh
go test -v -run FuncName/SubTaskName
```

Only run sub-test with name
```sh
go test -v -run /SubTaskName
```
