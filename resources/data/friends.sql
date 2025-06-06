/*
 Navicat Premium Dump SQL

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80041 (8.0.41)
 Source Host           : localhost:3306
 Source Schema         : chatify

 Target Server Type    : MySQL
 Target Server Version : 80041 (8.0.41)
 File Encoding         : 65001

 Date: 17/05/2025 20:16:36
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for friends
-- ----------------------------
DROP TABLE IF EXISTS `friends`;
CREATE TABLE `friends`  (
  `id` char(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT (UUID()),
  `mine` char(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `friend` char(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  FOREIGN KEY (`mine`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (`friend`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of friends
-- ----------------------------
INSERT INTO `friends` VALUES ('e497f609-331d-11f0-8dfc-e454e83a8e0d', '3553028d-a33b-42a7-9cd7-7a9b0c880192', 'aff2c966-7870-463e-a245-01e669ba8722', 'Insert by database.');

SET FOREIGN_KEY_CHECKS = 1;
