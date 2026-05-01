# Category 7: Functions

Learn to write reusable functions.

**Total Points: 250**

---

## 1. Function Basics (5 points)

Define and call a function.

**Expected output:**
```
Hello from function!
```

**Hint:** Use `func name() { }` syntax.

---

## 2. Function with Parameter (5 points)

Pass parameter to function.

**Expected output:**
```
Hello, Alice!
```

**Hint:** `func greet(name string) { }`

---

## 3. Return Value (10 points)

Return a value from function.

**Expected output:**
```
Sum: 15
```

**Hint:** Use `return` statement.

---

## 4. Return Named Value (10 points)

Use named return values.

**Expected output:**
```
Result: 15
```

**Hint:** Declare return type in function signature.

---

## 5. Multiple Return Values (10 points)

Return multiple values.

**Expected output:**
```
Quotient: 2, Remainder: 1
```

**Hint:** `(int, int)` as return type.

---

## 6. Error Return (15 points)

Return error from function.

**Expected output:**
```
Result: 5, Error: nil
Result: 0, Error: division by zero
```

**Hint:** Return `(int, error)`.

---

## 7. Variadic Function (10 points)

Accept variable arguments.

**Expected output:**
```
Sum: 15
Sum: 20
```

**Hint:** Use `...type` for variadic.

---

## 8. Anonymous Function (10 points)

Define function inline.

**Expected output:**
```
Result: 25
```

**Hint:** `func(x int) int { return x * x }()`

---

## 9. Function as Value (10 points)

Assign function to variable.

**Expected output:**
```
Square: 16
```

**Hint:** Functions are values in Go.

---

## 10. Higher-Order Function (15 points)

Pass function as parameter.

**Expected output:**
```
Result: 30
```

**Hint:** Function can accept other functions.

---

## 11. Closure (15 points)

Create closure.

**Expected output:**
```
10
11
12
```

**Hint:** Inner function accesses outer variable.

---

## 12. Recursion (15 points)

Use recursive function.

**Expected output:**
```
Factorial of 5: 120
```

**Hint:** Function calls itself.

---

## 13. Recursion Fibonacci (15 points)

Recursive Fibonacci.

**Expected output:**
```
Fibonacci(10): 55
```

**Hint:** Recursive with base cases.

---

## 14. Defer (10 points)

Use defer statement.

**Expected output:**
```
Started
Finished
Cleanup done
```

**Hint:** `defer` runs after function returns.

---

## 15. Multiple Defer (10 points)

Multiple defer statements.

**Expected output:**
```
1
2
3
(LIFO order)
```

**Hint:** Defer executes in LIFO order.

---

## 16. Callback Pattern (15 points)

Implement callback.

**Expected output:**
```
Transformed: [2 4 6 8 10]
```

**Hint:** Pass transform function.

---

## 17. Builder Pattern (15 points)

Use functional options pattern.

**Expected output:**
```
Name: Alice, Age: 25
```

**Hint:** Return struct from function.

---

## 18. Error Handling with Defer (15 points)

Use defer for cleanup on error.

**Expected output:**
```
Error occurred: file not found
Cleanup done
```

**Hint:** Combine defer with error checking.

---

## 19. Function Type (10 points)

Define function type.

**Expected output:**
```
Operation result: 8
```

**Hint:** `type Operation func(int, int) int`

---

## 20. Method on Function (15 points)

Use function as receiver.

**Expected output:**
```
Calculation: 20
```

**Hint:** Define method on function type.

---

## 21. Panic and Recover (20 points)

Handle panic with recover.

**Expected output:**
```
Recovered from panic: something went wrong
```

**Hint:** defer + recover() to handle panic.

---

## 22. Generator Pattern (15 points)

Create number generator.

**Expected output:**
```
1, 2, 3, 4, 5
```

**Hint:** Return function that generates numbers.

---

## 23. Pure Function (10 points)

Write pure function.

**Expected output:**
```
Double of 5: 10
```

**Hint:** Same input always gives same output.