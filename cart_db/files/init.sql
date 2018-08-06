create database cart;
use cart;

CREATE TABLE cart_list
(
user_id INTEGER,
product_id INTEGER,
qty INTEGER
) COMMENT='Lists current products in users carts';
