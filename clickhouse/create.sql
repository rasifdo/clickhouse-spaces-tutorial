CREATE TABLE IF NOT EXISTS tiered_logs (
    event_time DateTime,
    level String,
    message String
) ENGINE = MergeTree
ORDER BY (event_time)
TTL toDateTime(event_time) + INTERVAL 2 MINUTE TO VOLUME 'cold'
SETTINGS storage_policy = 's3_tiered';
