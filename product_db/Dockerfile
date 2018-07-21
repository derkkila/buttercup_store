FROM mysql
MAINTAINER  (me) <email>

# Copy the database schema to the /data directory
ADD files/init.sql /docker-entrypoint-initdb.d

ENV MYSQL_ROOT_PASSWORD=test

EXPOSE 3306
