CREATE TABLE profiles (
    user_id UUID PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    time_created TIMESTAMPTZ NOT NULL DEFAULT NOW()
);