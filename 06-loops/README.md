# Category 6: Loops

Master iteration in Go.

**Total Points: 250**

---

## 1. For Basic (5 points)

Use basic for loop.

**Expected output:**
```
0
1
2
3
4
```

**Hint:** `for i := 0; i < 5; i++ { ... }`

---

## 2. For Range (5 points)

Use for-range loop.

**Expected output:**
```
0 a
1 b
2 c
```

**Hint:** `for index, value := range slice { ... }`

---

## 3. While Style (10 points)

Use for as while loop.

**Expected output:**
```
Countdown: 5
4
3
2
1
0
```

**Hint:** `for condition { ... }`

---

## 4. Infinite Loop (10 points)

Create infinite loop with break.

**Expected output:**
```
1
2
3
```

**Hint:** `for { ... if condition { break } }`

---

## 5. Print 1 to 10 (5 points)

Print numbers 1 to 10.

**Expected output:**
```
1 2 3 4 5 6 7 8 9 10
```

**Hint:** Use for loop.

---

## 6. Print 10 to 1 (5 points)

Print numbers 10 to 1.

**Expected output:**
```
10 9 8 7 6 5 4 3 2 1
```

**Hint:** Use decrement in for loop.

---

## 7. Sum of Numbers (10 points)

Calculate sum of 1 to n.

**Expected output:**
```
Sum of 1 to 10: 55
```

**Hint:** Use loop to accumulate sum.

---

## 8. Even Numbers (10 points)

Print even numbers up to n.

**Expected output:**
```
Even numbers up to 10: 2 4 6 8 10
```

**Hint:** Use step of 2 or check modulo.

---

## 9. Multiplication Table (15 points)

Print multiplication table.

**Expected output:**
```
5 x 1 = 5
5 x 2 = 10
...
5 x 10 = 50
```

**Hint:** Nested loop or single loop with multiplication.

---

## 10. Count Characters (10 points)

Count characters in string.

**Expected output:**
```
"hello" has 5 characters
```

**Hint:** Use len() or for-range.

---

## 11. Find Maximum (15 points)

Find max in slice.

**Expected output:**
```
Max: 9
```

**Hint:** Iterate and track max value.

---

## 12. Find Minimum (10 points)

Find minimum in slice.

**Expected output:**
```
Min: 1
```

**Hint:** Similar to finding maximum.

---

## 13. Break Statement (10 points)

Use break to exit loop.

**Expected output:**
```
0 1 2 3
```

**Hint:** Break exits loop when condition met.

---

## 14. Continue Statement (10 points)

Use continue to skip iteration.

**Expected output:**
```
1 2 4 5
(skip 3)
```

**Hint:** Continue skips to next iteration.

---

## 15. Nested Loops (15 points)

Use nested for loops.

**Expected output:**
```
*
**
***
****
*****
```

**Hint:** Loop inside loop.

---

## 16. Pattern 1 (15 points)

Print pattern.

**Expected output:**
```
1
12
123
1234
12345
```

**Hint:** Nested loops with incrementing inner counter.

---

## 17. Pattern 2 (15 points)

Print reverse pattern.

**Expected output:**
```
12345
1234
123
12
1
```

**Hint:** Nested loops with decrement.

---

## 18. Prime Numbers (20 points)

Check if number is prime.

**Expected output:**
```
7 is prime: true
10 is prime: false
```

**Hint:** Check divisibility from 2 to sqrt(n).

---

## 19. Fibonacci (20 points)

Generate Fibonacci sequence.

**Expected output:**
```
0, 1, 1, 2, 3, 5, 8, 13, 21, 34
```

**Hint:** Each number is sum of previous two.

---

## 20. Factorial (15 points)

Calculate factorial.

**Expected output:**
```
5! = 120
```

**Hint:** Multiply all numbers from 1 to n.

---

## 21. Reverse Number (15 points)

Reverse a number.

**Expected output:**
```
12345 reversed: 54321
```

**Hint:** Use modulo and division.

---

## 22. Palindrome Check (15 points)

Check if number is palindrome.

**Expected output:**
```
12321 is palindrome: true
123 is palindrome: false
```

**Hint:** Compare reversed number with original.

---

## 23. Power of Number (15 points)

Calculate power.

**Expected output:**
```
2^10 = 1024
```

**Hint:** Multiply base by itself exponent times.

---

## 24. Sum of Digits (15 points)

Sum digits of a number.

**Expected output:**
```
123 sum of digits: 6
```

**Hint:** Use modulo to get each digit.

---

## 25. Count Vowels (20 points)

Count vowels in string.

**Expected output:**
```
"hello world" has 3 vowels
```

**Hint:** Iterate through string, check each char.

---

## 26. Binary Search (25 points)

Search in sorted array.

**Expected output:**
```
Found at index: 4
```

**Hint:** Use two pointers, narrow down range.