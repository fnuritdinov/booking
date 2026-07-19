CREATE TABLE IF NOT EXISTS bookings (
                         id SERIAL PRIMARY KEY,
                         user_id BIGINT,
                         movie_id BIGINT,
                         status INTEGER DEFAULT 1,
                         created_at TIMESTAMP DEFAULT NOW()
);