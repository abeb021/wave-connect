CREATE TABLE messages (
    id UUID PRIMARY KEY,
    text TEXT NOT NULL,
    sender TEXT NOT NULL,
    receiver TEXT NOT NULL,
    time_sent TIMESTAMPTZ NOT NULL DEFAULT NOW()
);