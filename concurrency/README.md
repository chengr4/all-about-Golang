# Concurrency in Golang

## Channels

- Like UNIX pipes (a way of sending data from one process to another)
- Things go in one end, come out the other
- In the same order they went in
- **Multiple readers and writers can share it safely**
- We can not write to a channel unless there is a reader

## Goroutines

- a goroutine is a unit of independent execution

Eg.

```golang
go func() {
    // do something
}()
```

How to stop a goroutine?

1. Use a channel or context to send a signal for completion
2. A well-defined loop terminating condition
3. The process is stopped