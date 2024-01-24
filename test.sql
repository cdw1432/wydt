/* user */
CREATE TABLE `users` (
	`handle` varchar(255),
	`password` varchar(255)
);
/* object */
CREATE TABLE objects (
    handle VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    start_at VARCHAR(255),
    end_at VARCHAR(255)
);

select * from objects
select * from users

SET SQL_SAFE_UPDATES = 0;

DELETE FROM users;