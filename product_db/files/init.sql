create database products;
use products;

CREATE TABLE product_list
(
id INTEGER AUTO_INCREMENT,
name TEXT,
description TEXT,
type TEXT,
category TEXT,
price FLOAT,
qty INT,
PRIMARY KEY (id)
) COMMENT='Lists products available for purchase';

INSERT INTO product_list (name,description,type,category,price,qty)
VALUES ("A Beginner's Guide to Collectd","With a growing need to effectively monitor your infrastructure, metrics have emerged as a way to gain insight to the trends and problems within your IT environment. collectd is an open source daemon that collects system and application performance metrics. With this data, collectd then has the ability to work alongside another tool, such as Splunk, to help identify trends, issues and relationships you may not be able to observe otherwise.<br /><br />
This e-book gives you a deep dive into what collectd is and how you can begin incorporating it into your organization\'s environment. Download your complimentary copy of A Beginner\'s Guide to collectd to learn how to:<br />
<ul>
<li>Get and configure collectd</li>
<li>Analyze collectd data</li>
<li>Use Splunk and collectd together to analyze large amounts of infrastructure data</li>", "book","how-to",9.99,1000);

CREATE TABLE product_images
(
id INTEGER,
imgname TEXT,
filepath TEXT
) COMMENT='Lists images for products';

INSERT INTO product_images (id,imgname,filepath)
VALUES (1,"primary_image","/wstatic/images/a-beginners-guide-to-collectd.png");
