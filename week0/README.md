# Week 01 - Functions

## Functions overview

### Prompts the user for input

```python
name = input("What's your name? ")
```

```go
var name string
fmt.Print("What's your name? ")
fmt.Scanln(&name)
```

### Outputs text to the console

```python
print("Hello, world!")
```

```go
fmt.Println("Hello, world!")
```

### Removes leading/trailing whitespace

```python
name = input("What's your name? ").strip()
```

```go
name = strings.TrimSpace(name)
```

### Converts a string to title case

```python
name = name.title()
```

```go
name = strings.Title(name)
```

### Converts a string to uppercase

```python
name = name.upper()
```

```go
name = strings.ToUpper(name)
```

### Converts a string to lowercase

```python
name = name.lower()
```

```go
name = strings.ToLower(name)
```

### Rounds a number

```python
result = round(3.14159, 2)  # 3.14
```

```go
result := math.Round(3.14159*100) / 100  // 3.14
```

### Formats a float to 2 decimal places

```python
price = 3.14159
print(f"{price:.2f}")  # 3.14
```

```go
price := 3.14159
fmt.Printf("%.2f\n", price)  // 3.14
```

# Defining and calling functions

Code sample in Python and Go

```python
def greet(name):
    return f"Hello, {name}!"

def main():
    print(greet("Alice"))

if __name__ == "__main__":
    main()
```

```go
package main

import "fmt"

func greet(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}

func main() {
    fmt.Println(greet("Alice"))
}
```