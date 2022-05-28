# UNIT TEST WITH GO

A simple project with an aritmethic operation (sum) for showing the power of unit testing in Go.

### Running project

```
go run main.go
```

### Running tests with coverage

```
go test -v --coverprofile=coverage.out ./... ./...
```

### Watching the coverage in browser

```
go tool cover -html=coverage.out
```
