create database bookstore;
use bookstore;
show tables;

create table users
(
    id       int primary key auto_increment,
    username varchar(100) not null unique,
    PASSWORD varchar(100) not null,
    email    varchar(100)
);


insert into users(username, PASSWORD, email) values('admin', '12345678', 'admin@panda.com');

select * from users;

delete from users where username = 'test_user_01';
drop table users;