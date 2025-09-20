CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    name TEXT,
    profile TEXT,
    email TEXT,
    phone TEXT,
    created_at TIMESTAMP,
    status TEXT
);

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    user_id TEXT REFERENCES users(id),
    amount NUMERIC,
    type TEXT,
    direction TEXT,
    city TEXT,
    country TEXT,
    latitude NUMERIC,
    longitude NUMERIC,
    device_id TEXT,
    timestamp TIMESTAMP
);

CREATE TABLE IF NOT EXISTS alerts (
    id SERIAL PRIMARY KEY,
    transaction_id INT REFERENCES transactions(id),
    type TEXT,
    priority TEXT,
    message TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);
