# Week 0 - Functions

## Running tests

Each problem has a `python/` and a `go/` directory. You can run tests for all problems at once or for a single one.

**All problems (Python + Go):**

```bash
make all
```

**Single problem:**

```bash
make test-4_tip_calculator
```

**Python only (from the problem's `python/` directory):**

```bash
cd 4_tip_calculator/python && pytest
```

**Go only (from the problem's `go/` directory):**

```bash
cd 4_tip_calculator/go && go test ./...
```