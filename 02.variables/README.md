# Go Variable Declarations

1. **Declare variables using `var`**

    **1.1. Declare with an explicit type**
    ```go
    var x int
    var name string
    ```
    > Variables get their zero values (0, "", false, etc.).

    **1.2. Declare and initialize**
    ```go
    var age int = 25
    var msg string = "Hello"
    ```

    **1.3. Let Go infer the type**
    ```go
    var pi = 3.14    // float64
    var ok = true    // bool
    ```

    **1.4. Declare multiple variables**
    ```go
    var a, b, c int
    var name, age = "Alice", 20
    ```

2. **Short variable declaration using `:=`**
    ```go
    x := 10
    msg := "hello"
    a, b := 1, 2
    ```
    > Go automatically infers the type.  
    > Cannot be used at the package level.

3. **Multi-variable block declaration `var (...)`**

    *Useful for grouping variables:*
    ```go
    var (
        name string = "Bob"
        age  int    = 30
        ok   bool
    )
    ```

4. **Constant declaration using `const`**

    *(Not a variable, but often grouped with declarations.)*
    ```go
    const PI = 3.14
    const SIZE int = 100
    ```

    **Multiple constants:**
    ```go
    const (
        A = 1
        B = 2
    )
    ```

5. **Blank identifier (`_`)**

    *Used to discard unwanted values:*
    ```go
    _, err := someFunction()
    ```

6. **Package-level (global) variables**

    *Can only use `var`, not `:=`:*
    ```go
    package main

    var Version = "1.0"
    var debug bool
    ```
