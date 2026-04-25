CREATE TABLE profiles (
    id UUID PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    bio TEXT NOT NULL DEFAULT '',
    avatar BYTEA,
    time_created TIMESTAMPTZ NOT NULL DEFAULT NOW()
);