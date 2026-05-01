# Category 8: Arrays

Learn about fixed-size arrays in Go.

**Total Points: 180**

---

## 1. Declare Array (5 points)

Create an array.

**Expected output:**
```
[0 0 0 0 0]
```

**Hint:** `var arr [5]int`

---

## 2. Initialize Array (5 points)

Initialize with values.

**Expected output:**
```
[1 2 3 4 5]
```

**Hint:** `arr := [5]int{1, 2, 3, 4, 5}`

---

## 3. Access Element (5 points)

Access array element by index.

**Expected output:**
```
First: 10
Last: 50
```

**Hint:** arr[index], zero-indexed.

---

## 4. Update Element (5 points)

Update array element.

**Expected output:**
```
[1 2 99 4 5]
```

**Hint:** arr[index] = value

---

## 5. Array Length (5 points)

Get array length.

**Expected output:**
```
Length: 5
```

**Hint:** Use len() function.

---

## 6. Iterate Array (10 points)

Loop through array.

**Expected output:**
```
Index: 0, Value: 1
Index: 1, Value: 2
Index: 2, Value: 3
```

**Hint:** Use for loop with range.

---

## 7. Sum Array (10 points)

Calculate sum of array elements.

**Expected output:**
```
Sum: 15
```

**Hint:** Loop and accumulate.

---

## 8. Find Max (15 points)

Find maximum value.

**Expected output:**
```
Max: 9
```

**Hint:** Track max while iterating.

---

## 9. Find Min (10 points)

Find minimum value.

**Expected output:**
```
Min: 1
```

**Hint:** Similar to finding max.

---

## 10. Array Reverse (15 points)

Reverse an array.

**Expected output:**
```
Original: [1 2 3 4 5]
Reversed: [5 4 3 2 1]
```

**Hint:** Swap elements from ends.

---

## 11. Copy Array (10 points)

Copy array to new array.

**Expected output:**
```
Original: [1 2 3]
Copy: [1 2 3]
```

**Hint:** Create new array and copy.

---

## 12. 2D Array (15 points)

Work with 2D array.

**Expected output:**
```
Matrix:
1 2 3
4 5 6
```

**Hint:** `var matrix [3][3]int`

---

## 13. 2D Array Sum (15 points)

Sum all elements in 2D array.

**Expected output:**
```
Total: 21
```

**Hint:** Nested loops.

---

## 14. Diagonal Sum (15 points)

Sum diagonal elements.

**Expected output:**
```
Diagonal sum: 15
```

**Hint:** Access arr[i][i].

---

## 15. Array as String (10 points)

Convert array to string.

**Expected output:**
```
[1, 2, 3, 4, 5]
```

**Hint:** Use fmt.Sprintf or strings.Join.

---

## 16. Search Element (15 points)

Linear search in array.

**Expected output:**
```
Found at index: 2
Not found: -1
```

**Hint:** Iterate and compare.

---

## 17. Insert Element (20 points)

Insert element at position.

**Expected output:**
```
[1 2 99 3 4 5]
```

**Hint:** Shift elements, then insert.

---

## 18. Delete Element (20 points)

Delete element at position.

**Expected output:**
```
[1 2 4 5]
```

**Hint:** Shift elements down.

---

## 19. Array of Strings (10 points)

Create array of strings.

**Expected output:**
```
[apple banana cherry]
```

**Hint:** `var fruits [3]string`

---

## 20. Sort Array (15 points)

Sort array (bubble sort).

**Expected output:**
```
Unsorted: [3 1 4 1 5]
Sorted: [1 1 3 4 5]
```

**Hint:** Compare and swap adjacent.