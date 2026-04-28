CREATE TABLE profiles (
    id UUID PRIMARY KEY,
    username TEXT UNIQUE,
    time_created TIMESTAMPTZ NOT NULL DEFAULT NOW()
);