# Category 27: Cache (Redis)

Redis caching strategies.

**Total Points: 120**

---

## 1. Set Cache (10 points)

Store in cache.

**Expected output:**
```
Cached
```

**Hint:** redis.Set

---

## 2. Get Cache (10 points)

Retrieve from cache.

**Expected output:**
```
Value: user123
```

**Hint:** redis.Get

---

## 3. Delete Cache (10 points)

Remove from cache.

**Expected output:**
```
Deleted
```

**Hint:** redis.Del

---

## 4. TTL (10 points)

Set expiration.

**Expected output:**
```
Expires in 1 hour
```

**Hint:** Set with EX option.

---

## 5. Cache-Aside (15 points)

Lazy loading.

**Expected output:**
```
Cache miss, DB fetched
```

**Hint:** Check, fetch, store.

---

## 6. Write-Through (15 points)

Write to cache and DB.

**Expected output:**
```
Stored in both
```

**Hint:** Write both.

---

## 7. Cache Invalidation (15 points)

Invalidate cache.

**Expected output:**
```
Invalidated
```

**Hint:** Delete on update.

---

## 8. Cache Keys (10 points)

Key naming.

**Expected output:**
```
user:1:name
```

**Hint:** Prefix patterns.

---

## 9. Serialize JSON (10 points)

Store JSON.

**Expected output:**
```
JSON stored
```

**Hint:** Marshal to JSON.

---

## 10. Distributed Lock (15 points)

Lock with Redis.

**Expected output:**
```
Lock acquired
```

**Hint:** SETNX.