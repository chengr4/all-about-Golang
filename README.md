# All About Golang

[first-try](./first-try): This folder includes some try of Golang

[types](./types)

## TOLEARN NEXT

+ `[]byte()`

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

## References

1. [Learn Go Programming - Golang Tutorial for Beginners (2019.6)](https://youtu.be/YS4e4q9oBaU)
2. [The Net Ninja; Go Tutorial (Golang) for Beginners (2021)](https://www.youtube.com/playlist?list=PL4cUxeGkcC9gC88BEo9czgyS72A3doDeM)
3. [Standard library](https://pkg.go.dev/std)
