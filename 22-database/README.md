# Category 22: Database (PostgreSQL)

Connect to database.

**Total Points: 160**

---

## 1. Connect to DB (10 points)

Establish database connection.

**Expected output:**
```
Connected to PostgreSQL
```

**Hint:** sql.Open("postgres", connStr)

---

## 2. Ping Database (5 points)

Verify connection.

**Expected output:**
```
PING OK
```

**Hint:** db.Ping()

---

## 3. Execute Query (10 points)

Run INSERT/UPDATE/DELETE.

**Expected output:**
```
Rows affected: 1
```

**Hint:** db.Exec()

---

## 4. Query Single Row (10 points)

Select one row.

**Expected output:**
```
Name: Alice
```

**Hint:** db.QueryRow()

---

## 5. Query Multiple Rows (10 points)

Select multiple rows.

**Expected output:**
```
[Alice Bob Charlie]
```

**Hint:** db.Query()

---

## 6. Scan Results (15 points)

Scan to struct.

**Expected output:**
```
User: Alice, Age: 25
```

**Hint:** Scan() to struct.

---

## 7. Prepare Statement (10 points)

Use prepared statement.

**Expected output:**
```
Prepared executed
```

**Hint:** db.Prepare()

---

## 8. Transaction (15 points)

Execute transaction.

**Expected output:**
```
Transaction committed
```

**Hint:** db.Begin(), tx.Commit()

---

## 9. Transaction Rollback (15 points)

Handle rollback.

**Expected output:**
```
Rolled back
```

**Hint:** tx.Rollback()

---

## 10. Connect with sqlx (15 points)

Use sqlx extension.

**Expected output:**
```
sqlx connected
```

**Hint:** sqlx.Connect()

---

## 11. Named Queries (10 points)

Use named parameters.

**Expected output:**
```
Query executed
```

**Hint:** Named parameter syntax.

---

## 12. In Clause (15 points)

Handle IN queries.

**Expected output:**
```
[1 2 3]
```

**Hint:** sqlx.In()

---

## 13. Null Values (10 points)

Handle NULL values.

**Expected output:**
```
Optional field handled
```

**Hint:** sql.NullString.

---

## 14. Connection Pool (10 points)

Configure pool.

**Expected output:**
```
Pool configured
```

**Hint:** db.SetMaxOpenConns()

---

## 15. Context with Query (10 points)

Cancel queries with context.

**Expected output:**
```
Query with timeout
```

**Hint:** QueryContext()