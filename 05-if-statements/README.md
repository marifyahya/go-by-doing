# Category 5: If Statements

Learn conditional logic in Go.

**Total Points: 300**

---

## 1. Basic If (5 points)

Use simple if statement.

**Expected output:**
```
Passed
```

**Hint:** `if condition { ... }`

---

## 2. If Else (5 points)

Use if-else.

**Expected output:**
```
Failed
```

**Hint:** `if condition { } else { }`

---

## 3. If Else If (10 points)

Chain multiple conditions.

**Expected output:**
```
Good
```

**Hint:** Use `else if` for multiple branches.

---

## 4. Comparison Operators (10 points)

Use various comparison operators.

**Expected output:**
```
5 > 3: true
5 < 3: false
5 == 5: true
5 != 3: true
```

**Hint:** ==, !=, <, >, <=, >=

---

## 5. Logical Operators (10 points)

Combine conditions with && and ||.

**Expected output:**
```
true && true: true
true && false: false
true || false: true
!true: false
```

**Hint:** && (AND), || (OR), ! (NOT)

---

## 6. Nested If (10 points)

Use nested if statements.

**Expected output:**
```
Pass with distinction
```

**Hint:** You can nest if statements.

---

## 7. Even or Odd (10 points)

Check if number is even or odd.

**Expected output:**
```
7 is odd
```

**Hint:** Use modulo `%` operator.

---

## 8. Positive Negative Zero (10 points)

Classify a number.

**Expected output:**
```
Positive
```

**Hint:** Check if > 0, < 0, or == 0.

---

## 9. Maximum (10 points)

Find the maximum of two numbers.

**Expected output:**
```
Max: 15
```

**Hint:** Use if-else or built-in.

---

## 10. Leap Year (15 points)

Check if a year is a leap year.

**Expected output:**
```
2024 is a leap year: true
2023 is a leap year: false
```

**Hint:** Divisible by 4, but not by 100 unless by 400.

---

## 11. Grade Calculator (15 points)

Convert score to letter grade.

**Expected output:**
```
Score: 85, Grade: B
```

**Hint:** A: 90-100, B: 80-89, C: 70-79, D: 60-69, F: <60

---

## 12. Age Category (15 points)

Categorize by age.

**Expected output:**
```
Age: 25, Category: Adult
```

**Hint:** Child: <13, Teen: 13-19, Adult: 20-64, Senior: 65+

---

## 13. BMI Calculator (20 points)

Calculate and categorize BMI.

**Expected output:**
```
Weight: 70, Height: 1.75, BMI: 22.86, Category: Normal
```

**Hint:** BMI = weight/(height*height). <18.5: Underweight, 18.5-25: Normal, 25-30: Overweight, >30: Obese

---

## 14. Triangle Type (20 points)

Determine triangle type.

**Expected output:**
```
3, 3, 3: Equilateral
3, 4, 5: Scalene
3, 3, 5: Invalid
```

**Hint:** Equilateral: all equal, Isosceles: 2 equal, Scalene: all different. Must satisfy a+b>c.

---

## 15. Day of Week (15 points)

Convert number to day name.

**Expected output:**
```
3: Wednesday
```

**Hint:** Use switch or if-else chain.

---

## 16. Switch Basic (10 points)

Use switch statement.

**Expected output:**
```
Tuesday
```

**Hint:** switch with multiple cases.

---

## 17. Switch Expression (10 points)

Switch on a value.

**Expected output:**
```
Weekend
```

**Hint:** Use switch value { ... }

---

## 18. Switch Multiple Values (10 points)

Handle multiple values in one case.

**Expected output:**
```
Weekday
```

**Hint:** case "Monday", "Tuesday", ...: 

---

## 19. Switch Default (10 points)

Use default case.

**Expected output:**
```
Invalid day
```

**Hint:** default catches unmatched cases.

---

## 20. Switch Fallthrough (15 points)

Understand fallthrough behavior.

**Expected output:**
```
2
3
```

**Hint:** fallthrough transfers to next case (use carefully).

---

## 21. Min Three Numbers (15 points)

Find minimum of three numbers.

**Expected output:**
```
Min: 3
```

**Hint:** Use nested if or min() twice.

---

## 22. FizzBuzz (20 points)

Classic FizzBuzz problem.

**Expected output:**
```
1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, FizzBuzz
```

**Hint:** Print Fizz for 3, Buzz for 5, FizzBuzz for both.

---

## 23. Number Sign (10 points)

Check if number is positive, negative, or zero.

**Expected output:**
```
-5 is negative
```

**Hint:** Use if-else chain.

---

## 24. Eligible to Vote (10 points)

Check voting eligibility.

**Expected output:**
```
Age 17: Not eligible
Age 18: Eligible
```

**Hint:** Must be 18 or older.

---

## 25. Password Strength (20 points)

Check password strength.

**Expected output:**
```
"abc": Weak
"abcd1234": Medium
"Abcd1234!": Strong
```

**Hint:** Weak: <8 chars, Medium: >=8 with number, Strong: >=8 + number + uppercase + special

---

## 26. Time Classification (15 points)

Classify time of day.

**Expected output:**
```
14: Afternoon
```

**Hint:** Morning: 5-11, Afternoon: 12-17, Evening: 18-21, Night: 22-4

---

## 27. Currency Exchange (15 points)

Convert between currencies.

**Expected output:**
```
100 USD to EUR: 92
```

**Hint:** USD to EUR: 0.92, USD to JPY: 149, etc.

---

## 28. Calculator If Else (20 points)

Simple calculator using if-else.

**Expected output:**
```
10 + 5 = 15
10 - 5 = 5
10 * 5 = 50
10 / 5 = 2
```

**Hint:** Use if-else for operation selection.

---

## 29. Divisibility Check (10 points)

Check if number is divisible by both 3 and 5.

**Expected output:**
```
15: Divisible by both
10: Divisible by 5 only
```

**Hint:** Use % operator.

---

## 30. Temperature Convert (15 points)

Convert temperature between Celsius and Fahrenheit.

**Expected output:**
```
0C = 32F
100C = 212F
32F = 0C
```

**Hint:** C = (F-32)*5/9, F = C*9/5 + 32