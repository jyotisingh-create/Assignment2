# 🏦 Banking System Backend

A banking system backend built with **Go**, **Gin**, **GORM**, and **PostgreSQL**.
This service manages banks, branches, customers, accounts, transactions, and loans while ensuring data consistency through database transactions.

---

## 📌 List Of Features

* Multi-bank & multi-branch support
* Customer registration & management
* Savings accounts with deposits & withdrawals
* Immutable transaction history
* Loan management with configurable tenure
* Fixed **12% yearly interest rate**
* Stores total payable amount at loan creation
* Loan linked with account type
* Auto-close loans when fully repaid
* RESTful APIs with proper HTTP status codes
* Database transactions for atomicity

---
## Tech Stack

- **Go 1.21+** - Backend language
- **Gin** - Web framework
- **GORM** - ORM
- **PostgreSQL** - Database
## 🚀 Quick Start

### Setup

### Install dependencies

```bash
go mod download
```

### Run Server

```bash
go run cmd/main.go
```

Server runs on:

```
http://localhost:8080
```

---

## 🔌 API Endpoints

| Method | Endpoint                    | Description             |
| ------ | --------------------------- | ----------------------- |
| POST   | /banks                      | Create bank             |
| GET    | /banks/{id}                 | Get bank with branches  |
| POST   | /branches                   | Create branch           |
| POST   | /customers                  | Register customer       |
| GET    | /customers/{id}             | Get customer profile    |
| POST   | /accounts/savings           | Open savings account    |
| GET    | /accounts/{id}              | Get account balance     |
| POST   | /accounts/{id}/deposit      | Deposit money           |
| POST   | /accounts/{id}/withdraw     | Withdraw money          |
| GET    | /accounts/{id}/transactions | Get transaction history |
| POST   | /loans                      | Take a loan             |
| GET    | /loans/{id}                 | Get loan details        |
| POST   | /loans/{id}/repay           | Repay loan              |
| GET    | /loans/{id}/interest        | Calculate interest      |

---

## 🧱 Project Structure

```
banking-system/
├── cmd/main.go
├── config/
├── models/
├── controllers/
├── services/
└── routes
```