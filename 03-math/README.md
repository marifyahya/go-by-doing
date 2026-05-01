# Category 3: Math

Practice mathematical operations in Go.

**Total Points: 60**

---

## 1. Basic Addition (5 points)

Add two numbers.

**Expected output:**
```
5 + 3 = 8
```

**Hint:** Use `+` operator.

---

## 2. Basic Arithmetic (5 points)

Perform multiple arithmetic operations.

**Expected output:**
```
10 + 5 = 15
10 - 5 = 5
10 * 5 = 50
10 / 5 = 2
```

**Hint:** Use +, -, *, / operators.

---

## 3. Modulo (5 points)

Use modulo operator.

**Expected output:**
```
17 % 5 = 2
```

**Hint:** Use `%` for remainder.

---

## 4. Order of Operations (5 points)

Follow correct operator precedence.

**Expected output:**
```
2 + 3 * 4 = 14
(2 + 3) * 4 = 20
```

**Hint:** Use parentheses to control order.

---

## 5. Increment Decrement (5 points)

Use ++ and -- operators.

**Expected output:**
```
Initial: 5
After ++: 6
After --: 5
```

**Hint:** ++ and -- are statements, not expressions in Go.

---

## 6. Compound Assignment (5 points)

Use +=, -=, *=, /=.

**Expected output:**
```
x = 10
x += 5: 15
x -= 3: 12
x *= 2: 24
```

**Hint:** Compound operators combine operation and assignment.

---

## 7. Math Package (10 points)

Use functions from math package.

**Expected output:**
```
Absolute: 5
Square root: 3
Power: 8
Ceil: 4
Floor: 3
```

**Hint:** Import "math" and use functions.

---

## 8. Integer vs Float Division (10 points)

Understand the difference.

**Expected output:**
```
Integer: 5 / 2 = 2
Float: 5.0 / 2.0 = 2.5
```

**Hint:** Division behavior differs for int vs float.

---

## 9. Min Max Functions (10 points)

Use math.Min and math.Max.

**Expected output:**
```
Min of 5, 10: 5
Max of 5, 10: 10
```

**Hint:** math.Min and math.Max take float64.

---

## 10. Large Numbers (10 points)

Handle large number operations.

**Expected output:**
```
1000000 * 1000000 = 1000000000000
```

**Hint:** Use int64 for large numbers.