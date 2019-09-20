     CREATE TABLE material (
         materid int unsigned NOT NULL AUTO_INCREMENT primary key,
         materdes CHAR(254)  NOT NULL unique key ,
         cid CHAR(32),
         customname CHAR(64),
         plating  CHAR(32),
         friction CHAR(32),
         thickness CHAR(32),
         salt CHAR(32),
         componentsolid CHAR(32),
         componentformat CHAR(32),
         factory CHAR(16),
         unit CHAR(16),
         money double,
         creattime  CHAR(32),
		 creatstamp BIGINT NOT NULL,
		 lastTime BIGINT NOT NULL
     )ENGINE = INNODB DEFAULT CHARSET = utf8mb4;material