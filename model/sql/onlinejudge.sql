CREATE TABLE `user`
(
    `id`            BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `user_name`     VARCHAR(16)  NOT NULL COMMENT '登录用户名',
    `password_hash` VARCHAR(255) NOT NULL COMMENT '密码哈希值',
    `status`        TINYINT(1) NOT NULL DEFAULT 1 COMMENT '用户状态: 1-启用, 0-禁用',
    `create_time`   TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`   TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time`   TIMESTAMP NULL COMMENT '删除时间',
    CONSTRAINT `pk_user_id` PRIMARY KEY (`id`),
    CONSTRAINT `uc_user_user_name` UNIQUE (`user_name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin
  COMMENT = '用户表';

CREATE TABLE `problem`
(
    `id`           BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `key`          VARCHAR(32)  NOT NULL COMMENT '题目标识',
    `title`        VARCHAR(50)  NOT NULL COMMENT '题目标题',
    `description`  TEXT NOT NULL COMMENT '题目描述',
    `input_desc`   TEXT NOT NULL COMMENT '输入格式说明',
    `output_desc`  TEXT NOT NULL COMMENT '输出格式说明',
    `sample_cases` JSON NOT NULL COMMENT '样例输入输出',
    `difficulty`   ENUM('easy', 'medium', 'hard') NOT NULL COMMENT '题目难度',
    `time_limit`   INT UNSIGNED NOT NULL COMMENT '程序运行时间限制(单位: 毫秒)',
    `memory_limit` INT UNSIGNED NOT NULL COMMENT '程序内存限制(单位: MB)',
    `tags`         VARCHAR(50) NOT NULL COMMENT '题目标签',
    `is_published` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否发布: 1-发布, 0-草稿',
    `create_time`  TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time`  TIMESTAMP NULL COMMENT '删除时间',
    CONSTRAINT `pk_problem_id` PRIMARY KEY (`id`),
    CONSTRAINT `uc_problem_key` UNIQUE (`key`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin
  COMMENT = '题目表';

CREATE TABLE `test_case`
(
    `id`              BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `key`             VARCHAR(32)  NOT NULL COMMENT '测试用例标识',
    `problem_key`     VARCHAR(32)  NOT NULL COMMENT '题目标识',
    `input`           TEXT NOT NULL COMMENT '测试用例输入',
    `expected_output` TEXT NOT NULL COMMENT '测试用例期望输出',
    `score`           TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '该用例分值(总分100)',
    `create_time`     TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`     TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time`     TIMESTAMP NULL COMMENT '删除时间',
    CONSTRAINT `pk_test_case_id` PRIMARY KEY (`id`),
    CONSTRAINT `uc_test_case_key` UNIQUE (`key`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin
  COMMENT = '测试用例表';

CREATE TABLE `submission`
(
    `id`               BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `key`              VARCHAR(32) NOT NULL COMMENT '提交记录标识',
    `user_id`          BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    `problem_key`      VARCHAR(32) NOT NULL COMMENT '题目标识',
    `code`             TEXT NOT NULL COMMENT '用户提交的代码',
    `language`         VARCHAR(50) NOT NULL COMMENT '编程语言',
    `status`           ENUM('pending', 'running', 'accepted', 'wrong_answer', 'time_limit_exceeded', 'memory_limit_exceeded', 'compile_error', 'runtime_error') NOT NULL DEFAULT 'pending' COMMENT '判题状态',
    `error_mag`        TEXT NULL COMMENT '错误信息',
    `score`            TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '得分(0-100)',
    `pass_case_count`  TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '通过用例数',
    `total_case_count` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '总用例数',
    `time_used`        INT NOT NULL COMMENT '代码最大运行耗时(单位: 毫秒)',
    `memory_used`      INT NOT NULL COMMENT '代码最大内存占用(单位: KB)',
    `submit_time`      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '提交时间',
    `create_time`      TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`      TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time`      TIMESTAMP NULL COMMENT '删除时间',
    CONSTRAINT `pk_submission_id` PRIMARY KEY (`id`),
    CONSTRAINT `uc_submission_key` UNIQUE (`key`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin
  COMMENT = '提交记录表';

CREATE TABLE `submission_case`
(
    `id`             BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `key`            VARCHAR(32)  NOT NULL COMMENT '用例执行结果标识',
    `submission_key` VARCHAR(32)  NOT NULL COMMENT '提交记录标识',
    `test_case_key`  VARCHAR(32)  NOT NULL COMMENT '测试用例标识',
    `user_id`        BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    `problem_key`    VARCHAR(32)  NOT NULL COMMENT '题目标识',
    `passed`         TINYINT(1) NOT NULL COMMENT '是否通过: 1-通过, 0-未通过',
    `actual_output`  TEXT NULL COMMENT '代码实际输出(NULL=执行失败)',
    `error_msg`      TEXT NULL COMMENT '执行错误信息',
    `time_used`      INT UNSIGNED NULL COMMENT '该用例执行耗时(单位：毫秒)',
    `memory_used`    INT UNSIGNED NULL COMMENT '该用例执行内存占用(单位: KB)',
    `create_time`    TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`    TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time`    TIMESTAMP NULL COMMENT '删除时间',
    CONSTRAINT `pk_submission_case_id` PRIMARY KEY (`id`),
    CONSTRAINT `uc_submission_case_key` UNIQUE (`key`),
    CONSTRAINT `uc_submission_case_submission_test_case` UNIQUE (`submission_key`, `test_case_key`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin
  COMMENT = '用例执行结果表';