# Category 14: Interfaces

Learn polymorphism in Go.

**Total Points: 160**

---

## 1. Define Interface (5 points)

Create an interface.

**Expected output:**
```
Shape has Area method
```

**Hint:** type Shape interface { }

---

## 2. Implement Interface (10 points)

Implement interface with struct.

**Expected output:**
```
Area: 25
```

**Hint:** Define methods matching interface.

---

## 3. Interface as Parameter (10 points)

Accept interface in function.

**Expected output:**
```
Total: 30
```

**Hint:** func PrintArea(s Shape) { }

---

## 4. Multiple Methods (10 points)

Interface with multiple methods.

**Expected output:**
```
Area: 25, Perimeter: 20
```

**Hint:** Define multiple methods.

---

## 5. Empty Interface (10 points)

Use empty interface.

**Expected output:**
```
Value: 42
Value: hello
```

**Hint:** interface{} accepts any type.

---

## 6. Type Assertion (15 points)

Assert interface type.

**Expected output:**
```
Type is int, Value: 42
```

**Hint:** value, ok := i.(type)

---

## 7. Type Switch (15 points)

Switch on interface type.

**Expected output:**
```
Type: int
```

**Hint:** switch v := i.(type) { }

---

## 8. Interface Composition (10 points)

Combine interfaces.

**Expected output:**
```
Has Area and Perimeter
```

**Hint:** Embed interfaces.

---

## 9. Error Interface (10 points)

Implement error interface.

**Expected output:**
```
Error: custom error
```

**Hint:** func (e Error) Error() string

---

## 10. Reader Interface (10 points)

Implement io.Reader.

**Expected output:**
```
Read 5 bytes
```

**Hint:** func (r) Read(p []byte) (n, error)

---

## 11. Writer Interface (10 points)

Implement io.Writer.

**Expected output:**
```
Wrote 5 bytes
```

**Hint:** func (w) Write(p []byte) (n, error)

---

## 12. Interface with Pointer (15 points)

Implement with pointer receiver.

**Expected output:**
```
Area: 25
```

**Hint:** Use *Type for implementation.

---

## 13. Nil Interface (10 points)

Handle nil interface.

**Expected output:**
```
Nil interface
```

**Hint:** var i Interface = nil

---

## 14. Interface Factory (15 points)

Return interface from function.

**Expected output:**
```
Area: 25
```

**Hint:** Return concrete type implementing interface.

---

## 15. Mock Interface (15 points)

Create mock for testing.

**Expected output:**
```
Mock called
```

**Hint:** Implement interface in test.