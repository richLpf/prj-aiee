CREATE DATABASE IF NOT EXISTS aiee;

CREATE TABLE IF NOT EXISTS aiee.permission (
  `id` int NOT NULL AUTO_INCREMENT,
  `key` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `type` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `desc` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `attribute` json DEFAULT NULL,
  `created_at` int DEFAULT NULL,
  `updated_at` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS aiee.role (
  `id` int NOT NULL AUTO_INCREMENT,
  `key` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `desc` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` int DEFAULT NULL,
  `created_at` int DEFAULT NULL,
  `updated_at` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


CREATE TABLE IF NOT EXISTS aiee.role_permission (
  `id` int NOT NULL AUTO_INCREMENT,
  `role_id` int DEFAULT NULL,
  `permission_id` int NOT NULL,
  `created_at` int DEFAULT NULL,
  `updated_at` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS aiee.user (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `password` longtext COLLATE utf8mb4_general_ci,
  `name` longtext COLLATE utf8mb4_general_ci,
  `age` bigint DEFAULT NULL,
  `status` int DEFAULT NULL,
  `created_at` int DEFAULT NULL,
  `updated_at` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS aiee.user_role (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `role_id` int NOT NULL,
  `role_name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_at` int DEFAULT NULL,
  `updated_at` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;