     CREATE TABLE goods (
         gid int unsigned NOT NULL AUTO_INCREMENT primary key,
         gname CHAR(32)  NOT NULL  unique key ,
         gtype CHAR(16),
         gformat CHAR(64),
         num TINYINT,
         unit CHAR(16),
         sid CHAR(32),
         suppliername CHAR(32),
         creattime CHAR(32) NOT NULL,
         creatstamp BIGINT NOT NULL,
         lasttime  BIGINT NOT NULL
     )ENGINE = INNODB DEFAULT CHARSET = utf8mb4;