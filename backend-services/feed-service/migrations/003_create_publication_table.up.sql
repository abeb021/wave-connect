CREATE TABLE profiles (
    id UUID PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    bio TEXT NOT NULL DEFAULT '',
    avatar BYTEA,
    time_created TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE processed_events (
    event_id TEXT PRIMARY KEY,
    event_type TEXT NOT NULL,
    processed_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);