# 在线判题系统

## 研发进度

### 用户管理模块

支持注册、登录的简单功能，设计数据库表结构如下：

``` sql
CREATE TABLE `user`
(
    `id`      BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `user_name` VARCHAR(50)  NOT NULL COMMENT '登录用户名',
    `password_hash` VARCHAR(255) NOT NULL COMMENT '密码哈希值',
    `status` BOOLEAN NOT NULL DEFAULT 1 COMMENT '用户状态: true-启用, false-禁用',
    `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time` TIMESTAMP NULL COMMENT '删除时间',
    CONSTRAINT `pk_id` PRIMARY KEY (`id`),
    CONSTRAINT `uc_user_name` UNIQUE (`user_name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin
  COMMENT = '用户表';
```




