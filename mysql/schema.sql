create database demo;
use demo;

CREATE TABLE `ships` (
  shipcode char(2) PRIMARY KEY,
  name varchar(36) NOT NULL,
  class varchar(36) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE reservations (
  resid int(11) NOT NULL AUTO_INCREMENT,
  vdsid varchar(36) NOT NULL,
  shipcode char(2) NOT NULL, 
  saildate date DEFAULT NULL,
  PRIMARY KEY (resid),
  FOREIGN KEY rez_shipcode (shipcode)
     REFERENCES ships(shipcode)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


