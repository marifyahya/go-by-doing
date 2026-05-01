# Category 1: Basics

Start here! These exercises will help you get familiar with Go syntax and running your first programs.

**Total Points: 60**

---

## 1. Hello World (5 points)

Write a program that prints "Hello, World!" to the screen.

**Expected output:**
```
Hello, World!
```

**Hint:** Use `fmt.Println()` from the `fmt` package.

---

## 2. Hello Name (5 points)

Write a program that prints a greeting with your name.

**Expected output:**
```
Hello, Marif!
```

**Hint:** Use `fmt.Print()` or `fmt.Println()` with a string.

---

## 3. Print Numbers (5 points)

Write a program that prints the numbers 1, 2, 3 on separate lines.

**Expected output:**
```
1
2
3
```

**Hint:** Use multiple `fmt.Println()` statements.

---

## 4. ASCII Art (5 points)

Write a program that prints a simple ASCII art.

**Expected output:**
```
  *
 ***
*****
  *
```

**Hint:** Use multiple Println with spaces.

---

## 5. Your Initials (5 points)

Write a program that prints your initials in large letters.

**Expected output (example):**
```
MMM
 MMM
 MMM
```

**Hint:** Use Println with spaces.

---

## 6. Comments Practice (5 points)

Write a program with comments explaining what it does. The program should print "Go is awesome!".

**Expected output:**
```
Go is awesome!
```

**Hint:** Use `//` for single-line comments.

---

## 7. Multi Line Comment (5 points)

Write a program using a multi-line comment block that prints your favorite quote.

**Expected output:**
```
"The best way to learn is to do."
```

**Hint:** Use `/* */` for multi-line comments.

---

## 8. Printf Basics (5 points)

Use `fmt.Printf()` to print formatted output.

**Expected output:**
```
Name: Go
Version: 1.21
```

**Hint:** Use `%s` for strings, `%d` for integers.

---

## 9. Printf Numbers (5 points)

Use `fmt.Printf()` with different format specifiers.

**Expected output:**
```
Integer: 42
Float: 3.14
Char: A
```

**Hint:** Use `%d` for int, `%f` for float, `%c` for char.

---

## 10. Escape Sequences (10 points)

Use escape sequences to format your output.

**Expected output:**
```
Tab:		test
New line:
Line 1
Line 2
Quote: "Hello"
Backslash: \
```

**Hint:** Use `\t`, `\n`, `\"`, `\\`.

---

## 11. Println vs Print vs Printf (10 points)

Compare the three print functions.

**Expected output:**
```
Print: HelloPrintln: Hello
Printf: Hello
```

**Hint:** Note the difference:
- Print: no newline
- Println: adds newline
- Printf: formatted, no newline

---

## 12. Multiple Packages (10 points)

Create a simple program that uses multiple functions from the fmt package.

**Expected output:**
```
This is a test.
Number: 123
Done!
```

**Hint:** Explore `fmt.Print`, `fmt.Println`, `fmt.Printf`, `fmt.Sprintf`