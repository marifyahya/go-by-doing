# Category 18: WaitGroup & Mutex

Learn synchronization primitives.

**Total Points: 150**

---

## 1. WaitGroup Basics (10 points)

Wait for goroutines.

**Expected output:**
```
All done!
```

**Hint:** sync.WaitGroup, Add, Done, Wait.

---

## 2. Multiple Workers (15 points)

Wait for multiple tasks.

**Expected output:**
```
Worker 1 done
Worker 2 done
All done
```

**Hint:** Add number of goroutines.

---

## 3. Mutex Basics (10 points)

Protect shared data.

**Expected output:**
```
Counter: 1000
```

**Hint:** sync.Mutex, Lock, Unlock.

---

## 4. Safe Counter (15 points)

Concurrent counter.

**Expected output:**
```
Final: 1000
```

**Hint:** Use mutex to protect.

---

## 5. RWMutex (15 points)

Read-write mutex.

**Expected output:**
```
Read operations many
Write operations few
```

**Hint:** sync.RWMutex.

---

## 6. Once (10 points)

Run code once.

**Expected output:**
```
Initialized once
```

**Hint:** sync.Once.

---

## 7. Map with Mutex (15 points)

Thread-safe map.

**Expected output:**
```
Safe map operations
```

**Hint:** Mutex + map.

---

## 8. Pool Pattern (15 points)

Object pool.

**Expected output:**
```
Pooled resource used
```

**Hint:** sync.Pool.

---

## 9. Cond (15 points)

Condition variable.

**Expected output:**
```
Signal received
```

**Hint:** sync.Cond.

---

## 10. Atomic Operations (15 points)

Use atomic package.

**Expected output:**
```
Value: 1000
```

**Hint:** atomic.AddInt64, etc.

---

## 11. Channel vs Mutex (15 points)

Compare approaches.

**Expected output:**
```
Both work correctly
```

**Hint:** Show both patterns.

---

## 12. Proper Cleanup (15 points)

Clean up resources.

**Expected output:**
```
Resources cleaned
```

**Hint:** defer unlock, etc.