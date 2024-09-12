-- starbucks.access_control definition

CREATE TABLE IF NOT EXISTS `access_control` (
  `user_id` int(11) NOT NULL COMMENT '사용자 ID',
  `access_token` varchar(255) NOT NULL COMMENT 'token(10분)',
  `expires_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '만료일시',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '생성일시',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `access_token` (`access_token`),
  CONSTRAINT `access_control_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;