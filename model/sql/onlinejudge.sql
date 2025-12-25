CREATE TABLE `user`
(
    `id`            BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `user_name`     VARCHAR(50)  NOT NULL COMMENT '登录用户名',
    `password_hash` VARCHAR(255) NOT NULL COMMENT '密码哈希值',
    `status`        BOOLEAN NOT NULL DEFAULT 1 COMMENT '用户状态: true-启用, false-禁用',
    `create_time`   TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`   TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time`   TIMESTAMP NULL COMMENT '删除时间',
    CONSTRAINT `pk_id` PRIMARY KEY (`id`),
    CONSTRAINT `uc_user_name` UNIQUE (`user_name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin
  COMMENT = '用户表';

CREATE TABLE `problem`
(
    `id`            BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `key`           VARCHAR(32)  NOT NULL COMMENT '题目标识',
    `title`         VARCHAR(50)  NOT NULL COMMENT '题目标题',
    `description`   TEXT NOT NULL COMMENT '题目描述',
    `input_desc`    TEXT NOT NULL COMMENT '输入格式说明',
    `output_desc`   TEXT NOT NULL COMMENT '输出格式说明',
    `sample_cases`  JSON NOT NULL COMMENT '样例输入输出',
    `difficulty`    VARCHAR(8) NOT NULL COMMENT '题目难度: easy/medium/hard',
    `time_limit`    INT NOT NULL COMMENT '程序运行时间限制(单位: 毫秒)',
    `memory_limit`  INT NOT NULL COMMENT '程序内存限制(单位: MB)',
    `tags`          VARCHAR(8) NOT NULL COMMENT '题目标签',
    `is_published`  BOOL NOT NULL DEFAULT FALSE COMMENT '是否发布',
    `create_time`   TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`   TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time`   TIMESTAMP NULL COMMENT '删除时间',
    CONSTRAINT `pk_id` PRIMARY KEY (`id`),
    CONSTRAINT `uc_key` UNIQUE (`key`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin
  COMMENT = '题目表';

CREATE TABLE `case`
(
    `id`          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `key`         VARCHAR(8)  NOT NULL COMMENT '测试用例标识',
    `problem_key` VARCHAR(50)  NOT NULL COMMENT '题目标识',
    `input`       TEXT NOT NULL COMMENT '测试用例输入',
    `output`      TEXT NOT NULL COMMENT '测试用例期望输出',
    `score`       INT NOT NULL DEFAULT 0 COMMENT '该用例分值(总分100)',
    `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time` TIMESTAMP NULL COMMENT '删除时间',
    CONSTRAINT `pk_id` PRIMARY KEY (`id`),
    CONSTRAINT `uc_key` UNIQUE (`key`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin
  COMMENT = '测试用例表';

CREATE TABLE `submission`
(
    `id`           BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `key`          VARCHAR(8)  NOT NULL COMMENT '提交记录标识',
    `user_id`      BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    `problem_key`  VARCHAR(50)  NOT NULL COMMENT '题目标识',
    `code`         TEXT NOT NULL COMMENT '用户提交的代码',
    `language`     VARCHAR(50) NOT NULL COMMENT '编程语言',
    `status`       VARCHAR(30) NOT NULL DEFAULT 'pending' COMMENT '判题状态',
    `error_msg`    TEXT NULL COMMENT '错误信息',
    `case_results` JSON NULL COMMENT '各测试用例执行结果',
    `score`        TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '得分',
    `time_used`    INT UNSIGNED NULL COMMENT '代码运行耗时(单位: 毫秒)',
    `memory_used`  INT UNSIGNED NULL COMMENT '代码内存占用(单位: KB)',
    `submit_time`  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '提交时间',
    `create_time`  TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time`  TIMESTAMP NULL COMMENT '删除时间',
    CONSTRAINT `pk_id` PRIMARY KEY (`id`),
    CONSTRAINT `uc_key_user_id_problem_key` UNIQUE (`key`, `user_id`, `problem_key`)

) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin
  COMMENT = '提交记录表';
