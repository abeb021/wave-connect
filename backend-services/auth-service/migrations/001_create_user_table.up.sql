CREATE TABLE users (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    password TEXT NOT NULL,
    email TEXT NOT NULL,
    time_created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
);