# Go Backend Project

This is a simple Go backend project that implements basic functionalities for user management, transaction processing, and balance handling. The project is structured to follow a modular approach, making it easy to understand and extend.

## Project Structure

```
go-backend-project
├── cmd
│   └── main.go                # Entry point of the application
├── config
│   └── config.go             # Configuration settings and environment variable loading
├── internal
│   ├── database
│   │   ├── migrations
│   │   │   ├── 001_create_users.sql      # SQL schema for users table
│   │   │   ├── 002_create_transactions.sql # SQL schema for transactions table
│   │   │   ├── 003_create_balances.sql    # SQL schema for balances table
│   │   │   └── 004_create_audit_logs.sql  # SQL schema for audit logs table
│   │   └── db.go                # Database connection and migration handling
│   ├── models
│   │   ├── balance.go           # Balance struct and related methods
│   │   ├── transaction.go       # Transaction struct and related methods
│   │   ├── user.go              # User struct and related methods
│   │   └── audit_log.go         # AuditLog struct for logging actions
│   ├── services
│   │   ├── auth_service.go      # UserService for registration, login, and role checks
│   │   ├── balance_service.go    # BalanceService for balance updates and retrieval
│   │   └── transaction_service.go # TransactionService for processing transactions
│   └── worker
│       └── worker_pool.go       # Worker pool for concurrent transaction processing
├── pkg
│   ├── logger
│   │   └── logger.go            # Simple logging system
│   └── middleware
│       ├── auth_middleware.go   # Middleware for authentication checks
│       ├── error_handler.go      # Middleware for error handling
│       └── validation_middleware.go # Middleware for request validation
├── go.mod                       # Go module definition
├── go.sum                       # Checksums for module dependencies
└── README.md                    # Project documentation
```

## Setup Instructions

1. **Clone the Repository**: 
   Clone this repository to your local machine.

2. **Install Dependencies**: 
   Run the following command to install the necessary dependencies:
   ```
   go mod tidy
   ```

3. **Set Up Environment Variables**: 
   Create a `.env` file in the root directory and define your environment variables (e.g., database connection string).

4. **Run Migrations**: 
   Ensure your database is set up and run the SQL migration files located in `internal/database/migrations` to create the necessary tables.

5. **Start the Application**: 
   Use the following command to run the application:
   ```
   go run cmd/main.go
   ```

## Usage

- **Authentication Endpoints**:
  - `POST /auth/register`: Register a new user.
  - `POST /auth/login`: Log in an existing user.
  - `POST /auth/refresh`: Refresh authentication tokens.

- **User Management**:
  - `GET /users`: Retrieve user information.
  - `PUT /users`: Update user information.
  - `DELETE /users`: Delete a user.

- **Transaction Management**:
  - `POST /transactions/credit`: Credit an amount to a user.
  - `POST /transactions/debit`: Debit an amount from a user.
  - `POST /transactions/transfer`: Transfer an amount between users.
  - `GET /transactions/history`: Retrieve transaction history.

- **Balance Management**:
  - `GET /balances/current`: Get the current balance of a user.
  - `GET /balances/historical`: Get historical balance information.

## Contributing

Feel free to contribute to this project by submitting issues or pull requests. Your feedback and contributions are welcome!