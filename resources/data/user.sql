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

 Date: 26/03/2025 20:45:42
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` char(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT (uuid()),
  `account` integer(8)  COLLATE 'utf8mb4_general_ci' NOT NULL COMMENT '账号',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码',
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COLLATE '头像',
  `created_time` timestamp NULL DEFAULT NULL COMMENT '创建时间,单位：秒',
  `updated_time` timestamp NULL DEFAULT NULL COMMENT '更新时间,单位：秒',
  `last_time` timestamp NULL DEFAULT NULL COMMENT '最后登录时间,单位：秒',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('3553028d-a33b-42a7-9cd7-7a9b0c880192', 10000001, 'admin', '$2a$10$CJ7PBrGBJkKIbP11i3s4MuWXbeJnl7o1EJ5zwL3U.LFfOYbWtze3G', 'Admin', '114ad59a4a2b71ecd92b4f2493457223e3c3528d787d88608f616f80a7ec6b04.jpg', '2025-03-30 21:32:35', '2025-04-03 21:41:56', '2025-04-29 21:36:21');
INSERT INTO `user` VALUES ('4a98633b-9a7d-4de0-bbdd-27783cc5cab6', 12331566, 'user', '$2a$10$YxYuW6QarYtCrWpujEUx3u3n9spEsKZ37jqIVsK.CzHYHlnOiNPNy', NULL, NULL, '2025-04-29 21:36:54', '2025-04-29 21:36:54', NULL);
INSERT INTO `user` VALUES ('aff2c966-7870-463e-a245-01e669ba8722', 10001209, 'lanxiz', '$2a$10$aoAxyfFEW9kq6Egf1T2FR.BubM44/Q6eD4QEBdKih3ngut0Aok502', NULL, NULL, '2025-04-29 21:37:17', '2025-04-29 21:37:17', NULL);

SET FOREIGN_KEY_CHECKS = 1;
