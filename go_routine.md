# Concurrency and Go Routines in Go

## Introduction
Concurrent programming is a hot topic in Go, especially for those new to the language. This guide will introduce you to Go routines, synchronization, parallelism, and best practices for writing efficient and concurrent applications.

## Go Routines
Go routines are the foundation of concurrency in Go. They allow your application to handle multiple tasks simultaneously. Here's how to create a Go routine:

1. **Creating a Go Routine**:
   - Use the `go` keyword before a function call to run it as a Go routine.
   - Example:
     ```go
     func sayHello() {
         fmt.Println("Hello")
     }

     func main() {
         go sayHello() // Runs sayHello in a Go routine
         time.Sleep(100 * time.Millisecond) // Ensure the main function waits
     }
     ```

2. **Green Threads**:
   - Go routines are lightweight "green threads" managed by the Go runtime.
   - Unlike OS threads, which are heavy (1MB stack size), Go routines start with small stack spaces and are cheap to create and destroy.
   - The Go scheduler maps Go routines onto OS threads, allowing efficient use of CPU cores.

## Synchronization
When working with Go routines, synchronization is crucial to coordinate tasks. Two key tools are:

1. **Wait Groups**:
   - Used to wait for a collection of Go routines to finish.
   - Example:
     ```go
     var wg sync.WaitGroup
     wg.Add(1)
     go func() {
         defer wg.Done()
         sayHello()
     }()
     wg.Wait()
     ```

2. **Mutexes**:
   - Prevent race conditions by ensuring only one Go routine accesses shared data at a time.
   - Example:
     ```go
     var mu sync.Mutex
     var counter int
     go func() {
         mu.Lock()
         counter++
         mu.Unlock()
     }()
     ```

## Parallelism
While concurrency allows handling multiple tasks, parallelism enables executing them simultaneously across multiple CPU cores. To achieve parallelism:
- Ensure your Go application runs on multiple cores (not limited to the Go Playground, which uses only one core).
- Use Go routines and synchronization primitives to distribute work across cores.

## Best Practices
1. **Avoid Race Conditions**:
   - Use mutexes or channels to protect shared resources.
2. **Limit Go Routine Creation**:
   - While Go routines are lightweight, avoid creating excessive numbers unnecessarily.
3. **Use Tools**:
   - Leverage Go's race detector (`-race` flag) to identify and fix concurrency issues.

## Conclusion
Go routines provide a powerful and efficient way to handle concurrency in Go. By understanding synchronization, parallelism, and best practices, you can build highly concurrent and performant applications.
```