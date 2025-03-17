## **Arrays, Slices, and Maps in Golang**  

Golang provides **arrays**, **slices**, and **maps** to handle collections of data efficiently. Let's explore them in detail.

---

# **1. Arrays in Golang**  
An **array** is a fixed-size collection of elements of the same type.

### **1.1 Declaring and Initializing an Array**
```go
package main

import "fmt"

func main() {
    // Declare an array of size 3
    var arr [3]int
    fmt.Println(arr) // Output: [0 0 0] (default values)

    // Assign values
    arr[0] = 10
    arr[1] = 20
    arr[2] = 30
    fmt.Println(arr) // Output: [10 20 30]

    // Initialize an array with values
    arr2 := [3]int{5, 10, 15}
    fmt.Println(arr2) // Output: [5 10 15]

    // Let Go determine the size automatically
    arr3 := [...]int{1, 2, 3, 4, 5}
    fmt.Println(arr3) // Output: [1 2 3 4 5]
}
```

---

### **1.2 Iterating Over an Array**
```go
for i, v := range arr3 {
    fmt.Println("Index:", i, "Value:", v)
}
```

---

### **1.3 Multidimensional Arrays**
```go
var matrix [2][3]int // 2 rows, 3 columns
matrix[0][1] = 7
fmt.Println(matrix) // Output: [[0 7 0] [0 0 0]]
```

---

# **2. Slices in Golang**  
A **slice** is a dynamically-sized, flexible view of an array.

### **2.1 Declaring and Initializing a Slice**
```go
// Create a slice from an array
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4] // Slices elements from index 1 to 3 (excluding 4)
fmt.Println(slice) // Output: [2 3 4]

// Declare a slice directly
slice2 := []int{10, 20, 30}
fmt.Println(slice2) // Output: [10 20 30]

// Create an empty slice with make
slice3 := make([]int, 5) // Initializes with 5 zero values
fmt.Println(slice3)      // Output: [0 0 0 0 0]
```

---

### **2.2 Appending to a Slice**
```go
slice := []int{1, 2, 3}
slice = append(slice, 4, 5)
fmt.Println(slice) // Output: [1 2 3 4 5]
```

---

### **2.3 Copying a Slice**
```go
src := []int{10, 20, 30}
dest := make([]int, len(src))
copy(dest, src)
fmt.Println(dest) // Output: [10 20 30]
```

---

### **2.4 Slices Are References to Arrays**
Modifying a slice affects the underlying array:

```go
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4]
slice[0] = 99
fmt.Println(arr)   // Output: [1 99 3 4 5]
fmt.Println(slice) // Output: [99 3 4]
```

---

# **3. Maps in Golang**  
A **map** is a collection of key-value pairs, similar to a dictionary.

### **3.1 Declaring and Initializing a Map**
```go
// Declare a map
var myMap map[string]int

// Initialize a map
myMap = make(map[string]int)

// Declare and initialize a map
scores := map[string]int{
    "Alice": 90,
    "Bob":   85,
    "Eve":   92,
}

// Add key-value pairs
scores["Charlie"] = 88

fmt.Println(scores) // Output: map[Alice:90 Bob:85 Eve:92 Charlie:88]
```

---

### **3.2 Accessing and Checking Keys**
```go
val := scores["Alice"]
fmt.Println(val) // Output: 90

// Check if a key exists
if val, exists := scores["John"]; exists {
    fmt.Println("John's score:", val)
} else {
    fmt.Println("John not found")
}
```

---

### **3.3 Deleting a Key**
```go
delete(scores, "Bob")
fmt.Println(scores) // Output: map[Alice:90 Eve:92 Charlie:88]
```

---

### **3.4 Iterating Over a Map**
```go
for key, value := range scores {
    fmt.Println("Name:", key, "Score:", value)
}
```

---

# **4. Differences Between Arrays, Slices, and Maps**
| Feature  | Arrays  | Slices  | Maps  |
|----------|---------|---------|---------|
| Size     | Fixed   | Dynamic | Dynamic |
| Mutability | Immutable (size) | Mutable (size) | Mutable (keys & values) |
| Indexing | Numeric | Numeric | Key-based |
| Use Case | Fixed-size data | Dynamic lists | Key-value storage |

---

✅ Use **arrays** when size is fixed.  
✅ Use **slices** for flexible collections.  
✅ Use **maps** for key-value lookups.