# Lesson 4: Channels and Stopping Goroutines

## Objective

In this lesson, you will learn how to:

- Communicate between goroutines using channels  
- Control and stop running goroutines safely  
- Build a simple concurrent system that can shut down cleanly  

---

## Introduction

In the previous lesson, you created multiple goroutines that ran at the same time.

Now, we will:

1. Let goroutines communicate
2. Learn how to stop them safely

These are essential patterns used in real-world Go systems.

---

## Channels

A channel allows goroutines to send and receive data.

Think of a channel as a pipe:

- One goroutine sends data into the pipe  
- Another goroutine reads data from it  

---

## Example: Basic Channel

```go
package main

import "fmt"

func main() {
    messageChannel := make(chan string)

    go func() {
        messageChannel <- "Hello from a goroutine!"
    }()

    msg := <-messageChannel
    fmt.Println(msg)
}
```

### Explanation

- make(chan string) creates a channel
- channel <- value sends data
- <-channel receives data

The program waits until a value is received.

---

## Coin Watcher with Channels

Instead of printing directly, goroutines will send messages to main.

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func watchCoin(name string, ch chan string) {
    for i := 0; i < 5; i++ {
        price := rand.Intn(100)
        ch <- fmt.Sprintf("%s price is $%d", name, price)
        time.Sleep(1 * time.Second)
    }
}

func main() {
    ch := make(chan string)

    go watchCoin("Bitcoin", ch)
    go watchCoin("Ethereum", ch)

    for i := 0; i < 10; i++ {
        fmt.Println(<-ch)
    }
}
```

### Key Idea

- Goroutines produce data
- The main function consumes data

---

## Stopping Goroutines

So far, goroutines stop only when their loop ends.

In real applications, we often need to stop them early.

---

## Using a Stop Channel

We can send a signal to tell goroutines to stop.

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func watchCoin(name string, ch chan string, stop chan bool) {
    for {
        select {
        case <-stop:
            fmt.Println(name, "stopping...")
            return
        default:
            price := rand.Intn(100)
            ch <- fmt.Sprintf("%s price is $%d", name, price)
            time.Sleep(1 * time.Second)
        }
    }
}

func main() {
    ch := make(chan string)
    stop := make(chan bool)

    go watchCoin("Bitcoin", ch, stop)
    go watchCoin("Ethereum", ch, stop)

    go func() {
        time.Sleep(5 * time.Second)
        stop <- true
        stop <- true
    }()

    for i := 0; i < 10; i++ {
        fmt.Println(<-ch)
    }
}
```

---

## Explanation

- stop is a signal channel  
- select lets the goroutine wait for multiple events  
- When a stop signal is received, the goroutine exits  

---

## Important Concepts

- Channels enable safe communication between goroutines  
- select allows waiting on multiple channels  
- Goroutines should always have a way to stop  

---

## Exercises

1. Add another coin watcher  
2. Change one watcher to update every 500ms  
3. Print a custom message when a goroutine stops  
4. Stop the program after 3 seconds instead of 5  

---

## Challenge

Can you stop all goroutines using only one signal?

---

## Summary

In this lesson, you learned:

- How to use channels for communication  
- How to coordinate goroutines  
- How to stop concurrent processes safely  
