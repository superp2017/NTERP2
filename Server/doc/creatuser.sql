     CREATE TABLE user (
         uid int unsigned NOT NULL AUTO_INCREMENT primary key,
         uname CHAR(16)  NOT NULL  unique key ,
         sex CHAR(8),
		 age     TINYINT,
         cell CHAR(11),
		 department CHAR(16),
         salary  INT,
		 loginname  CHAR(16) NOT NULL unique key,
         logincode    CHAR(16) NOT NULL,
         author TINYINT NOT NULL,
         intime CHAR(32),
         outtime CHAR(32),
         creattime CHAR(32) NOT NULL,
         creatstamp BIGINT NOT NULL,
         lasttime  BIGINT NOT NULL
     )ENGINE = INNODB DEFAULT CHARSET = utf8mb4;