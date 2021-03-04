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



create table books
(
    id       int primary key auto_increment,
    title    varchar(100)  not null,
    author   varchar(100)  not null,
    price    double(11, 2) not null,
    sales    int           not null,
    stock    int           not null,
    img_path varchar(100)
);

select * from books;

INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('1984 (Signet Classics)', 'George Orwell', 9.99,100,100,'41aM4xOZxaL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('A Brief History of Time', 'Stephen Hawking', 18.00,100,100,'51+GySc8ExL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('A Heartbreaking Work of Staggering Genius', 'Dave Eggers', 16.95,100,100,'51em6Mq-+gL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('A Long Way Gone: Memoirs of a Boy Soldier', 'Ishmael Beah', 15.00,100,100,'51IxoPlODPL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('A Wrinkle in Time (A Wrinkle in Time Quintet)', 'Madeleine L\'Engle', 6.99,100,100,'414KmDe2xWL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('Selected Stories, 1968-1994', 'Alice Munro', 23.03,100,100,'41Zm+Hnq8TL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('All the President\'s Men', 'Bob Woodward', 21.97,100,100,'61lSD6cxtOL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('Angela\'s Ashes: A Memoir', 'Frank McCourt', 18.00,100,100,'51JhmNksrBL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('Bel Canto (Harper Perennial Deluxe Editions)', 'Ann Patchett', 18.99,100,100,'51LOpWsGN5L');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('Beloved', 'Toni Morrison', 16.00,100,100,'41Rdzbiqh7L');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('Breath, Eyes, Memory (Oprah\'s Book Club)', 'Edwidge Danticat', 15.00,100,100,'41yQjfWGKPL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('Charlie and the Chocolate Factory', 'Roald Dahl', 31.69,100,100,'51EkR6mNwEL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('Cutting for Stone', 'Abraham Verghese', 17.95,100,100,'61enXVybbjL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('Love Medicine: Newly Revised Edition (P.S.)', 'Louise Erdrich', 16.99,100,100,'41LZMNiu3xL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('Man\'s Search for Meaning', 'Viktor E. Frankl', 15.00,100,100,'41nLuZrSbNL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('Diary of a Wimpy Kid, Book 1', 'Jeff Kinney', 14.28,100,100,'51znoqd9roL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('Dune (Dune Chronicles, Book 1)', 'Frank Herbert', 18.00,100,100,'41yAdpcoZkL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('Fahrenheit 451', 'Ray Bradbury', 17.00,100,100,'41qI9quGIdL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('Gone Girl', 'Gillian Flynn', 17.00,100,100,'41cSJI7PfHL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('Goodnight Moon', 'Margaret Wise Brown', 18.99,100,100,'617F0bkL+WL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('Great Expectations (Dover Thrift Editions)', 'Charles Dickens', 6.99,100,100,'51i715XqsYL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('Guns, Germs, and Steel: The Fates of Human Societies', 'Jared Diamond Ph.D.', 29.95,100,100,'41xWuVx+LVL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('Harry Potter and the Sorcerer\'s Stone', 'J.K. Rowling', 10.99,100,100,'51HSkTKlauL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('In Cold Blood', 'Truman Capote', 16.00,100,100,'41yJPKGYz4L');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('Invisible Man', 'Ralph Ellison', 16.00,100,100,'51NIZqFLvJL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('Life After Life: A Novel', 'Kate Atkinson', 18.00,100,100,'41QDxGRAPZL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('Little House on the Prairie: Full Color Edition', 'Laura Ingalls Wilder', 16.99,100,100,'51BlvJuPBdL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('Love in the Time of Cholera (Oprah\'s Book Club)', 'Gabriel Garcia Marquez', 16.95,100,100,'41+HorgT4YL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('Love Medicine: Newly Revised Edition (P.S.)', 'Louise Erdrich', 16.99,100,100,'41LZMNiu3xL');
INSERT INTO books (title,author,price,sales,stock,img_path) VALUES('Man\'s Search for Meaning', 'Viktor E. Frankl', 15.00,100,100,'41nLuZrSbNL');
