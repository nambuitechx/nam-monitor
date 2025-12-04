CREATE TABLE health_logs (
    id uuid,
    status text,
    created_at timestamp,
    PRIMARY KEY (id, created_at)
) WITH CLUSTERING ORDER BY (created_at DESC);
