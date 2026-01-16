CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY,
    email VARCHAR(128) UNIQUE NOT NULL,
    password_hash VARCHAR(256) NOT NULL,
    email_verified BOOLEAN NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE
);