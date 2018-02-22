---
title: "Static code analysis using Go vet"
date: 2017-08-23
lastmod: 2018-01-02
description: "A look at the kind of problems go's static analysis tool can detect in your code."
draft: false
tags: ["golang", "go", "vet", "static"]
categories: ["software"]
---

If you've been using Go for a while you might have heard about `vet`.
To vet means to make a careful and critical examination of something. This is
a pretty good name for what the tool does.

Vet is a pretty interesting tool. For people familiar with `gcc` based tools, it
can be compared to the warnings `gcc` emits. It inspects source files and outputs
recommendation or "warnings".

# Invoking vet

You can technically use `go vet`, but its interface is awkward for most.
It only works with packages names, which you can obtain through `go list`.

Instead you can use `go tool vet`. This version is simpler to work with, allowing
you to work with directories or files (as long as they are in a single package).
It you can also list the optional flags using `go tool vet --help`.

Here's a list of what the tool can do:

## Locks are not passed by value (`-copylocks`) - A favourite.
```golang

// Function with a mutex argument
func GoodLocker(l *sync.Mutex) {}
func BadLocker(l sync.Mutex) {} // passes lock by value: sync.Mutex

// Struct with a mutex field
type Locker1 struct {
        l sync.RWMutex
}
func (l *Locker1) GoodMethod() {}
func (l Locker1) BadMethod() {} // passes lock by value: Locker1 contains sync.RWMutex

func main() {
        var l sync.Mutex
        GoodLocker(&l)
        BadLocker(l) // call of BadLocker copies lock value: sync.Mutex
}
```

## Useless assignment (`-assign`)
```golang
func main() {
        s := "hey"
        s = s // self-assignment of s to s
}
```

## Mistaken sync/atomic usage (`-atomic`)
```golang
func main() {
        n := uint64(0)
        n = atomic.AddUint64(&n, 1) // direct assignment to atomic value
}
```

## Mistakes involving boolean operators (`-bool`)
```golang
func main() {
        var i, j int
        _ = i != 10 || i != 2           // suspect or: i != 10 || i != 2
        _ = i == 1 || 1 == i || j == 0  // redundant or: i == 1 || 1 == i
}
```

## Not using key to identify struct fields (`-composites`)
```golang
import "container/list"
func main() {
        _ = list.Element{
                "First", // container/list.Element composite literal uses unkeyed fields
        }
        _ = list.Element{
                Value: "First",
        }
}
```

## Nil function comparison (`-nilfunc`)
```golang
func F() {}
var f = F

func main() {
        if f == nil {
                return 
        }
        if F == nil { // comparison of function F == nil is always false

        }
}
```

## Printf invocation error (`-printf`)
```golang
func main() {
        fmt.Printf("%s", "hello")
        fmt.Printf("%d", "hello") // arg "hello" for printf verb %d of wrong type: string
}
```

## Unsafe concurrent access of range variables (`-rangeloops`)
```golang
func main() {
        s := []int{11, 22, 33, 44, 55, 66}
        for i := range s {
                go func() {
                        fmt.Printf("bad %d\n", i) // range variable i captured by func literal
                }
        }
        for i := range s {
                // Correct usage
                go func(i int) {
                        fmt.Printf("good %d\n", i)
                }(i)
        }
}
```

## Variable shadowing (`-shadow`). Not active by default
```golang
func example(w io.Writer, buf bytes.Buffer) (err error) {
        if buf.Len() > 0 {
                _, err := buf.ReadByte()
                if err != nil {
                        return err // declaration of "err" shadows declaration...
                }
        }
        return err
}
```

## Improper usage of struct field tags (`-structtags`)
```golang
type Struct struct {
        A string `json:"a"`
        B int `json:"a"` // struct field B repeats json tag "a" also at...
        C bool `"bad"`   // sruct field tag `"bad"` not compatible with reflect.StructTag.Get:
                         // bad syntax for struct tag key
}
```

## Unreachable code, A.K.A. dead code (`-unreachable`)
```golang
func simpleMergeError() int {
        a := "hello"
        return 0
        fmt.Println("unreachable") // unreachable code
        return 1
}
func anotherSimpleMergeError(a int) {
        switch a {
        case 1:
                return
        default:
                return
        }
        fmt.Println("unreachable") // unreachable code
}
func lessObviousCase() {
        for {
                for {
                        break
                }
        }
        fmt.Println("unreachable") // unreachable code
}
```
