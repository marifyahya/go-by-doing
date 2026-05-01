# Category 26: Email (SMTP)

Send emails in Go.

**Total Points: 120**

---

## 1. Send Plain Email (10 points)

Send simple email.

**Expected output:**
```
Email sent
```

**Hint:** net/smtp.SendMail.

---

## 2. HTML Email (10 points)

Send HTML email.

**Expected output:**
```
HTML email sent
```

**Hint:** Set Content-Type: text/html.

---

## 3. Attachments (15 points)

Add attachments.

**Expected output:**
```
Email with attachment sent
```

**Hint:** multipart email.

---

## 4. Multiple Recipients (10 points)

Send to multiple.

**Expected output:**
```
Sent to all
```

**Hint:** Multiple To addresses.

---

## 5. Email Templates (15 points)

Use templates.

**Expected output:**
```
Template rendered
```

**Hint:** text/template.

---

## 6. Config Email (10 points)

SMTP configuration.

**Expected output:**
```
Configured
```

**Hint:** env vars for config.

---

## 7. Queue Integration (15 points)

Queue emails.

**Expected output:**
```
Email queued
```

**Hint:** Send to queue, worker sends.

---

## 8. Welcome Email (15 points)

Send welcome.

**Expected output:**
```
Welcome email sent
```

**Hint:** Trigger on register.

---

## 9. Reset Password Email (15 points)

Send reset email.

**Expected output:**
```
Reset email sent
```

**Hint:** With reset link.

---

## 10. Batch Email (10 points)

Send bulk.

**Expected output:**
```
All sent
```

**Hint:** Loop + send.