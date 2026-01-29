CREATE TABLE IF NOT EXISTS messages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    text TEXT NOT NULL,
    sender TEXT NOT NULL,
    receiver TEXT NOT NULL,
    time_sent TIMESTAMPTZ NOT NULL DEFAULT NOW()
)