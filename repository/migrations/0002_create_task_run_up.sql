CREATE TABLE task_run (
    id SERIAL PRIMARY KEY,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    total_runtime INTERVAL,
    directory_path TEXT,
    interval INTERVAL,
    magic_string TEXT,
    files_added TEXT[],
    files_deleted TEXT[],
    total_occurrences INT,
    status TEXT
);
