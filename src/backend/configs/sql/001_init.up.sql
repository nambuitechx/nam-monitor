-- Create table
CREATE TABLE IF NOT EXISTS health_logs (
    id VARCHAR(36) Primary Key,
    status VARCHAR (36),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
