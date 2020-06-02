/*
 Navicat Premium Data Transfer

 Source Server         : mysql57
 Source Server Type    : MySQL
 Source Server Version : 50721
 Source Host           : localhost:3307
 Source Schema         : jt

 Target Server Type    : MySQL
 Target Server Version : 50721
 File Encoding         : 65001

 Date: 04/05/2020 16:55:36
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for course_main
-- ----------------------------
DROP TABLE IF EXISTS `course_main`;
CREATE TABLE `course_main`  (
  `course_id` int(11) NOT NULL AUTO_INCREMENT,
  `course_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `course_disp_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `course_price` decimal(10, 2) NULL DEFAULT 0.00,
  `course_price2` decimal(10, 0) NULL DEFAULT 0,
  `addtime` timestamp(0) NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`course_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of course_main
-- ----------------------------
INSERT INTO `course_main` VALUES (1, '4个月成为架构师', '4个月成为架构师,不辛苦、无需流汗', 9.90, 0, '2020-05-04 00:36:45');
INSERT INTO `course_main` VALUES (2, '60天精通python全栈', '60天精通python全栈', 10.00, 0, '2020-05-04 00:36:48');

SET FOREIGN_KEY_CHECKS = 1;
