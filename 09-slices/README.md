# Category 9: Slices

Learn dynamic arrays in Go.

**Total Points: 180**

---

## 1. Declare Slice (5 points)

Create an empty slice.

**Expected output:**
```
[]
```

**Hint:** `var slice []int` or `make([]int, 0)`

---

## 2. Initialize Slice (5 points)

Initialize with values.

**Expected output:**
```
[1 2 3 4 5]
```

**Hint:** `[]int{1, 2, 3, 4, 5}`

---

## 3. Append (10 points)

Add element to slice.

**Expected output:**
```
[1 2 3]
[1 2 3 4]
[1 2 3 4 5]
```

**Hint:** Use append() function.

---

## 4. Append Multiple (10 points)

Append multiple elements.

**Expected output:**
```
[1 2 3 4 5 6]
```

**Hint:** Use spread operator: append(slice, vals...)

---

## 5. Make Slice (10 points)

Create slice with make.

**Expected output:**
```
Length: 3, Capacity: 3
[0 0 0]
```

**Hint:** make([]int, length, capacity)

---

## 6. Copy Slice (10 points)

Copy slice elements.

**Expected output:**
```
Source: [1 2 3]
Copy: [1 2 3]
```

**Hint:** Use copy() function.

---

## 7. Slice Slicing (10 points)

Get subslice.

**Expected output:**
```
Original: [1 2 3 4 5]
Sub: [2 3 4]
```

**Hint:** slice[start:end]

---

## 8. Reslice (10 points)

Modify slice via reslicing.

**Expected output:**
```
[1 2 3 4 5]
[1 2 3]
[1 2 3 4 5]
```

**Hint:** Reslice shares underlying array.

---

## 9. Len vs Cap (10 points)

Understand length vs capacity.

**Expected output:**
```
Length: 3, Capacity: 5
Length: 5, Capacity: 10
```

**Hint:** len() vs cap()

---

## 10. Filter Slice (15 points)

Filter elements.

**Expected output:**
```
Original: [1 2 3 4 5 6]
Evens: [2 4 6]
```

**Hint:** Create new slice with condition.

---

## 11. Map Slice (15 points)

Transform elements.

**Expected output:**
```
Original: [1 2 3]
Doubled: [2 4 6]
```

**Hint:** Apply function to each element.

---

## 12. Sort Slice (15 points)

Sort slice.

**Expected output:**
```
Unsorted: [3 1 4 1 5]
Sorted: [1 1 3 4 5]
```

**Hint:** sort.Ints(slice)

---

## 13. Reverse Slice (10 points)

Reverse slice order.

**Expected output:**
```
[1 2 3 4 5]
[5 4 3 2 1]
```

**Hint:** Swap elements from ends.

---

## 14. Remove Duplicates (15 points)

Remove duplicate values.

**Expected output:**
```
[1 2 3 4 5]
```

**Hint:** Use map to track unique.

---

## 15. Chunk Slice (20 points)

Split into chunks.

**Expected output:**
```
[[1 2] [3 4] [5]]
```

**Hint:** Loop with step size.

---

## 16. Flatten Slice (20 points)

Flatten nested slices.

**Expected output:**
```
[1 2 3 4 5 6]
```

**Hint:** Append each sub-slice.

---

## 17. Rotate Slice (15 points)

Rotate elements.

**Expected output:**
```
[1 2 3 4 5]
After rotate 2: [3 4 5 1 2]
```

**Hint:** Use two slices and append.

---

## 18. Insert at Position (20 points)

Insert element at index.

**Expected output:**
```
[1 2 3 4 5]
After insert 99 at 2: [1 2 99 3 4 5]
```

**Hint:** Append and shift.

---

## 19. Delete at Position (15 points)

Delete element at index.

**Expected output:**
```
[1 2 3 4 5]
After delete at 2: [1 2 4 5]
```

**Hint:** Append and skip.

---

## 20. Merge Slices (10 points)

Merge two slices.

**Expected output:**
```
[1 2 3] + [4 5 6] = [1 2 3 4 5 6]
```

**Hint:** Use append.