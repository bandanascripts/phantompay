
# Phantom Pay

Phantom Pay is an end-to-end wallet app with auth using JWTs signed and verified using RSA with option for key rotation and uses database transactions to have rollback in our hands.

## Features

- User registration & login
- JWT authentication using **RSA** key pairs
  - Private key used to sign access tokens (stored in Redis)
  - Public key used to verify tokens
- **Key rotation** support (RSA keys are swappable in Redis)
- Middleware protection for all routes
- Database **transactions with rollback** for critical flows (e.g., wallet operations)

## Tech Stack
- Golang
- MySQL
- Gin
- Redis
- JWTs
- Gorm

## Getting Started 

1 . Clone the Repository 
``` bash
git clone https://github.com/bandanascripts/phantompay
cd phantom pay
``` 

2 . Set up Env file
``` dotenv 
PORT=your_port
DB_ADDRESS=your_db_address
```

3 . Build the app 

```bash
go build
```

🟥 Important: Make sure Redis is running locally on the default port (6379). You can start Redis using Docker or a local install before running the service.

## API Endpoints 

## 📡 API Endpoints

| Method | Endpoint              | Description                     | Auth Required |
|--------|-----------------------|---------------------------------|---------------|
| POST   | `phantompay/register`           | Register a new user             | ❌            |
| POST   | `phantompay/login`              | Login and get JWT tokens        | ❌            |
| POST   | `phantompay/wallet`      | Create a wallet for the user    | ✅            |
| POST   | `phantompay/deposit`     | Deposit money into wallet       | ✅            |
| POST   | `phantompay/withdraw`    | Withdraw money from wallet      | ✅            |
| POST   | `phantompay/transaction`    | Transfer money to another user  | ✅            |
| GET    | `phantompay/transactionhistory`| Get transaction history         | ✅            |

All authenticated routes require a valid JWT in the `Authorization` header.

## License

This project is licensed under the [MIT License](LICENSE).

Feel free to use, modify, and distribute — just leave a star if you found it helpful ⭐
