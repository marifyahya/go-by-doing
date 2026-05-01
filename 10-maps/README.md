# Category 10: Maps

Learn key-value storage in Go.

**Total Points: 180**

---

## 1. Declare Map (5 points)

Create an empty map.

**Expected output:**
```
map[]
```

**Hint:** `var m map[string]int` or `make(map[string]int)`

---

## 2. Initialize Map (5 points)

Initialize with values.

**Expected output:**
```
map[alice:25 bob:30]
```

**Hint:** `map[string]int{"alice": 25, "bob": 30}`

---

## 3. Add Element (5 points)

Add key-value to map.

**Expected output:**
```
map[alice:25]
map[alice:25 bob:30]
```

**Hint:** m["key"] = value

---

## 4. Get Value (5 points)

Retrieve value by key.

**Expected output:**
```
Alice: 25
Charlie: (not found)
```

**Hint:** m["key"], check exists with comma-ok.

---

## 5. Update Value (10 points)

Update existing value.

**Expected output:**
```
Before: alice: 25
After: alice: 26
```

**Hint:** m["key"] = newValue

---

## 6. Delete Key (5 points)

Remove key from map.

**Expected output:**
```
map[alice:25]
After delete: map[]
```

**Hint:** delete(map, key)

---

## 7. Check Key Exists (10 points)

Check if key exists.

**Expected output:**
```
alice exists: true
charlie exists: false
```

**Hint:** value, exists := m["key"]

---

## 8. Iterate Map (10 points)

Loop through map.

**Expected output:**
```
name: alice, age: 25
name: bob, age: 30
```

**Hint:** for key, value := range m { }

---

## 9. Map Length (5 points)

Get number of keys.

**Expected output:**
```
Length: 2
```

**Hint:** len(m)

---

## 10. Nested Map (10 points)

Create nested maps.

**Expected output:**
```
Alice's city: New York
```

**Hint:** map[string]map[string]string

---

## 11. Map of Slices (10 points)

Map with slice values.

**Expected output:**
```
fruits: [apple orange]
```

**Hint:** map[string][]string

---

## 12. Map of Structs (15 points)

Map with struct values.

**Expected output:**
```
User: Alice, Email: alice@test.com
```

**Hint:** map[string]struct{...}

---

## 13. Count Frequency (15 points)

Count occurrences.

**Expected output:**
```
a: 3, b: 2, c: 1
```

**Hint:** Use map to track counts.

---

## 14. Invert Map (15 points)

Swap keys and values.

**Expected output:**
```
Original: {a:1, b:2}
Inverted: {1:a, 2:b}
```

**Hint:** Iterate and swap.

---

## 15. Filter Map (15 points)

Filter map by condition.

**Expected output:**
```
Original: {a:1, b:2, c:3}
Filtered (>1): {b:2, c:3}
```

**Hint:** Create new map with condition.

---

## 16. Map Keys (10 points)

Get all keys.

**Expected output:**
```
[alice bob charlie]
```

**Hint:** Iterate and collect keys.

---

## 17. Map Values (10 points)

Get all values.

**Expected output:**
```
[25 30 35]
```

**Hint:** Iterate and collect values.

---

## 18. Merge Maps (10 points)

Combine two maps.

**Expected output:**
```
{first: 1, second: 2, third: 3}
```

**Hint:** Iterate and add to main map.

---

## 19. Group by (20 points)

Group items by property.

**Expected output:**
```
Adult: [Alice Bob]
Child: [Charlie]
```

**Hint:** Use map of slices.

---

## 20. Remove Duplicates (15 points)

Remove duplicate strings.

**Expected output:**
```
Input: [a, b, a, c, b]
Unique: [a b c]
```

**Hint:** Use map to track unique.