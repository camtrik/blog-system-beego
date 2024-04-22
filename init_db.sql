CREATE DATABASE IF NOT EXISTS blog_test;
USE blog_test

CREATE TABLE IF NOT EXISTS entries (
    id INT AUTO_INCREMENT,
    title TEXT,
    content TEXT,
    created DATETIME,
    primary key (id)
);