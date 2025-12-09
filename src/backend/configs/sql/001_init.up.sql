-- Create tables
CREATE TABLE IF NOT EXISTS hosts (
    id VARCHAR(36) PRIMARY KEY,
    url VARCHAR (1000) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS health_logs (
    id VARCHAR(36) PRIMARY KEY,
    host_id VARCHAR(36) NOT NULL,
    status VARCHAR (36) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    CONSTRAINT fk_health_logs_hosts
        FOREIGN KEY (host_id)
        REFERENCES hosts(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);
