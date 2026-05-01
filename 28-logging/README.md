# Category 28: Logging

Structured logging.

**Total Points: 120**

---

## 1. Basic Log (10 points)

Use log package.

**Expected output:**
```
2024/01/01 12:00:00 INFO Starting application
```

**Hint:** log.Println.

---

## 2. Custom Logger (10 points)

Create custom logger.

**Expected output:**
```
APP: Message
```

**Hint:** log.New.

---

## 3. Log Levels (15 points)

Implement levels.

**Expected output:**
```
[DEBUG] debug message
[INFO] info message
[ERROR] error message
```

**Hint:** Check level before log.

---

## 4. JSON Log (15 points)

Structured JSON.

**Expected output:**
```
{"level":"info","message":"request","ip":"127.0.0.1"}
```

**Hint:** Use zerolog or zap.

---

## 5. Request Logger (15 points)

Log HTTP requests.

**Expected output:**
```
GET /api/users 200 10ms
```

**Hint:** Middleware logging.

---

## 6. Error Logger (10 points)

Log errors.

**Expected output:**
```
ERROR: failed to connect: connection refused
```

**Hint:** Log with context.

---

## 7. File Logging (10 points)

Log to file.

**Expected output:**
```
Logged to app.log
```

**Hint:** os.OpenFile.

---

## 8. Log Rotation (15 points)

Rotate logs.

**Expected output:**
```
Rotated to new file
```

**Hint:** Daily/size-based rotation.

---

## 9. Context Logger (10 points)

Log with context.

**Expected output:**
```
UserID: 123 action: login
```

**Hint:** Add fields.

---

## 10. Audit Log (10 points)

Track important actions.

**Expected output:**
```
AUDIT: user=123 action=delete target=456
```

**Hint:** Separate audit log.