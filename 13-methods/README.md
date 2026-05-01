# Category 13: Methods

Learn methods on types in Go.

**Total Points: 200**

---

## 1. Value Receiver (5 points)

Create method with value receiver.

**Expected output:**
```
Area: 30
```

**Hint:** func (p Person) Method() { }

---

## 2. Pointer Receiver (10 points)

Create method with pointer receiver.

**Expected output:**
```
Area: 30
After: 60
```

**Hint:** func (p *Person) Method() { }

---

## 3. Method on Struct (5 points)

Add method to struct.

**Expected output:**
```
Alice is 25 years old
```

**Hint:** Define method on struct type.

---

## 4. Method on Int (10 points)

Add method to int.

**Expected output:**
```
5 is even: true
```

**Hint:** Cannot directly extend, use custom type.

---

## 5. Method on String (10 points)

Add method to string.

**Expected output:**
```
hello: 5
```

**Hint:** Create custom string type.

---

## 6. Method Chaining (15 points)

Return same type for chaining.

**Expected output:**
```
Result: 10
```

**Hint:** Return *Type for chaining.

---

## 7. Multiple Methods (10 points)

Define multiple methods.

**Expected output:**
```
Alice
25
```

**Hint:** Multiple methods on same type.

---

## 8. Override Method (15 points)

Understand method overriding.

**Expected output:**
```
Employee: Alice, TechCorp
```

**Hint:** Embedding + method override.

---

## 9. Call Other Methods (10 points)

Call method from same struct.

**Expected output:**
```
Area: 30
Perimeter: 22
```

**Hint:** p.otherMethod()

---

## 10. Method with Parameters (10 points)

Method with input parameters.

**Expected output:**
```
Distance: 5
```

**Hint:** func (p Point) Distance(other Point) float64

---

## 11. Method with Return (10 points)

Method returning value.

**Expected output:**
```
Full Name: Alice Smith
```

**Hint:** func (p Person) GetName() string

---

## 12. Method on Slice (15 points)

Add method to custom slice.

**Expected output:**
```
Sum: 15
```

**Hint:** Define type alias for slice.

---

## 13. Method on Map (15 points)

Add method to custom map.

**Expected output:**
```
Keys: [a b c]
```

**Hint:** Define custom map type.

---

## 14. Generic Method (20 points)

Create generic-style method.

**Expected output:**
```
Converted: 42
```

**Hint:** Use interface{} or type parameters (Go 1.18+)

---

## 15. Validate Method (15 points)

Create validation method.

**Expected output:**
```
Valid: true
Valid: false
```

**Hint:** Return bool + error.

---

## 16. Builder Pattern (15 points)

Use builder pattern with methods.

**Expected output:**
```
User: Alice, Age: 25, Admin: true
```

**Hint:** Return *Type for chaining.

---

## 17. Decorator Pattern (15 points)

Wrap method behavior.

**Expected output:**
```
[LOG] Area called
Area: 30
[LOG] Done
```

**Hint:** Add logging in method.

---

## 18. Interface with Methods (10 points)

Implement interface via methods.

**Expected output:**
```
Area: 30
```

**Hint:** Define interface, implement with methods.