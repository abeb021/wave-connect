CREATE TABLE comments (
    id UUID PRIMARY KEY,
    pub_id UUID NOT NULL REFERENCES publications(id) ON DELETE CASCADE,
    text TEXT NOT NULL,
    user_id TEXT NOT NULL,
    time_created TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
