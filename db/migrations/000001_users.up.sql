CREATE TABLE users (
    id INT PRIMARY KEY,
    first_name VARCHAR(255),
    family_name VARCHAR(255),
    email VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP NULL
);