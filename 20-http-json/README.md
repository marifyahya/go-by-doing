# Category 20: HTTP JSON

JSON handling in HTTP.

**Total Points: 160**

---

## 1. Marshal JSON (10 points)

Convert struct to JSON.

**Expected output:**
```
{"name":"Alice","age":25}
```

**Hint:** json.Marshal()

---

## 2. Unmarshal JSON (10 points)

Parse JSON to struct.

**Expected output:**
```
Alice
```

**Hint:** json.Unmarshal()

---

## 3. JSON Tags (10 points)

Use JSON field tags.

**Expected output:**
```
{"user_name":"alice","user_age":25}
```

**Hint:** \`json:"field_name"\`

---

## 4. Decode JSON Body (15 points)

Decode request body.

**Expected output:**
```
Decoded: Alice
```

**Hint:** json.NewDecoder(r.Body)

---

## 5. Encode JSON Response (10 points)

Encode to response.

**Expected output:**
```
Content-Type: application/json
{"data":"value"}
```

**Hint:** json.NewEncoder(w)

---

## 6. JSON Array (10 points)

Handle JSON arrays.

**Expected output:**
```
[1,2,3,4,5]
```

**Hint:** []int, []string.

---

## 7. Map to JSON (10 points)

Convert map to JSON.

**Expected output:**
```
{"key1":"value1","key2":"value2"}
```

**Hint:** map[string]interface{}

---

## 8. Nested JSON (15 points)

Handle nested structures.

**Expected output:**
```
{"user":{"name":"Alice","email":"alice@test.com"}}
```

**Hint:** Nested structs.

---

## 9. Partial JSON (10 points)

Partial unmarshal.

**Expected output:**
```
Only name decoded
```

**Hint:** Partial struct fields.

---

## 10. Omit Empty (10 points)

Skip empty fields.

**Expected output:**
```
{"name":"Alice"}
```

**Hint:** omitempty tag.

---

## 11. Custom Marshal (15 points)

Custom JSON marshaling.

**Expected output:**
```
Custom output
```

**Hint:** Implement MarshalJSON.

---

## 12. Raw JSON (10 points)

Handle raw JSON.

**Expected output:**
```
Raw JSON preserved
```

**Hint:** json.RawMessage.

---

## 13. Pretty Print (10 points)

Format JSON output.

**Expected output:**
```
{
  "name": "Alice"
}
```

**Hint:** json.Indent.

---

## 14. Error JSON (10 points)

Return error as JSON.

**Expected output:**
```
{"error":"message"}
```

**Hint:** Proper error response.

---

## 15. Valid JSON (10 points)

Validate JSON.

**Expected output:**
```
Valid: true
```

**Hint:** json.Valid()

---

## 16. Streaming JSON (15 points)

Process large JSON.

**Expected output:**
```
Stream processed
```

**Hint:** Decoder with token.