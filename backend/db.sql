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

CREATE TABLE Post (
    post_uuid VARCHAR(36) PRIMARY KEY,
    post_creator VARCHAR(36) NOT NULL,
    post_content VARCHAR(280) NOT NULL,
    time TIMESTAMP NOT NULL,
    FOREIGN KEY (post_creator) REFERENCES User(uuid) ON DELETE CASCADE
);
