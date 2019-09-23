user     CREATE TABLE customer (
         cid int unsigned NOT NULL AUTO_INCREMENT primary key,
         cname CHAR(32)  NOT NULL unique key,
         addr CHAR(64),
         tel  CHAR(11),
         contactname CHAR(16),
         contactcell CHAR(11),
         bankname CHAR(16),
         banknumber CHAR(32),
         bankbranch CHAR(16),
         certificatesnum CHAR(64),
         note CHAR(64),
         creattime  CHAR(32),
		 creatstamp BIGINT NOT NULL,
		 lastTime BIGINT NOT NULL
     )ENGINE = INNODB DEFAULT CHARSET = utf8mb4;
     