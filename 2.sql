create table users(
id int primary key auto_increment,
username varchar(100) not null,
password varchar(100) not null,
email varchar(100) not null,
phonenum varhcar(100) not null);

CREATE TABLE `books` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `author` varchar(100) NOT NULL,
  `price` float NOT NULL,
  `sales` int DEFAULT NULL,
  `stock` int DEFAULT NULL,
  `classification` varchar(100) DEFAULT NULL,
  `publisher` varchar(100) DEFAULT NULL,
  `imgpath` varchar(100) DEFAULT NULL,
  `ebook` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `title` (`title`)
);

CREATE TABLE sessions(
sessionid VARCHAR(100) PRIMARY KEY,
username VARCHAR(100) NOT NULL,
userid INT NOT NULL,
FOREIGN KEY(userid) REFERENCES users(id)
);

CREATE TABLE carts(
id VARCHAR(100) PRIMARY KEY,
totalcount INT NOT NULL,
totalamount DOUBLE(19,2) NOT NULL,
userid INT NOT NULL,
FOREIGN KEY(userid) REFERENCES users(id)
);

CREATE TABLE cartitmes(
id INT PRIMARY KEY AUTO_INCREMENT,
count INT NOT NULL,
amount DOUBLE(19,2) NOT NULL,
bookid INT NOT NULL,
cartid VARCHAR(100) NOT NULL,
FOREIGN KEY(bookid) REFERENCES books(id),
FOREIGN KEY(cartid) REFERENCES carts(id)
);

CREATE TABLE orders(
id VARCHAR(100) PRIMARY KEY,
createtime DATETIME NOT NULL,
totalcount INT NOT NULL,
totalamount DOUBLE(19,2) NOT NULL,
state INT NOT NULL,
userid INT,
FOREIGN KEY(userid) REFERENCES users(id)
);

CREATE TABLE orderitems(
id INT PRIMARY KEY AUTO_INCREMENT,
COUNT INT NOT NULL,
amount DOUBLE(19,2) NOT NULL,
title VARCHAR(100) NOT NULL,
author VARCHAR(100) NOT NULL,
price DOUBLE(19,2) NOT NULL,
classification VARCHAR(100) NOT NULL,
publisher VARCHAR(100) NOT NULL,
imgpath VARCHAR(100) NOT NULL,
orderid VARCHAR(100) NOT NULL,
FOREIGN KEY(orderid) REFERENCES orders(id)
);
create table roots(
id int primary key not null auto_increment,
username varchar(100) not null unique,
password varchar(100) not null
);
