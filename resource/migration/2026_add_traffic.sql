CREATE TABLE IF NOT EXISTS traffic_plans (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    monthly_limit BIGINT NOT NULL,
    auto_action TEXT DEFAULT 'none',
    notify BOOLEAN DEFAULT 1
);

ALTER TABLE servers ADD COLUMN traffic_plan_id INTEGER DEFAULT NULL;

CREATE TABLE IF NOT EXISTS traffic_usage (
    server_id INTEGER,
    year INTEGER,
    month INTEGER,
    rx BIGINT DEFAULT 0,
    tx BIGINT DEFAULT 0,
    exceeded BOOLEAN DEFAULT 0,
    PRIMARY KEY (server_id, year, month)
);
