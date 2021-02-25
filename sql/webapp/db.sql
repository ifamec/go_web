show databases;
create database test;
use test;
create table users
(
    id       int primary key auto_increment,
    username varchar(100) not null unique,
    PASSWORD varchar(100) not null,
    email    varchar(100)
);
show tables;
select * from users;