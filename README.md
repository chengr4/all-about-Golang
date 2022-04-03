# All About Golang

[first-try](./tutorials/first-try): This folder includes some try of Golang

[Golang types](./go-types)

## TOLEARN NEXT

- `[]byte()`
- struct tag

> https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go

## Module Management

- `go mod init <name with path>`: init project
- `go build`
- `go list -m all`: 查看 dependency
- `go get`: add the new dependency
- `go mod tidy`: 自動刪除未使用的 dependencies

## Attrbiute

+ Strong and statically typed
+ Excellent communnity
+ Simpicity
+ Fast compile times
+ Garbage collected: Do not need to focus on manage memory
+ built-in testing support
+ Object-oriented (but in it's own way)
+ Go is a **pass by value** language
  + Go makes "copies" of values when passed into functions
  > See `first-try/pass-by-value.go`

## types

> [GOTO go types](./go-types)

| Non-Pointer Values | Pointer Wrapper Values |
| ------------------ | ---------------------- |
| string             | slice                  |
| int                | map                    |
| float              | function               |
| boolean            |                        |
| **array**          |                        |
| **struct**         |                        |

## FAQ

### 1. Why a New Language?

At the moment just before `Golang` was created:

+ `Python`: Easy to use but slow (interpreted language)
+ `Java`: Increasingly complex type system
+ `C/C++`: Complex type system and Slow compile times (E.g. `C++` applications written 10 years age still need to compile today)

### How to run a file

`go run <name of the file>`

### What is `go.mod`

When your code imports packages contained in other modules, you manage those dependencies through your code's own module. That module is defined by a go.mod file that tracks the modules that provide those packages. That go.mod file stays with your code, including in your source code repositor

```
go mod init example.com/m
```

## Package types

### main (executable)

```
go run *.go
go build -o exec
```

### non main (non executable)

`import "github.com/xxxx/xxxxx"`

## References

1. [Learn Go Programming - Golang Tutorial for Beginners (2019.6)](https://youtu.be/YS4e4q9oBaU)
2. [The Net Ninja; Go Tutorial (Golang) for Beginners (2021)](https://www.youtube.com/playlist?list=PL4cUxeGkcC9gC88BEo9czgyS72A3doDeM)
3. [Standard library](https://pkg.go.dev/std)

### High Concurrency

1. https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html?fbclid=IwAR0Qs3jYR4dJBkINCn6BW5lPizwWNPBZRI_8Q6nW27uUdD3M-mdjr3pWKzE

> ps. 高效能 backend system 的很好導讀文
