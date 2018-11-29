create database demo;

CREATE TABLE `reservations` (
  `resid` int(11) NOT NULL AUTO_INCREMENT,
  `vdsid` varchar(32) DEFAULT NULL,
  `shipcode` enum('AL','OA','SY','VI','HM') DEFAULT NULL,
  `saildate` date DEFAULT NULL,
  PRIMARY KEY (`resid`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

