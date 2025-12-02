1. **Classic for Loop**
>Used when you know the number of iterations.
>Initialization
>Condition
>Post statement
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

2. **While-like Loop**
>Go does not have a while keyword, but you can achieve the same behavior:
```go
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}
```

3. **Infinite Loop**
>Used when running tasks continuously until an explicit break.
```go
for {
    fmt.Println("Running forever...")
}
```

4. **Loop Control Keywords**
>Go provides:
>break — exit the loop immediately
>continue — skip to the next iteration

```go
for i := 0; i < 5; i++ {
    if i == 2 {
        continue
    }
    if i == 4 {
        break
    }
    fmt.Println(i)
}
```

6. **Looping Over Collections with range**
>range is used to iterate over arrays, slices, strings, maps, and channels.

```go
numbers := []int{1, 2, 3}

for index, value := range numbers {
    fmt.Println(index, value)
}
```
>When you don't need the index or value, you can use _:

```go
for _, v := range numbers {
    fmt.Println(v)
}

```
7. **Loop Labels**
>Used to break or continue outer loops.

```go
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if j == 1 {
            break outer
        }
    }
}

```
8. **Important Notes**

>Go encourages readability and avoids complex loop constructs.
>The for loop is versatile enough to replace all other loop types.
>Avoid heavy use of labels unless necessary.