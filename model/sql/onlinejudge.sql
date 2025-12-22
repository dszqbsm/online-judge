CREATE TABLE `demo`
(
    `id`      VARCHAR(128) NOT NULL,
    `demo_id` VARCHAR(32)  NOT NULL,
    CONSTRAINT `pk_id` PRIMARY KEY (`id`),
    CONSTRAINT `uc_id_demo_id` UNIQUE (`id`, `demo_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;
