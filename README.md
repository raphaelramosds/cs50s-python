# CS50's Introduction to Programming with Python

Each problem set is implemented in **Python** and/or **Go** as a way to review Python and learn Go side-by-side.

- Python tests use [pytest](https://docs.pytest.org/)
- Go tests use the standard `go test` toolchain

---

## Language Features by Problem Set

### Week 0 — Functions

| Problem | Python | Go |
|---|---|---|
| `0_indoor` | `str.lower()`, type hints, `if __name__ == "__main__"` | `strings.ToLower`, exported functions (capitalized), package layout |
| `1_playback` | `str.replace()`, type hints | `strings.Replace` with `-1` (replace all) |
| `2_making_faces` | Chained method calls, return type annotations, Unicode literals | `strings.ReplaceAll`, intermediate variables, Unicode literals |
| `3_einstein` | Module-level constants, numeric underscores (`90_000_000_000_000_000`), `int` type hints | `const` declaration, explicit unsigned types (`uint`, `uint64`), type casting |
| `4_tip_calculator` | Custom exceptions via class inheritance, `re.fullmatch`, f-string with precision (`:.2f`), string slicing | Custom error types via `Error() string` interface, `regexp.MustCompile`, multiple return values `(T, error)`, `strconv` for type conversion |

### Week 1 — Conditionals

| Problem | Python | Go |
|---|---|---|
| `0_deep_thought` | — | `if`/`else`, logical OR (`\|\|`), string equality |
| `1_bank` | — | `strings.HasPrefix`, `strings.ToLower`, byte indexing (`s[0]`), `if`/`else if`, `strconv.Itoa` |
| `2_extensions` | — | `switch`/`case`/`default`, `regexp.MustCompile`, `FindStringSubmatch`, regex capture groups, `strings.TrimSpace` |
| `3_math_interpreter` | `re.match` with capture groups (`.group()`), `str.strip/lower/replace`, `float()`, floor division (`//`), `ZeroDivisionError`, `ValueError` | — |
| `4_meal_time` | `str.split(":")`, `float()`, numeric range checks (`>=`/`<=`), helper function for unit conversion | — |

### Week 2 — Loops

| Problem | Python | Go |
|---|---|---|
| `0_camel_case` | Character-by-character iteration, `str.isupper()`, conditional expression, string concatenation | — |
| `1_coke_machine` | Module-level list constant (`VALID_COINS`), `str.isnumeric()`, `dict` as mutable accumulator (pass-by-reference semantics), generator expression inside `str.join()`, `while` loop on dict value, testable wrapper (`_static` variant) separating I/O from logic | — |

---

## Python Features Overview

- **Type hints** — parameter and return annotations on all functions:
  ```python
  def indoor(txt: str) -> str:
  ```
- **String methods** — `.lower()`, `.replace()`, `.strip()`, chained:
  ```python
  exp.strip().lower().replace(' ', '')
  ```
- **Numeric literals** — underscores for readability:
  ```python
  C2 = 90_000_000_000_000_000
  ```
- **f-strings** — with format specifiers:
  ```python
  f"Leave ${tip:.2f}"
  ```
- **Regular expressions** — `re.fullmatch` for strict matching, `re.match` with capture groups:
  ```python
  re.fullmatch(r"\$\d{1,3}(\.\d{1,2})?", dollars)
  exp_operands.group(1)
  ```
- **Floor division** — `//` operator for integer-valued results:
  ```python
  return x // y
  ```
- **Built-in exceptions** — raising directly with a message:
  ```python
  raise ZeroDivisionError()
  ```
- **Custom exceptions** — subclassing `ValueError`, calling `super().__init__`:
  ```python
  class InvalidDollarStr(ValueError):
      def __init__(self, txt: str):
          super().__init__(f"dollar: {txt} should be formatted as $##.##")
  ```
- **String splitting** — to unpack structured tokens:
  ```python
  hh, mm = time_str.strip().split(":")
  ```
- **Module guard** — entry point pattern:
  ```python
  if __name__ == "__main__":
  ```
- **String slicing** — strip a leading character:
  ```python
  float(dollars[1:])
  ```
- **`str.isupper()`** — character classification for camelCase → snake_case:
  ```python
  c = '_' + char.lower() if char.isupper() else char
  ```
- **`str.isnumeric()`** — input validation before `int()` conversion:
  ```python
  if not coin.isnumeric():
      return False
  ```
- **`dict` as mutable accumulator** — mutations inside a function are visible to the caller:
  ```python
  invoice["total_inserted"] = invoice["total_inserted"] + coin
  ```
- **Generator expressions** — inline inside `str.join()` to avoid an intermediate list:
  ```python
  ", ".join(str(i) for i in VALID_COINS)
  ```
- **`while` loop** — condition checked against a dict value each iteration:
  ```python
  while invoice["total_inserted"] < COKE_PRICE:
  ```
- **Testability pattern** — `_static` variant accepts a list instead of reading from stdin:
  ```python
  def coke_machine_static(inputs: list) -> dict:  # testable
  def coke_machine() -> None:                     # interactive
  ```

## Go Features Overview

- **Package organization** — each problem in its own package with exported (capitalized) functions:
  ```go
  package tip
  func DollarsToFloat(dollars string) (float32, error)
  ```
- **Multiple return values** — `(T, error)` pattern for fallible operations:
  ```go
  func DollarsToFloat(dollars string) (float32, error) { ... }
  dollarsFloat, err := DollarsToFloat(dollars)
  ```
- **Custom error types** — named types implementing the `error` interface via `Error() string`:
  ```go
  type InvalidDollarStr string
  func (txt InvalidDollarStr) Error() string {
      return fmt.Sprintf("dollar: %s should be formatted as $##.##", string(txt))
  }
  ```
- **`strings` package** — `ToLower`, `HasPrefix`, `TrimSpace`, `ReplaceAll`:
  ```go
  strings.ToLower(greeting)
  strings.HasPrefix(greeting, "hello")
  strings.ReplaceAll(dollars, "$", "")
  ```
- **`regexp` package** — `MustCompile`, `MatchString`, `FindStringSubmatch` with capture groups:
  ```go
  re := regexp.MustCompile(`.*\.(\w+)$`)
  match := re.FindStringSubmatch(filename)  // match[1] = captured extension
  ```
- **`strconv` package** — `ParseFloat`, `Itoa`, `FormatFloat` for string↔numeric conversion:
  ```go
  strconv.ParseFloat(dollarsStrParsed, 32)
  strconv.Itoa(amount)
  strconv.FormatFloat(float64(tip), 'f', 2, 32)
  ```
- **`const`** — typed and untyped constant declarations:
  ```go
  const C2 = 90000000000000000
  ```
- **Explicit numeric types** — `uint`, `uint64`, `float32`, `float64` with explicit casts:
  ```go
  func Einstein(mass uint) uint64 {
      return uint64(mass) * C2
  }
  ```
- **`switch`/`case`** — multi-branch dispatch with a `default` fallback:
  ```go
  switch match[1] {
  case "jpg", "jpeg": return "image/jpeg"
  case "png":         return "image/png"
  default:            return "application/octet-stream"
  }
  ```
- **Byte indexing** — `s[0]` to inspect the first byte of a string:
  ```go
  greeting[0] == 'h'
  ```
- **Testing** — table-driven style using `*testing.T` and `t.Errorf`
