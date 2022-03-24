CREATE DATABASE golessons;

CREATE TABLE employees (
  name       VARCHAR(100),
  email      VARCHAR(255),
  mobileno   VARCHAR(100) NOT NULL,
  dept       VARCHAR(100),
  birthdate  DATE,
  createdat  TIMESTAMP,
  UNIQUE INDEX(email)
) ENGINE=INNODB;


CREATE USER 'dbuser'@localhost IDENTIFIED BY 'user2022';

GRANT ALL PRIVILEGES ON golessons.* TO 'dbuser'@localhost;

FLUSH PRIVILEGES;

SHOW GRANTS FOR 'dbuser'@localhost;

DROP USER 'dbuser'@localhost;