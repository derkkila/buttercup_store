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

INSERT INTO product_list (name,description,type,category,price,qty)
VALUES ("A Beginner's Guide to Kubernetes Monitoring","Application container technology, like Kubernetes, is revolutionizing app development, bringing previously unimagined flexibility and efficiency to the development process. But a good container monitoring solution is still necessary for dynamic container-based environments to unify container data with other infrastructure data — only then can you have better contextualization and root cause analysis.<br /><br />
  In this guide, we'll define the importance and benefits of software container monitoring in container vendors today — like Docker, Kubernetes, RedHat OpenShift, or Amazon EKS. Learn what a typical and successful Kubernetes deployment looks like and how to effectively monitor its orchestration.<br /><br />
  Download your copy of A Beginner's Guide to Kubernetes Monitoring to learn:<br />
  <ul>
  <li>About containers and their benefits</li>
  <li>How to deploy Kubernetes</li>
  <li>How to monitor Kubernetes deployments</li>
  </ul>","book","education",11.99,10000);

CREATE TABLE product_images
(
id INTEGER,
imgname TEXT,
filepath TEXT
) COMMENT='Lists images for products';

INSERT INTO product_images (id,imgname,filepath)
VALUES (1,"primary_image","/wstatic/images/a-beginners-guide-to-collectd.png");

INSERT INTO product_images (id,imgname,filepath)
VALUES (2,"primary_image","/wstatic/images/a-beginners-guide-to-collectd.png");
