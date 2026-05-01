# Category 24: Authentication

JWT and password handling.

**Total Points: 160**

---

## 1. Password Hash (10 points)

Hash password with bcrypt.

**Expected output:**
```
Hash generated: $2a$...
```

**Hint:** bcrypt.GenerateFromPassword.

---

## 2. Password Verify (10 points)

Verify password.

**Expected output:**
```
Valid: true
```

**Hint:** bcrypt.CompareHashAndPassword.

---

## 3. JWT Generate (15 points)

Create JWT token.

**Expected output:**
```
Token: eyJhbG...
```

**Hint:** jwt.New, jwt.SignedString.

---

## 4. JWT Verify (10 points)

Verify JWT token.

**Expected output:**
```
Valid: true
```

**Hint:** jwt.Parse, Validate.

---

## 5. JWT Claims (10 points)

Custom claims.

**Expected output:**
```
UserID: 123
```

**Hint:** Custom struct for claims.

---

## 6. JWT Middleware (15 points)

Protect routes.

**Expected output:**
```
401 Unauthorized
```

**Hint:** Check Authorization header.

---

## 7. Register User (15 points)

User registration.

**Expected output:**
```
User registered: alice
```

**Hint:** Hash password, store.

---

## 8. Login (15 points)

User login.

**Expected output:**
```
Token generated
```

**Hint:** Verify credentials, return JWT.

---

## 9. Refresh Token (15 points)

Refresh JWT.

**Expected output:**
```
New token
```

**Hint:** Generate new token.

---

## 10. Logout (10 points)

Invalidate token.

**Expected output:**
```
Token blacklisted
```

**Hint:** Add to blacklist.

---

## 11. Role Based Access (15 points)

Check user role.

**Expected output:**
```
Admin access granted
```

**Hint:** Include role in claims.

---

## 12. Password Reset (15 points)

Reset password flow.

**Expected output:**
```
Reset email sent
```

**Hint:** Generate reset token.

---

## 13. Rate Limit Login (15 points)

Limit login attempts.

**Expected output:**
```
Too many attempts
```

**Hint:** Track attempts, block.