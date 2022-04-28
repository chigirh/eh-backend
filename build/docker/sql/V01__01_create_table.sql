CREATE TABLE IF NOT EXISTS `users`(
    `user_id` VARCHAR(64) comment 'user id',
    `first_name` VARCHAR(300),
    `family_name` VARCHAR(300),
    PRIMARY KEY(`user_id`)
);

CREATE TABLE IF NOT EXISTS `passwords`(
    `user_id` VARCHAR(64) comment 'user id',
    `password` text comment 'password is encrypted of SHA-256',
    PRIMARY KEY(`user_id`),
    FOREIGN KEY `fk_users`(`user_id`) REFERENCES `users`(`user_id`)
);