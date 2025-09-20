
CREATE TABLE transactions (
    transaction_id VARCHAR(100) PRIMARY KEY,
    user_id VARCHAR(100) REFERENCES users(id),
    amount DECIMAL(15,2),
    transaction_date TIMESTAMP,
    location_city VARCHAR(100),
    location_country VARCHAR(50)
);

CREATE TABLE transaction_analysis (
    id SERIAL PRIMARY KEY,
    transaction_id VARCHAR(100) REFERENCES transactions(transaction_id),
    risk_score INT,
    risk_level VARCHAR(20),
    approved BOOLEAN
);

CREATE TABLE transaction_alerts (
    id SERIAL PRIMARY KEY,
    analysis_id INT REFERENCES transaction_analysis(id),
    message TEXT
    type TEXT
    priority TEXT
);