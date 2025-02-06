CREATE TABLE wallets (
    address TEXT PRIMARY KEY,
    balance DECIMAL(10, 2) NOT NULL
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    from_address REFERENCES wallets(address) NOT NULL,
    to_address REFERENCES wallets(address) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);