ALTER TABLE users
ADD discord_id VARCHAR(255) UNIQUE,
    access_token VARCHAR(255) UNIQUE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now();