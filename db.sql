DROP TABLE IF EXISTS `students`;
CREATE TABLE `students` (
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`stu_id` int(10) DEFAULT NULL,
	`cost` float(6,1) DEFAULT NULL,
	`date` varchar(10) DEFAULT NULL,
	`time` varchar(10) DEFAULT NULL,
	`restaurant` varchar(30) DEFAULT NULL,
	`place` varchar(30) DEFAULT NULL,
	PRIMARY KEY (`id`),
	KEY `stu_id` (`stu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;
