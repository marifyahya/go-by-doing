# Category 21: HTTP Middleware

Learn middleware pattern.

**Total Points: 120**

---

## 1. Basic Middleware (10 points)

Wrap handler with middleware.

**Expected output:**
```
Middleware -> Handler
```

**Hint:** Return http.HandlerFunc.

---

## 2. Logging Middleware (10 points)

Log requests.

**Expected output:**
```
GET /path - 200 - 10ms
```

**Hint:** Log before/after.

---

## 3. Auth Middleware (15 points)

Check authorization.

**Expected output:**
```
401 Unauthorized
```

**Hint:** Check header.

---

## 4. CORS Middleware (10 points)

Handle CORS.

**Expected output:**
```
Access-Control-Allow-Origin: *
```

**Hint:** Set CORS headers.

---

## 5. Recovery Middleware (15 points)

Recover from panic.

**Expected output:**
```
Recovered from panic
```

**Hint:** defer + recover().

---

## 6. Rate Limiter (15 points)

Limit requests.

**Expected output:**
```
429 Too Many Requests
```

**Hint:** Count requests.

---

## 7. Header Middleware (10 points)

Set response headers.

**Expected output:**
```
X-Custom-Header: value
```

**Hint:** w.Header().Set()

---

## 8. Middleware Chain (15 points)

Chain multiple middleware.

**Expected output:**
```
Log -> Auth -> Handler
```

**Hint:** Compose middleware.

---

## 9. Context Middleware (10 points)

Pass data via context.

**Expected output:**
```
User ID: 123
```

**Hint:** r.Context().

---

## 10. Middleware with Config (10 points)

Configurable middleware.

**Expected output:**
```
Custom config used
```

**Hint:** Pass config.