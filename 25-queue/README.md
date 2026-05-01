# Category 25: Queue (Redis)

Message queue with Redis.

**Total Points: 120**

---

## 1. Redis Connection (10 points)

Connect to Redis.

**Expected output:**
```
Redis connected: PONG
```

**Hint:** redis.NewClient, Ping.

---

## 2. Enqueue (10 points)

Push to queue.

**Expected output:**
```
Job queued
```

**Hint:** LPUSH to list.

---

## 3. Dequeue (10 points)

Pop from queue.

**Expected output:**
```
Job: {type: "email"}
```

**Hint:** BRPOP from list.

---

## 4. Worker Pattern (15 points)

Process jobs.

**Expected output:**
```
Processing job: email
```

**Hint:** Infinite loop + BRPOP.

---

## 5. Job Types (10 points)

Handle job types.

**Expected output:**
```
Email job processed
```

**Hint:** Type field in job.

---

## 6. Retry Logic (15 points)

Retry failed jobs.

**Expected output:**
```
Retrying job
```

**Hint:** Requeue on failure.

---

## 7. Delayed Jobs (15 points)

Schedule jobs.

**Expected output:**
```
Job scheduled for later
```

**Hint:** Sorted set with timestamp.

---

## 8. Job Status (15 points)

Track job status.

**Expected output:**
```
Status: completed
```

**Hint:** Store status in Redis.

---

## 9. Pub/Sub (15 points)

Publish/Subscribe.

**Expected output:**
```
Message received
```

**Hint:** Subscribe to channel.

---

## 10. Multiple Workers (15 points)

Scale workers.

**Expected output:**
```
Worker 1 processing
Worker 2 processing
```

**Hint:** Multiple goroutines.