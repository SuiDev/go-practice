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
