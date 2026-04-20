# Wave Connect API Endpoints

## Gateway (recommended entrypoint, port 8080)

| Method | Endpoint | Body | Description |
|--------|----------|------|-------------|
| POST | `/api/auth/register` | `{ "username": string, "email": string, "password": string }` | Register new user. Returns user object. |
| POST | `/api/auth/login` | `{ "username": string, "email": string, "password": string }` | Login. Returns JWT in `Authorization: Bearer <token>` header. |
| GET | `/api/auth/{id}` | - | Get user by ID (requires JWT). |
| DELETE | `/api/auth/{id}` | - | Delete user (requires JWT). |
| POST | `/api/feed/` | `{ "text": string }` | Create publication (requires JWT). |
| GET | `/api/feed/{id}` | - | Get publication by ID (requires JWT). |
| PUT | `/api/feed/{id}` | `{ "text": string }` | Update publication text (requires JWT). |
| DELETE | `/api/feed/{id}` | - | Delete publication (requires JWT). |
| POST | `/api/profile/` | `{ "username": string }` | Create profile (requires JWT). |
| GET | `/api/profile/{id}` | - | Get profile by ID (requires JWT). |
| PUT | `/api/profile/{id}` | `{ "username": string }` | Update profile (requires JWT). |
| DELETE | `/api/profile/{id}` | - | Delete profile (requires JWT). |
| POST | `/api/chat/` | `{ "text": string, "sender": string, "receiver": string }` | Create message (requires JWT). |
| GET | `/api/chat/{id}` | - | Get message by ID (requires JWT). |
| PUT | `/api/chat/{id}` | `{ "text": string }` | Update message text (requires JWT). |
| DELETE | `/api/chat/{id}` | - | Delete message (requires JWT). |
| GET | `/api/chat/ws` | - | Chat websocket endpoint (requires JWT). |

## Direct Service Endpoints (Docker ports)

### Auth Service (port 8081)

| Method | Endpoint | Body | Description |
|--------|----------|------|-------------|
| POST | `/api/auth/register` | `{ "username": string, "email": string, "password": string }` | Register new user. |
| POST | `/api/auth/login` | `{ "username": string, "email": string, "password": string }` | Login and return JWT in response `Authorization` header. |
| GET | `/api/auth/{id}` | - | Get user by ID (stub). |
| DELETE | `/api/auth/{id}` | - | Delete user (stub). |

### Feed Service (port 8083)

| Method | Endpoint | Body | Description |
|--------|----------|------|-------------|
| POST | `/api/feed/` | `{ "text": string }` | Create publication. |
| GET | `/api/feed/{id}` | - | Get publication by ID. |
| PUT | `/api/feed/{id}` | `{ "text": string }` | Update publication text. |
| DELETE | `/api/feed/{id}` | - | Delete publication. |

### Profile Service (port 8084)

| Method | Endpoint | Body | Description |
|--------|----------|------|-------------|
| POST | `/api/profile/` | `{ "username": string }` | Create profile. |
| GET | `/api/profile/{id}` | - | Get profile by ID. |
| PUT | `/api/profile/{id}` | `{ "username": string }` | Update profile username. |
| DELETE | `/api/profile/{id}` | - | Delete profile. |

### Chat Service (port 8082)

| Method | Endpoint | Body | Description |
|--------|----------|------|-------------|
| POST | `/api/chat/` | `{ "text": string, "sender": string, "receiver": string }` | Create message. Returns message with id and timeSent. |
| GET | `/api/chat/{id}` | - | Get message by ID. |
| PUT | `/api/chat/{id}` | `{ "text": string }` | Update message text. |
| DELETE | `/api/chat/{id}` | - | Delete message. |
| GET | `/api/chat/ws` | - | Websocket endpoint for realtime chat. |

## Base URLs (Docker)

- Gateway: `http://localhost:8080`
- Auth: `http://localhost:8081`
- Chat: `http://localhost:8082`
- Feed: `http://localhost:8083`
- Profile: `http://localhost:8084`
