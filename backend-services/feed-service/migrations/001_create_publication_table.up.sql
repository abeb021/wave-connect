CREATE TABLE publications (
    post_id UUID PRIMARY KEY,
    text TEXT NOT NULL,
    user_id TEXT NOT NULL,
    time_created TIMESTAMPTZ NOT NULL DEFAULT NOW()
);