# Category 19: HTTP Basics

Build web servers in Go.

**Total Points: 200**

---

## 1. Hello Server (10 points)

Create basic HTTP server.

**Expected output:**
```
Server running on :8080
(curl localhost:8080 returns "Hello")
```

**Hint:** http.ListenAndServe

---

## 2. HandleFunc (5 points)

Register handler function.

**Expected output:**
```
Handler registered
```

**Hint:** http.HandleFunc

---

## 3. Request Method (10 points)

Check request method.

**Expected output:**
```
Method: GET
```

**Hint:** r.Method

---

## 4. Query Parameters (10 points)

Parse URL query.

**Expected output:**
```
Name: John
Age: 25
```

**Hint:** r.URL.Query().Get()

---

## 5. Path Parameters (15 points)

Extract path segments.

**Expected output:**
```
User ID: 123
```

**Hint:** r.URL.Path

---

## 6. Form Data (10 points)

Read form data.

**Expected output:**
```
Username: john
```

**Hint:** r.Form.Get()

---

## 7. Headers (10 points)

Read request headers.

**Expected output:**
```
Content-Type: application/json
```

**Hint:** r.Header.Get()

---

## 8. Response Writer (10 points)

Write response.

**Expected output:**
```
Response sent
```

**Hint:** w.Write()

---

## 9. Status Code (10 points)

Set status code.

**Expected output:**
```
Status: 200 OK
```

**Hint:** w.WriteHeader()

---

## 10. JSON Response (15 points)

Return JSON.

**Expected output:**
```
{"message":"success"}
```

**Hint:** Set Content-Type header.

---

## 11. POST Handler (10 points)

Handle POST request.

**Expected output:**
```
Data processed
```

**Hint:** Check r.Method == "POST"

---

## 12. Custom 404 (10 points)

Return 404.

**Expected output:**
```
404 Not Found
```

**Hint:** http.NotFound()

---

## 13. Redirect (10 points)

Handle redirect.

**Expected output:**
```
Redirected to /new
```

**Hint:** http.Redirect()

---

## 14. ServerMux (15 points)

Use router.

**Expected output:**
```
Routes registered
```

**Hint:** http.NewServeMux()

---

## 15. Multiple Handlers (10 points)

Handle multiple paths.

**Expected output:**
```
/ -> home
/about -> about
```

**Hint:** Multiple http.HandleFunc

---

## 16. Request Body (15 points)

Read request body.

**Expected output:**
```
Body length: 100
```

**Hint:** r.Body, io.ReadAll

---

## 17. URL Parsing (10 points)

Parse URL components.

**Expected output:**
```
Path: /users
Query: name=john
```

**Hint:** r.URL fields.

---

## 18. Static Files (15 points)

Serve static files.

**Expected output:**
```
CSS/JS served
```

**Hint:** http.FileServer

---

## 19. HTTPS Server (10 points)

Create HTTPS server.

**Expected output:**
```
HTTPS server running
```

**Hint:** ListenAndServeTLS

---

## 20. Graceful Shutdown (15 points)

Handle server shutdown.

**Expected output:**
```
Server stopped gracefully
```

**Hint:** http.Server with context.