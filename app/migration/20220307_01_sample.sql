CREATE TABLE `user_address` (
`id` bigint unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'id',
`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
`user_id` bigint NOT NULL DEFAULT 0 COMMENT '用户id',
`is_default` boolean NOT NULL DEFAULT FALSE COMMENT '是否默认地址',
`province` varchar(16) NOT NULL DEFAULT '' COMMENT '省',
`receive_at` bigint NOT NULL DEFAULT 0 COMMENT '收货时间',
INDEX idx_user_address_user_default(`user_id`,`is_default`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户地址信息表';
