# Category 15: Errors

Learn error handling in Go.

**Total Points: 140**

---

## 1. Return Error (5 points)

Return error from function.

**Expected output:**
```
Error: something went wrong
```

**Hint:** return nil, errors.New("message")

---

## 2. Check Error (5 points)

Check if error exists.

**Expected output:**
```
Failed: something went wrong
```

**Hint:** if err != nil { }

---

## 3. Custom Error Type (10 points)

Create custom error type.

**Expected output:**
```
ValidationError: name is required
```

**Hint:** Define struct implementing Error().

---

## 4. Error with Details (10 points)

Include extra error details.

**Expected output:**
```
InvalidInput: field=name, reason=required
```

**Hint:** Custom struct with fields.

---

## 5. Error Wrapping (15 points)

Wrap errors with context.

**Expected output:**
```
process failed: original error
```

**Hint:** fmt.Errorf("context: %w", err)

---

## 6. Error Chain (15 points)

Unwrap error chain.

**Expected output:**
```
All errors: process -> read -> file
```

**Hint:** errors.Is and errors.As.

---

## 7. Sentinel Errors (10 points)

Define predefined errors.

**Expected output:**
```
ErrNotFound returned
```

**Hint:** var ErrNotFound = errors.New("not found")

---

## 8. Type Assertion Error (10 points)

Check specific error type.

**Expected output:**
```
Validation error: name required
```

**Hint:** errors.As for type assertion.

---

## 9. Handle Multiple Errors (10 points)

Handle different error types.

**Expected output:**
```
Validation error
```

**Hint:** switch err.(type) { }

---

## 10. Error Logging (10 points)

Log errors properly.

**Expected output:**
```
2024/01/01 ERROR: operation failed: error details
```

**Hint:** log.Printf or structured logger.

---

## 11. Recover from Panic (15 points)

Use recover to handle panic.

**Expected output:**
```
Recovered: panic occurred
```

**Hint:** defer func() { recover() }()

---

## 12. Custom Error Formatting (10 points)

Implement Error() string method.

**Expected output:**
```
[ERROR] code=500, msg=internal error
```

**Hint:** Custom Error() implementation.

---

## 13. Error in Go Style (10 points)

Return error as last value.

**Expected output:**
```
Result: , Error: division by zero
```

**Hint:** Convention: return (value, error).

---

## 14. Error Checking Patterns (15 points)

Use common error patterns.

**Expected output:**
```
Not found
```

**Hint:** os.IsNotExist, errors.Is, etc.