# Wave Connect API Endpoints

## Gateway (recommended entrypoint, port 8080)

Authentication model:

- Public routes: `POST /api/auth/register`, `POST /api/auth/login`
- All other `/api/auth/*`, `/api/feed/*`, `/api/profile/*`, `/api/chat/*` require JWT
- JWT is read from `Authorization: Bearer <token>` (or `jwt` cookie)
- Gateway forwards caller as `X-User-ID` to downstream services

| Method | Endpoint | Body | Description |
|--------|----------|------|-------------|
| POST | `/api/auth/register` | `{ "email": string, "password": string }` | Register new user. Returns user object. |
| POST | `/api/auth/login` | `{ "email": string, "password": string }` | Login. Returns JWT in `Authorization: Bearer <token>` header (HTTP 202). |
| GET | `/api/auth/id/` | - | Get current user by JWT subject (requires JWT). |
| DELETE | `/api/auth/` | - | Delete current user by JWT subject (requires JWT). |
| POST | `/api/feed/` | `{ "text": string }` | Create publication (requires JWT). |
| GET | `/api/feed/` | - | Get feed list (requires JWT). |
| GET | `/api/feed/user/{userID}` | - | Get publications for a user (requires JWT). |
| GET | `/api/feed/{id}` | - | Get publication by ID (requires JWT). |
| PUT | `/api/feed/{id}` | `{ "text": string }` | Update own publication text (requires JWT). |
| DELETE | `/api/feed/{id}` | - | Delete own publication (requires JWT). |
| POST | `/api/feed/{pubID}/comment/` | `{ "text": string }` | Create comment on publication (requires JWT). |
| GET | `/api/feed/{pubID}/comment/` | - | Get comments for publication (requires JWT). |
| DELETE | `/api/feed/comment/{id}` | - | Delete own comment (requires JWT). |
| POST | `/api/profile/` | `{ "username": string }` | Create profile (requires JWT). |
| GET | `/api/profile/{id}` | - | Get profile by ID (requires JWT). |
| GET | `/api/profile/username/{username}` | - | Get profile by username (requires JWT). |
| PUT | `/api/profile/` | `{ "username": string, "bio": string }` | Update own profile (requires JWT). |
| DELETE | `/api/profile/` | - | Delete own profile (requires JWT). |
| PUT | `/api/profile/avatar/` | raw bytes (max 5MB) | Update own avatar (requires JWT). |
| GET | `/api/profile/avatar/{id}` | - | Get avatar bytes by user ID (requires JWT). |
| GET | `/api/chat/conversation` | - | Get full conversation list for current user (requires JWT). |
| GET | `/api/chat/conversation/{peerID}` | - | Get conversation with specific peer (requires JWT). |
| GET | `/api/chat/{id}` | - | Get message by ID (requires JWT). |
| PUT | `/api/chat/{id}` | `{ "text": string }` | Update own message text (requires JWT). |
| DELETE | `/api/chat/{id}` | - | Delete own message (requires JWT). |
| GET | `/api/chat/ws` | - | WebSocket endpoint (requires JWT). |

Notes:

- `POST /api/chat/` exists in chat handler code but is not wired in routes.

## Direct Service Endpoints (Docker ports)

### Auth Service (port 8081)

| Method | Endpoint | Body | Description |
|--------|----------|------|-------------|
| POST | `/api/auth/register` | `{ "email": string, "password": string }` | Register new user. |
| POST | `/api/auth/login` | `{ "email": string, "password": string }` | Login and return JWT in response `Authorization` header (HTTP 202). |
| GET | `/api/auth/id/` | - | Get current user by `X-User-ID`. |
| DELETE | `/api/auth/` | - | Delete current user by `X-User-ID`. |

### Feed Service (port 8083)

| Method | Endpoint | Body | Description |
|--------|----------|------|-------------|
| POST | `/api/feed/` | `{ "text": string }` | Create publication. |
| GET | `/api/feed/` | - | Get feed list with joined profile projection. |
| GET | `/api/feed/user/{userID}` | - | Get user publications. |
| GET | `/api/feed/{id}` | - | Get publication by ID. |
| PUT | `/api/feed/{id}` | `{ "text": string }` | Update own publication text. |
| DELETE | `/api/feed/{id}` | - | Delete own publication. |
| POST | `/api/feed/{pubID}/comment/` | `{ "text": string }` | Create comment. |
| GET | `/api/feed/{pubID}/comment/` | - | List comments for publication. |
| DELETE | `/api/feed/comment/{id}` | - | Delete own comment. |

### Profile Service (port 8084)

| Method | Endpoint | Body | Description |
|--------|----------|------|-------------|
| POST | `/api/profile/` | `{ "username": string }` | Create profile. |
| GET | `/api/profile/{id}` | - | Get profile by ID. |
| GET | `/api/profile/username/{username}` | - | Get profile by username. |
| PUT | `/api/profile/` | `{ "username": string, "bio": string }` | Update own profile. |
| DELETE | `/api/profile/` | - | Delete own profile. |
| PUT | `/api/profile/avatar/` | raw bytes (max 5MB) | Update avatar. |
| GET | `/api/profile/avatar/{id}` | - | Get avatar bytes. |

### Chat Service (port 8082)

| Method | Endpoint | Body | Description |
|--------|----------|------|-------------|
| GET | `/health` | - | Health endpoint. |
| GET | `/api/chat/conversation` | - | Get all user conversations/messages. |
| GET | `/api/chat/conversation/{peerID}` | - | Get conversation with peer. |
| GET | `/api/chat/{id}` | - | Get message by ID. |
| PUT | `/api/chat/{id}` | `{ "text": string }` | Update own message text. |
| DELETE | `/api/chat/{id}` | - | Delete own message. |
| GET | `/api/chat/ws` | websocket frames | WebSocket endpoint for realtime chat. |

## Base URLs (Docker)

- Gateway: `http://localhost:8080`
- Auth: `http://localhost:8081`
- Chat: `http://localhost:8082`
- Feed: `http://localhost:8083`
- Profile: `http://localhost:8084`
