# gotensor

[![Build status](https://github.com/gotensor/gotensor/workflows/test/badge.svg)](https://github.com/gotensor/gotensor)
[![go.dev reference](https://pkg.go.dev/badge/github.com/gotensor/gotensor)](https://pkg.go.dev/github.com/gotensor/gotensor)
[![Go Report Card](https://goreportcard.com/badge/github.com/gotensor/gotensor)](https://goreportcard.com/report/github.com/gotensor/gotensor)


tensor package for golang

## Install

```
go get -u github.com/gotensor/gotensor
```

## How to use

Creating a new tensor:

```go
data := []interface{}{1, 2, 3, 4, 5, 6}
shape := []int{2, 3}
tx := tensor.NewTensor(data, shape)

fmt.Println(tx)
// Shape: [2 3]
// 1 2 3
// 4 5 6

```

Creating a tensor with all zeros:

```go
shape := []int{3, 3}
tx := tensor.Zeros(shape)

fmt.Println(tx)
// Shape: [3 3]
// 0 0 0
// 0 0 0
// 0 0 0
```

Creating a tensor with all ones:

```go
shape := []int{2, 2}
tx := tensor.Ones(shape)

fmt.Println(tx)
// Shape: [2 2]
// 1 1
// 1 1
```

Getting a single value from the tensor:

```go
data := []interface{}{1, 2, 3, 4, 5, 6}
shape := []int{2, 3}
tx := tensor.NewTensor(data, shape)
value, err := tx.Get(0, 1)
if err != nil {
    fmt.Println(err)
}
fmt.Println(value) // output: 2
```

Setting a single value in the tensor:

```go
data := []interface{}{1, 2, 3, 4, 5, 6}
shape := []int{2, 3}
tx := tensor.NewTensor(data, shape)
err := tx.Set(7, 1, 2)
if err != nil {
    fmt.Println(err)
}
fmt.Println(tx)
// Shape: [2 3]
// 1 2 3
// 4 5 7

```

## License

MIT
