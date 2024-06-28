CREATE DATABASE sample_db;
\c sample_db;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE jobs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO users (name) VALUES
('Alice'),
('Bob'),
('Charlie');

INSERT INTO jobs (name, description) VALUES
('Engineer', 'Responsible for developing software.'),
('Designer', 'Creates design specifications for software.'),
('Manager', 'Manages the project and team.');
