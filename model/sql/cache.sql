-- 产品表
CREATE TABLE `product`
(
    `key`         VARCHAR(32) NOT NULL,
    `secret`      VARCHAR(32) NOT NULL,
    `name`        VARCHAR(32) NOT NULL,
    `category`    VARCHAR(64) NOT NULL DEFAULT '',
    `node_type`   INT         NOT NULL DEFAULT 0 COMMENT '0 直连设备，1 网关设备，2 网关子设备',
    `data_proto`  INT         NOT NULL DEFAULT 0 COMMENT '0 标准协议，1 自定义协议',
    `comm_mode`   INT         NOT NULL DEFAULT 0 COMMENT '0 Wi-Fi，1 Wi-Fi-蓝牙，2 蓝牙，3 Wi-Fi-ZigBee，4 ZigBee',
    `thing_model` LONGTEXT    NOT NULL,
    `create_time` TIMESTAMP   NULL     DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP   NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_time` TIMESTAMP   NULL,
    CONSTRAINT `pk_key` PRIMARY KEY (`key`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;