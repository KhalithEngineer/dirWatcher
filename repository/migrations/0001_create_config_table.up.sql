CREATE TABLE CONFIG 
(
    id SERIAL PRIMARY KEY 
    directory VARCHAR(255),
    magic_string VARCHAR(255),
    time_interval INTEGER
)
