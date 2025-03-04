CREATE DATABASE icecreams_catalog;
CREATE DATABASE user_data;

\c icecreams_catalog
CREATE TABLE icecreams (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    ingredients TEXT NOT NULL,
    production_date VARCHAR(32) NOT NULL,
    best_before VARCHAR(32) NOT NULL,
    price REAL NOT NULL,
    quantity INTEGER DEFAULT 1 NOT NULL
);

\c user_data
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    user_name VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    secret_word_hash VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
);

\c user_data
CREATE TABLE refresh_tokens (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    token TEXT NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    revoked_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
