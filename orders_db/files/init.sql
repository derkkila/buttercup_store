create database order;
use orders;

CREATE TABLE order_list
(
order_id INTEGER,
user_id INTEGER,
total DECIMAL(12, 2),
qty INTEGER
) COMMENT='Lists orders placed by users';

CREATE TABLE order_products
(
order_id INTEGER,
product_id INTEGER,
qty INTEGER
) COMMENT = "Lists products purchased per order";
