## **Control Structures in Golang**  

Golang provides control structures such as **if**, **switch**, and **loops** (`for`). Let's break them down with examples.

---

# **1. If-Else Statements**  

### **1.1 Basic If Statement**
```go
if x > 10 {
    fmt.Println("x is greater than 10")
}
```

### **1.2 If-Else Statement**
```go
if x%2 == 0 {
    fmt.Println("x is even")
} else {
    fmt.Println("x is odd")
}
```

### **1.3 If-Else If-Else**
```go
if x < 0 {
    fmt.Println("x is negative")
} else if x == 0 {
    fmt.Println("x is zero")
} else {
    fmt.Println("x is positive")
}
```

### **1.4 If with Short Variable Declaration**
```go
if y := 10; y > 5 {
    fmt.Println("y is greater than 5")
}
```
> The variable `y` is only accessible within the `if` block.

---

# **2. Switch Statements**  

### **2.1 Basic Switch Statement**
```go
switch day := "Monday"; day {
case "Monday":
    fmt.Println("Start of the week!")
case "Friday":
    fmt.Println("Weekend is coming!")
default:
    fmt.Println("It's a normal day.")
}
```

### **2.2 Multiple Cases in One Condition**
```go
switch day := "Saturday"; day {
case "Saturday", "Sunday":
    fmt.Println("It's the weekend!")
default:
    fmt.Println("It's a weekday.")
}
```

### **2.3 Switch Without an Expression**
```go
num := 15
switch {
case num < 0:
    fmt.Println("Negative number")
case num == 0:
    fmt.Println("Zero")
case num > 0:
    fmt.Println("Positive number")
}
```
> This works like an **if-else if-else** chain.

### **2.4 Fallthrough Keyword**
```go
switch num := 2; num {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
    fallthrough
case 3:
    fmt.Println("Three") // Executed because of fallthrough
default:
    fmt.Println("Other number")
}
```

---

# **3. Loops in Golang**  

## **3.1 For Loop (Basic)**
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

## **3.2 For Loop Without Initialization and Post Statement**
```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

## **3.3 Infinite Loop**
```go
for {
    fmt.Println("This will run forever")
    break // Prevents infinite loop
}
```

## **3.4 For Loop with Range**
```go
numbers := []int{10, 20, 30}
for index, value := range numbers {
    fmt.Println("Index:", index, "Value:", value)
}
```
> Use `_` to ignore the index:
```go
for _, value := range numbers {
    fmt.Println(value)
}
```

## **3.5 Looping Over a Map**
```go
data := map[string]int{"Alice": 25, "Bob": 30}
for key, value := range data {
    fmt.Println(key, "is", value, "years old")
}
```

## **3.6 Looping Over a String (Rune by Rune)**
```go
for i, char := range "hello" {
    fmt.Printf("Index: %d, Character: %c\n", i, char)
}
```

---

# **4. Break, Continue, and Goto**  

### **4.1 Break Statement**
```go
for i := 1; i <= 5; i++ {
    if i == 3 {
        break // Stops the loop
    }
    fmt.Println(i)
}
// Output: 1 2
```

### **4.2 Continue Statement**
```go
for i := 1; i <= 5; i++ {
    if i == 3 {
        continue // Skips this iteration
    }
    fmt.Println(i)
}
// Output: 1 2 4 5
```

### **4.3 Goto Statement**
```go
i := 1
for i <= 5 {
    if i == 3 {
        goto skip // Jumps to the label
    }
    fmt.Println(i)
    i++
}
skip:
fmt.Println("Skipped part of the loop")
```

---

# **Summary Table**  

| Control Structure | Description | Example |
|------------------|-------------|---------|
| **If-Else** | Conditional branching | `if x > 5 { fmt.Println(x) }` |
| **Switch** | Multi-case selection | `switch day { case "Monday": fmt.Println("Start") }` |
| **For Loop** | Iteration | `for i := 0; i < 5; i++ { fmt.Println(i) }` |
| **Range** | Iterate over collections | `for k, v := range map {}` |
| **Break** | Exit loop | `if i == 3 { break }` |
| **Continue** | Skip iteration | `if i == 3 { continue }` |
| **Goto** | Jump to a label | `goto skip` |