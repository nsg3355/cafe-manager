-- payhere.product_info definition

CREATE TABLE IF NOT EXISTS `product_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '상품 ID',
  `user_id` int(11) NOT NULL COMMENT '사용자 ID',
  `category` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '카테고리',
  `price` decimal(10,2) NOT NULL COMMENT '가격',
  `cost` decimal(10,2) NOT NULL COMMENT '원가',
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '이름',
  `initial` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '초성',
  `description` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '설명',
  `barcode` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '바코드',
  `expiration_date` date NOT NULL COMMENT '유통기한',
  `size` enum('small','large') COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '사이즈',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '생성일시',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '수정일시',
  PRIMARY KEY (`id`),
  UNIQUE KEY `barcode` (`barcode`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `product_info_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;