create database bookstore;
use bookstore;
show tables;

-- users
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


-- books
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


-- currentPageNumber, pageSize = limit (pageNumber-1)*pageSize pageSize
select * from books limit 0, 4;
select * from books limit 4, 4;
select * from books limit 8, 4;

select count(*) from books;

-- session
create table sessions
(
    session_id varchar(100) primary key,
    username   varchar(100) not null,
    user_id    int not null,
    foreign key (user_id) references users(id)
);

select * from sessions;


-- cart
create table carts
(
    id           varchar(100) primary key,
    total_count  int    not null,
    total_amount double not null,
    user_id      int    not null,
    foreign key (user_id) references users (id)
);
select * from carts;

-- cartItems
create table cart_items
(
    id      int primary key auto_increment,
    COUNT   int not null,
    amount  double not null,
    book_id int not null,
    cart_id varchar(100) not null,
    foreign key (book_id) references books(id),
    foreign key (cart_id) references carts(id)
);
select * from cart_items;

-- Orders
create table orders
(
    id           varchar(100) primary key,
    timestamp    datetime      not null,
    total_count  int           not null,
    total_amount double(11, 2) not null,
    state        int           not null,
    user_id      int           not null,
    foreign key (user_id) references users (id)
);

select * from orders;

-- Order Items
create table order_items
(
    id       int primary key auto_increment,
    count    int           not null,
    amount   double(11, 2) not null,
    title    varchar(100)  not null,
    author   varchar(100)  not null,
    price    double(11, 2) not null,
    img_path varchar(100)  not null,
    order_id varchar(100)  not null,
    foreign key (order_id) references orders (id)

);

select * from order_items;