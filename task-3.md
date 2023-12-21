# Setup MySQL

```bash
docker run --name database-access -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql

docker exec -it database-access /bin/sh
sh-4.4# mysql -psecret
mysql> create database gowiki;
mysql> use gowiki;

mysql> CREATE TABLE page (
  id      INT AUTO_INCREMENT NOT NULL,
  title   VARCHAR(255) NOT NULL,
  body    TEXT NOT NULL,
  author  VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);

mysq> INSERT INTO page
  (title, body, author)
VALUES
  ('Blue Train', 'Body - Blue Train', 'John Coltrane'),
  ('Giant Steps', 'Body - Giant Steps', 'John Coltrane'),
  ('Jeru', 'Body - Jeru', 'Gerry Mulligan'),
  ('Sarah Vaughan', 'Body - Sarah Vaughan', 'Sarah Vaughan');

mysql> SELECT * from page;
```
