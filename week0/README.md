# Week 01 - Functions

## Functions overview

| Description | Python | Go |
|---|---|---|
| Prompts the user for input | `input()` | `fmt.Scanln()` |
| Outputs text to the console | `print()` | `fmt.Println()` |
| Removes leading/trailing whitespace | `str.strip()` | `strings.TrimSpace()` |
| Converts a string to title case | `str.title()` | `strings.Title()` |
| Converts a string to uppercase | `str.upper()` | `strings.ToUpper()` |
| Converts a string to lowercase | `str.lower()` | `strings.ToLower()` |
| Rounds a number | `round(number, ndigits)` | `math.Round()` |
| Formats a float to 2 decimal places | `print(f"{variable:.2f}")` | `fmt.Printf("%.2f", variable)` |

# Defining and calling functions

Code sample in Python:

```python
def greet(name):
    return f"Hello, {name}!"

def main():
    print(greet("Alice"))

if __name__ == "__main__":
    main()
```

Code sample in Go:

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