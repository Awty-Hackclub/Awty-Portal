CREATE DATABASE Awty_Portal;
USE Awty_Portal;

CREATE TABLE User (
    uuid VARCHAR(36) PRIMARY KEY,
    username VARCHAR(40) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password TEXT NOT NULL,
    description TEXT,
    UNIQUE KEY (username)
);
