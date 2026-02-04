# Wave Connect API Endpoints

## Auth Service (port 8081)

| Method | Endpoint | Body | Description |
|--------|----------|------|-------------|
| POST | `/api/auth/register` | `{ "username": string, "email": string, "password": string }` | Register new user. Returns user object. |
| POST | `/api/auth/login` | `{ "username": string, "email": string, "password": string }` | Login. Returns JWT in `Authorization: Bearer <token>` header. |
| GET | `/api/auth/{id}` | - | Get user by ID (stub) |
| DELETE | `/api/auth/{id}` | - | Delete user (stub) |

## Chat Service (port 8080)

| Method | Endpoint | Body | Description |
|--------|----------|------|-------------|
| POST | `/api/message` | `{ "text": string, "sender": string, "receiver": string }` | Create message. Returns message with id & timeSent. |
| GET | `/api/message/{id}` | - | Get message by ID |
| PUT | `/api/message/{id}` | `{ "text": string }` | Update message text |
| DELETE | `/api/message/{id}` | - | Delete message |

## Base URLs (Docker)

- Auth: `http://localhost:8081`
- Chat: `http://localhost:8080`
