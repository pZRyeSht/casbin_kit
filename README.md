# casbin_kit
the casbin kit sample code!!!
## mysql initialize
```mysql
CREATE TABLE IF NOT EXISTS `casbin_rule`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `ptype` varchar(100) DEFAULT NULL,
    `v0` varchar(100) DEFAULT NULL,
    `v1` varchar(100) DEFAULT NULL,
    `v2` varchar(100) DEFAULT NULL,
    `v3` varchar(100) DEFAULT NULL,
    `v4` varchar(100) DEFAULT NULL,
    `v5` varchar(100) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES
	('p', '0', '/api/v1/casbin', 'POST', '', '', ''),
	('p', '0', '/api/v1/casbin/list', 'POST', '', '', ''),
	('p', '0', '/api/v1/health', 'GET', '', '', '');
```
