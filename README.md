# GO ARGS

A simple tool for delivering arguements.

## Usage

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