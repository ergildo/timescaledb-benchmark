CREATE EXTENSION IF NOT EXISTS timescaledb;
CREATE TABLE cpu_usage(
                          ts    TIMESTAMPTZ,
                          host  TEXT,
                          usage DOUBLE PRECISION
);
