# casbin_kit
the casbin kit sample code!!!
## mysql initialize
```mysql
INSERT INTO `casbin_test`.`casbin_rule` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES
	('p', '0', '/api/v1/casbin', 'POST', '', '', ''),
	('p', '0', '/api/v1/casbin/list', 'POST', '', '', ''),
	('p', '0', '/api/v1/health', 'GET', '', '', '');
```
