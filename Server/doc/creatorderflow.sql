 CREATE TABLE orderflow (
         orderid CHAR(32) NOT NULL ,
         flow VARCHAR(1024)
)ENGINE = INNODB DEFAULT CHARSET = utf8mb4;