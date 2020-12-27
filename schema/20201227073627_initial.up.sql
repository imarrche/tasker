CREATE TABLE projects (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(500) NOT NULL,
    description VARCHAR(1000)
);

CREATE TABLE columns (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    index INTEGER NOT NULL UNIQUE,
    project_id INTEGER REFERENCES projects (id) NOT NULL
);

CREATE TABLE tasks (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(500) NOT NULL,
    description VARCHAR(5000),
    index INTEGER NOT NULL UNIQUE,
    column_id INTEGER REFERENCES columns (id) NOT NULL
);

CREATE TABLE comments (
    id BIGSERIAL PRIMARY KEY,
    text VARCHAR(5000) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    task_id INTEGER REFERENCES tasks (id) NOT NULL
);