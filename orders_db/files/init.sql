create database orders;
use orders;

CREATE TABLE order_list
(
order_id INTEGER NOT NULL AUTO_INCREMENT,
user_id INTEGER,
total DECIMAL(12, 2),
qty INTEGER,
primary key (order_id)
) COMMENT='Lists orders placed by users';

CREATE TABLE order_products
(
order_id INTEGER,
product_id INTEGER,
qty INTEGER
) COMMENT = "Lists products purchased per order";

CREATE TABLE order_emails
(
order_id INTEGER,
email TEXT,
first_name TEXT,
last_name TEXT
) COMMENT = "Lists emails for each order";
