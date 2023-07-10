CREATE TABLE `audit` (
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    `name` VARCHAR(50) NOT NULL,
    `date` DATETIME NULL,
    `value` INT DEFAULT 0,
    UNIQUE (`name`, `date`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `user` (
    `id`   VARCHAR(50)  NOT NULL,
    `name` VARCHAR(190) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `team` (
    `id`   VARCHAR(50)  NOT NULL,
    `name` VARCHAR(190) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `user_m2m_team` (
    `user_id` VARCHAR(50)  NOT NULL,
    `team_id` VARCHAR(50)  NOT NULL,
    PRIMARY KEY (`user_id`, `team_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `t` (
    `labels` JSON NOT NULL DEFAULT ('[]')
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;