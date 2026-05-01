# Category 4: Strings

Work with string manipulation in Go.

**Total Points: 90**

---

## 1. String Length (5 points)

Get the length of a string.

**Expected output:**
```
Length: 5
```

**Hint:** Use `len()` function.

---

## 2. String Concatenation (5 points)

Combine strings.

**Expected output:**
```
Hello World
```

**Hint:** Use `+` operator.

---

## 3. Sprintf (5 points)

Use fmt.Sprintf for string creation.

**Expected output:**
```
Name: Alice, Age: 25
```

**Hint:** `fmt.Sprintf` returns a string.

---

## 4. Index Access (5 points)

Access character by index.

**Expected output:**
```
First char: H
Last char: o
```

**Hint:** String is immutable, access via index.

---

## 5. Substring (10 points)

Extract part of a string.

**Expected output:**
```
First 5: Hello
From 6: World
Middle: lo Wo
```

**Hint:** Use slicing: `s[start:end]`.

---

## 6. Contains (10 points)

Check if string contains substring.

**Expected output:**
```
Contains "ello": true
Contains "xyz": false
```

**Hint:** Use `strings.Contains()`.

---

## 7. To Upper/Lower (10 points)

Convert case.

**Expected output:**
```
Upper: HELLO
Lower: hello
Title: Hello World
```

**Hint:** Use `strings.ToUpper`, `ToLower`, `ToTitle`.

---

## 8. Trim Spaces (10 points)

Remove leading/trailing spaces.

**Expected output:**
```
Trimmed: Hello World
TrimLeft: Hello   
TrimRight:   Hello
```

**Hint:** Use `strings.Trim`, `TrimSpace`.

---

## 9. Split (10 points)

Split string into slice.

**Expected output:**
```
Parts: [one two three]
Count: 3
```

**Hint:** Use `strings.Split()`.

---

## 10. Join (10 points)

Join slice into string.

**Expected output:**
```
Joined: one-two-three
```

**Hint:** Use `strings.Join()`.

---

## 11. Replace (10 points)

Replace substrings.

**Expected output:**
```
Original: Hello World
Replaced: Hello Go
```

**Hint:** Use `strings.Replace()`.

---

## 12. Has Prefix/Suffix (10 points)

Check start and end of string.

**Expected output:**
```
HasPrefix: true
HasSuffix: true
```

**Hint:** Use `strings.HasPrefix`, `HasSuffix`.