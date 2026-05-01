# Category 2: Variables

Learn how to declare and use variables in Go.

**Total Points: 80**

---

## 1. Declare Variable (5 points)

Declare a variable and print it.

**Expected output:**
```
42
```

**Hint:** Use `var name type` syntax.

---

## 2. Declare and Initialize (5 points)

Declare a variable with initial value.

**Expected output:**
```
Hello
```

**Hint:** `var name type = value`

---

## 3. Short Declaration (5 points)

Use short variable declaration.

**Expected output:**
```
Go
```

**Hint:** Use `:=` for short declaration.

---

## 4. Multiple Variables (5 points)

Declare multiple variables at once.

**Expected output:**
```
Alice 25
Bob 30
```

**Hint:** Use `var ( ... )` block.

---

## 5. Type Inference (5 points)

Let Go infer the type.

**Expected output:**
```
100
3.14
true
```

**Hint:** `:=` infers type from value.

---

## 6. Zero Values (10 points)

Print zero values of different types.

**Expected output:**
```
int: 0
float: 0
bool: false
string:
```

**Hint:** Uninitialized variables have zero values.

---

## 7. Constant (5 points)

Declare and use a constant.

**Expected output:**
```
Pi is approximately 3.14159
```

**Hint:** Use `const` keyword.

---

## 8. Multiple Constants (5 points)

Declare multiple constants.

**Expected output:**
```
Red: #FF0000
Green: #00FF00
Blue: #0000FF
```

**Hint:** Use `const ( ... )` block.

---

## 9. Iota (10 points)

Use iota for sequential constants.

**Expected output:**
```
0
1
2
```

**Hint:** `iota` generates sequential numbers.

---

## 10. Type Conversion (10 points)

Convert between types.

**Expected output:**
```
int to float: 42.00
float to int: 3
int to string: 42
```

**Hint:** Use `type(value)` for conversion.

---

## 11. Shadowing (10 points)

Understand variable shadowing.

**Expected output:**
```
Outer: 1
Inner: 2
Outer after: 1
```

**Hint:** Inner variable shadows outer.

---

## 12. Blank Identifier (10 points)

Use blank identifier for unused values.

**Expected output:**
```
Value: 10
Ignored the rest
```

**Hint:** Use `_` for ignored values.