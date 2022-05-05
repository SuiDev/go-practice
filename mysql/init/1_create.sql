CREATE DATABASE golang_db;
use golang_db;

CREATE TABLE users (
    id INT(16) AUTO_INCREMENT NOT NULL,
    uuid VARCHAR(64) NOT NULL UNIQUE,
    name VARCHAR(64),
    email VARCHAR(64),
    password VARCHAR(64),
    create_at DATETIME,
    PRIMARY KEY (id)
);

CREATE TABLE todos (
    id INT(16) AUTO_INCREMENT NOT NULL,
    content TEXT,
    user_id INT(16),
    create_at DATETIME,
    PRIMARY KEY (id)
);
