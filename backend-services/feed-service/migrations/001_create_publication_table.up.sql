CREATE TABLE publications (
    id UUID PRIMARY KEY,
    text TEXT NOT NULL,
    user_id TEXT NOT NULL,
    time_created TIMESTAMPTZ NOT NULL DEFAULT NOW()
);