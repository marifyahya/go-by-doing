# Category 16: Packages

Learn package organization.

**Total Points: 100**

---

## 1. Create Package (5 points)

Create custom package.

**Expected output:**
```
Utils loaded
```

**Hint:** package utils in separate folder.

---

## 2. Import Package (5 points)

Import and use package.

**Expected output:**
```
Hello from utils
```

**Hint:** import "path/to/package"

---

## 3. Export Function (5 points)

Export function from package.

**Expected output:**
```
Exported function called
```

**Hint:** Start with capital letter.

---

## 4. Unexported Function (5 points)

Keep function private.

**Expected output:**
```
Only accessible in package
```

**Hint:** Start with lowercase.

---

## 5. Package Init (10 points)

Run initialization code.

**Expected output:**
```
Init called
```

**Hint:** func init() { }

---

## 6. Package Variables (10 points)

Use package-level variables.

**Expected output:**
```
Config loaded
```

**Hint:** Define var at package level.

---

## 7. Package Alias (5 points)

Use package alias.

**Expected output:**
```
Called as u
```

**Hint:** import u "utils"

---

## 8. Dot Import (10 points)

Import without prefix.

**Expected output:**
```
Direct call
```

**Hint:** import . "package"

---

## 9. Blank Import (5 points)

Use blank identifier.

**Expected output:**
```
Driver registered
```

**Hint:** import _ "driver"

---

## 10. Multiple Files (10 points)

Split code across files.

**Expected output:**
```
Functions from multiple files
```

**Hint:** Multiple .go files same package.

---

## 11. Subpackages (10 points)

Create nested packages.

**Expected output:**
```
subpackage function
```

**Hint:** package/subpackage

---

## 12. Go Modules (10 points)

Use go mod.

**Expected output:**
```
Module: github.com/user/project
```

**Hint:** go mod init module name

---

## 13. External Packages (10 points)

Import external packages.

**Expected output:**
```
Package version: v1.0.0
```

**Hint:** go get package

---

## 14. Vendor (10 points)

Use vendor directory.

**Expected output:**
```
Using vendored package
```

**Hint:** go mod vendor