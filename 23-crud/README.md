# Category 23: CRUD Operations

Create, Read, Update, Delete.

**Total Points: 200**

---

## 1. Create User (15 points)

Insert new record.

**Expected output:**
```
User created with ID: 1
```

**Hint:** INSERT with RETURNING.

---

## 2. Read User by ID (10 points)

Select single record.

**Expected output:**
```
User: Alice, Email: alice@test.com
```

**Hint:** WHERE id = ?

---

## 3. Read All Users (10 points)

Select all records.

**Expected output:**
```
[Alice, Bob, Charlie]
```

**Hint:** SELECT without WHERE.

---

## 4. Update User (15 points)

Update record.

**Expected output:**
```
User updated
```

**Hint:** UPDATE with WHERE.

---

## 5. Delete User (10 points)

Delete record.

**Expected output:**
```
User deleted
```

**Hint:** DELETE with WHERE.

---

## 6. Pagination (15 points)

Paginate results.

**Expected output:**
```
Page 1 of 3
[user1, user2, user3]
```

**Hint:** LIMIT and OFFSET.

---

## 7. Filtering (15 points)

Filter with WHERE.

**Expected output:**
```
Active users: 5
```

**Hint:** WHERE conditions.

---

## 8. Sorting (10 points)

Sort results.

**Expected output:**
```
[Zack, Alice, Bob]
```

**Hint:** ORDER BY.

---

## 9. Search (15 points)

Search with LIKE.

**Expected output:**
```
Found: 2
```

**Hint:** LIKE %pattern%.

---

## 10. Count Records (10 points)

Count with aggregate.

**Expected output:**
```
Total: 100
```

**Hint:** COUNT(*).

---

## 11. Batch Create (15 points)

Insert multiple records.

**Expected output:**
```
10 users created
```

**Hint:** Loop or batch insert.

---

## 12. Batch Update (15 points)

Update multiple records.

**Expected output:**
```
5 updated
```

**Hint:** Update with IN clause.

---

## 13. Upsert (15 points)

Insert or update.

**Expected output:**
```
Inserted/Updated
```

**Hint:** ON CONFLICT DO UPDATE.

---

## 14. Soft Delete (10 points)

Mark as deleted.

**Expected output:**
```
Marked as deleted
```

**Hint:** UPDATE deleted_at.

---

## 15. Validation (15 points)

Validate before CRUD.

**Expected output:**
```
Validation failed: email required
```

**Hint:** Check before execute.

---

## 16. Error Handling (10 points)

Handle CRUD errors.

**Expected output:**
```
Error: duplicate key
```

**Hint:** Check specific errors.