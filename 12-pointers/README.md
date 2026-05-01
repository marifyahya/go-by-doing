# Category 12: Pointers

Learn memory addresses and references.

**Total Points: 120**

---

## 1. Declare Pointer (5 points)

Create a pointer.

**Expected output:**
```
0xc000012345
```

**Hint:** var p *int

---

## 2. Address Of (5 points)

Get address of variable.

**Expected output:**
```
Address: 0xc000012345
Value: 10
```

**Hint:** Use & operator.

---

## 3. Dereference Pointer (5 points)

Access value through pointer.

**Expected output:**
```
Value: 10
```

**Hint:** Use * operator.

---

## 4. Pointer to Struct (10 points)

Create pointer to struct.

**Expected output:**
```
Name: Alice
Age: 25
```

**Hint:** p := &Person{Name: "Alice"}

---

## 5. Modify via Pointer (10 points)

Modify value through pointer.

**Expected output:**
```
Before: 10
After: 20
```

**Hint:** *ptr = newValue

---

## 6. Nil Pointer (5 points)

Understand nil pointer.

**Expected output:**
```
nil
```

**Hint:** Uninitialized pointer is nil.

---

## 7. Check Nil (10 points)

Check if pointer is nil.

**Expected output:**
```
Pointer is nil
```

**Hint:** if p == nil { }

---

## 8. New Function (10 points)

Use new() to create pointer.

**Expected output:**
```
0
```

**Hint:** p := new(int)

---

## 9. Pointer to Slice (10 points)

Modify slice via pointer.

**Expected output:**
```
[1 2 99 4 5]
```

**Hint:** Pass pointer to slice.

---

## 10. Pointer to Map (10 points)

Modify map via pointer.

**Expected output:**
```
{a:1 b:2}
```

**Hint:** Pass pointer to map.

---

## 11. Return Pointer (10 points)

Return pointer from function.

**Expected output:**
```
Alice
```

**Hint:** Return *Person from function.

---

## 12. Double Pointer (15 points)

Use pointer to pointer.

**Expected output:**
```
Value: 10
```

**Hint:** **int for double pointer.

---

## 13. Pointer Array (10 points)

Create array of pointers.

**Expected output:**
```
1 2 3
```

**Hint:** []*int

---

## 14. Go Pointer Restrictions (10 points)

Understand pointer arithmetic.

**Expected output:**
```
Error: pointer arithmetic not allowed
```

**Hint:** Go doesn't allow pointer arithmetic.

---

## 15. Stack vs Heap (15 points)

Understand memory allocation.

**Expected output:**
```
Value on heap
```

**Hint:** Compiler decides stack vs heap.