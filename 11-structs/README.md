# Category 11: Structs

Learn custom data types in Go.

**Total Points: 200**

---

## 1. Define Struct (5 points)

Create a struct type.

**Expected output:**
```
Name: , Age: 0
```

**Hint:** type Person struct { ... }

---

## 2. Create Instance (5 points)

Create struct instance.

**Expected output:**
```
Alice, 25
```

**Hint:** person := Person{Name: "Alice", Age: 25}

---

## 3. Access Fields (5 points)

Access struct fields.

**Expected output:**
```
Alice
25
```

**Hint:** person.Name, person.Age

---

## 4. Pointer to Struct (10 points)

Use struct pointer.

**Expected output:**
```
Alice
25
```

**Hint:** p := &Person{Name: "Alice"}

---

## 5. Nested Struct (15 points)

Create nested struct.

**Expected output:**
```
Alice lives in New York, USA
```

**Hint:** struct inside struct.

---

## 6. Anonymous Struct (10 points)

Create anonymous struct.

**Expected output:**
```
John, 30
```

**Hint:** var person struct { ... }

---

## 7. Struct with Methods (10 points)

Add method to struct.

**Expected output:**
```
Alice (ID: 1)
```

**Hint:** func (p Person) Method() { }

---

## 8. Struct Comparison (10 points)

Compare two structs.

**Expected output:**
```
Equal: true
```

**Hint:** Use == for comparison.

---

## 9. Struct as Map Value (10 points)

Use struct as map value.

**Expected output:**
```
Alice: {Name:Alice Age:25 Email:alice@test.com}
```

**Hint:** map[string]Person

---

## 10. Slice of Structs (10 points)

Create slice of structs.

**Expected output:**
```
[Alice Bob Charlie]
```

**Hint:** []Person{...}

---

## 11. JSON to Struct (15 points)

Unmarshal JSON to struct.

**Expected output:**
```
Alice
```

**Hint:** json.Unmarshal into struct

---

## 12. Struct to JSON (15 points)

Marshal struct to JSON.

**Expected output:**
```
{"name":"Alice","age":25}
```

**Hint:** json.Marshal from struct

---

## 13. Tags (10 points)

Use struct field tags.

**Expected output:**
```
alice
```

**Hint:** \`json:"name"\`

---

## 14. Embedding (15 points)

Use struct embedding.

**Expected output:**
```
Employee: Alice, Company: TechCorp
```

**Hint:** Embed struct in struct.

---

## 15. Constructor Pattern (15 points)

Create constructor function.

**Expected output:**
```
Created: Alice
```

**Hint:** func NewPerson() *Person { }

---

## 16. Multiple Types (10 points)

Use multiple struct types.

**Expected output:**
```
Person: Alice
Address: NYC
```

**Hint:** Different struct types.

---

## 17. Validate Struct (20 points)

Validate struct fields.

**Expected output:**
```
Valid
Invalid: name required
```

**Hint:** Create validation function.

---

## 18. Copy Struct (10 points)

Create struct copy.

**Expected output:**
```
Original: Alice
Copy: Alice
```

**Hint:** Create new instance with same values.

---

## 19. Zero Value Struct (5 points)

Understand zero struct.

**Expected output:**
```
{ 0 }
```

**Hint:** var p Person

---

## 20. Custom String (15 points)

Implement String() method.

**Expected output:**
```
Person: Alice, Age: 25
```

**Hint:** func (p Person) String() string { }