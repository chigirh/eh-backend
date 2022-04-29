CREATE TABLE IF NOT EXISTS `users`(
    `user_id` VARCHAR(64) comment 'user id',
    `first_name` VARCHAR(300) NOT NULL,
    `family_name` VARCHAR(300) NOT NULL,
    PRIMARY KEY(`user_id`)
);

CREATE TABLE IF NOT EXISTS `roles`(
    `user_id` VARCHAR(64) comment 'user id',
    `role` VARCHAR(10) comment 'ADMIN,CORP,GENE',
    PRIMARY KEY(`user_id`,`role`)
);

CREATE TABLE IF NOT EXISTS `passwords`(
    `user_id` VARCHAR(64) comment 'user id',
    `password` text NOT NULL comment 'password is encrypted of SHA-256',
    PRIMARY KEY(`user_id`),
    FOREIGN KEY `fk_users`(`user_id`) REFERENCES `users`(`user_id`)
);

CREATE TABLE IF NOT EXISTS `schedules`(
    `schedule_id` VARCHAR(64),
    `user_id` VARCHAR(64),
    `date` DATE NOT NULL,
    `period` INTEGER,
    PRIMARY KEY(`schedule_id`),
    UNIQUE `unique_idx` (`user_id`, `date`, `period`),
    INDEX `date_idx` (`date`, `period`),
    FOREIGN KEY `fk_users`(`user_id`) REFERENCES `users`(`user_id`),
    FOREIGN KEY `fk_m_schedule`(`period`) REFERENCES `m_schedule`(`period`)
);
