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

CREATE TABLE entries (
    id SERIAL PRIMARY KEY,
    user_id int,
    job_id int,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (job_id) REFERENCES jobs(id)
);

INSERT INTO users (name) VALUES
('Alice'),
('Bob'),
('Charlie'),
('David'),
('Eva'),
('Frank'),
('Grace'),
('Hannah'),
('Ivy'),
('Jack'),
('Kelly'),
('Leo'),
('Mona'),
('Nina'),
('Oscar');


INSERT INTO jobs (name, description) VALUES
('Engineer', 'Responsible for developing software.'),
('Designer', 'Creates design specifications for software.'),
('Manager', 'Manages the project and team.'),
('Analyst', 'Analyzes and interprets data to guide decision making.'),
('Developer', 'Develops software solutions.'),
('Tester', 'Tests software to identify bugs and issues.'),
('Administrator', 'Maintains and supports IT systems and infrastructure.'),
('Consultant', 'Provides expert advice in a particular area.'),
('Support Specialist', 'Provides technical support and troubleshooting.'),
('Architect', 'Designs and plans the structure of software systems.'),
('Researcher', 'Conducts research to advance knowledge in a field.'),
('Writer', 'Creates content for various types of publications.'),
('Salesperson', 'Sells products or services to customers.'),
('Marketing Specialist', 'Develops strategies to promote products or services.'),
('Project Coordinator', 'Assists in managing and coordinating projects.');

INSERT INTO entries (user_id, job_id) VALUES
(1, 1), (1, 2), (1, 3), (1, 4), (1, 5),
(2, 6), (2, 7), (2, 8), (2, 9), (2, 10),
(3, 11), (3, 12), (3, 13), (3, 14), (3, 15),
(4, 1), (4, 6), (4, 11), (4, 5), (4, 10),
(5, 2), (5, 7), (5, 12), (5, 3), (5, 15),
(6, 8), (6, 13), (6, 4), (6, 9), (6, 14),
(7, 2), (7, 10), (7, 5), (7, 13), (7, 11),
(8, 3), (8, 9), (8, 14), (8, 6), (8, 12),
(9, 4), (9, 15), (9, 7), (9, 11), (9, 8),
(10, 5), (10, 14), (10, 12), (10, 3), (10, 9),
(11, 6), (11, 13), (11, 2), (11, 10), (11, 7),
(12, 7), (12, 4), (12, 11), (12, 15), (12, 8),
(13, 1), (13, 12), (13, 5), (13, 14), (13, 3),
(14, 8), (14, 9), (14, 6), (14, 15), (14, 2),
(15, 11), (15, 4), (15, 13), (15, 12), (15, 1);