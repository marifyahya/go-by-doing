# Category 17: Concurrency (Goroutines)

Learn concurrent programming in Go.

**Total Points: 250**

---

## 1. Basic Goroutine (5 points)

Start a goroutine.

**Expected output:**
```
Main done
(background task runs)
```

**Hint:** go function()

---

## 2. Multiple Goroutines (10 points)

Start multiple goroutines.

**Expected output:**
```
Task 1 done
Task 2 done
```

**Hint:** Multiple go statements.

---

## 3. Channel Basics (10 points)

Create and use channel.

**Expected output:**
```
Received: 42
```

**Hint:** make(chan int), <-ch, ch<-

---

## 4. Send and Receive (10 points)

Communication via channel.

**Expected output:**
```
Sent: 10
Received: 10
```

**Hint:** Two-way communication.

---

## 5. Buffered Channel (10 points)

Create buffered channel.

**Expected output:**
```
Channel with buffer 3
```

**Hint:** make(chan int, 3)

---

## 6. Close Channel (10 points)

Close channel after use.

**Expected output:**
```
Values: 1 2 3
Closed
```

**Hint:** close(ch), check ok

---

## 7. Select Statement (15 points)

Handle multiple channels.

**Expected output:**
```
Received from ch1: 1
```

**Hint:** select { case ... }

---

## 8. Select with Default (10 points)

Non-blocking select.

**Expected output:**
```
No message available
```

**Hint:** default case in select.

---

## 9. Select Timeout (15 points)

Implement timeout.

**Expected output:**
```
Timeout!
```

**Hint:** select with time.After.

---

## 10. Select Multiple Channels (15 points)

Handle many channels.

**Expected output:**
```
Received from one of the channels
```

**Hint:** Multiple case statements.

---

## 11. Worker Pool (20 points)

Create worker pool pattern.

**Expected output:**
```
Worker 1 processed job
Worker 2 processed job
```

**Hint:** Multiple workers from channel.

---

## 12. Fan Out (20 points)

Distribute work to workers.

**Expected output:**
```
Results collected from workers
```

**Hint:** Multiple goroutines processing.

---

## 13. Fan In (20 points)

Combine results.

**Expected output:**
```
Combined result
```

**Hint:** Merge channels into one.

---

## 14. Pipeline (20 points)

Chain processing stages.

**Expected output:**
```
Stage 1 -> Stage 2 -> Stage 3
```

**Hint:** Connect stages with channels.

---

## 15. Generator Pattern (15 points)

Create data generator.

**Expected output:**
```
1 2 3 4 5 ...
```

**Hint:** Return channel from function.

---

## 16. Context in Goroutine (15 points)

Pass context to goroutine.

**Expected output:**
```
Context done
```

**Hint:** Pass context to function.

---

## 17. Context Cancellation (15 points)

Cancel goroutine with context.

**Expected output:**
```
Cancelled
```

**Hint:** ctx.Done() in select.

---

## 18. Context Timeout (15 points)

Timeout with context.

**Expected output:**
```
Timed out
```

**Hint:** context.WithTimeout.

---

## 19. Context Value (10 points)

Pass values via context.

**Expected output:**
```
UserID: 123
```

**Hint:** context.WithValue.

---

## 20. Race Condition (20 points)

Demonstrate race condition.

**Expected output:**
```
Final: (variable)
```

**Hint:** Show need for synchronization.

---

## 21. Deadlock Prevention (15 points)

Avoid deadlocks.

**Expected output:**
```
No deadlock
```

**Hint:** Proper channel ordering.