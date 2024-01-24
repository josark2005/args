# GO ARGS

[![Go Reference](https://pkg.go.dev/badge/github.com/josark2005/args.svg)](https://pkg.go.dev/github.com/josark2005/args)

A simple tool for delivering arguements.

## Usage

### Get Package

```
go get github.com/josark2005/args
```

### Initialize

```go
a := args.NewWarp()
```

### Set Value

```go
a.Set(key, value)
```

### Get Value

```go
a.Get(key)
# or
a.Value(key)
```

### Remove Value

```go
a.Remove(key)
```