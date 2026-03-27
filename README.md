# Go User API (Fixed Version)

## 👨‍💻 Author

Fixed and improved by **Anouar Mezgualli**

---

## 📌 About

This project was originally a broken Golang backend used for interview practice.
It had several issues like bad practices, bugs, and security problems.

I worked on fixing and improving it to make it more stable and secure.

---

## ✅ What I Fixed

* 🔐 Passwords are now hashed using bcrypt
* 🛡️ Fixed SQL injection (using parameterized queries)
* ⚙️ Improved database connection (no longer opened on every request)
* ✅ Added input validation (empty fields, weak passwords)
* 🌍 Removed hardcoded config (using `.env`)
* 📡 Improved API responses

---

## ⚡ How to Run

### 1. Start database

```bash
docker-compose up
```

### 2. Create `.env` file

```env
DB_USER=admin
DB_PASSWORD=admin
DB_NAME=test_repo
DB_HOST=localhost
DB_PORT=5432
```

### 3. Run the server

```bash
go run ./cmd
```

---

## 📬 API

### POST `/user`

#### Request:

```json
{
  "username": "john",
  "password": "123456"
}
```

---

## 🎯 Goal

This project was mainly to practice:

* Fixing bugs
* Improving code
* Applying basic backend best practices

---
