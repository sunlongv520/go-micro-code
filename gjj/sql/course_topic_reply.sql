/*
Navicat MySQL Data Transfer

Source Server         : 192.168.1.100
Source Server Version : 50553
Source Host           : 192.168.1.101:3306
Source Database       : test

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2020-06-03 22:33:30
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for course_topic_reply
-- ----------------------------
DROP TABLE IF EXISTS `course_topic_reply`;
CREATE TABLE `course_topic_reply` (
  `id` int(11) NOT NULL,
  `topic_id` int(11) NOT NULL COMMENT '短评ID',
  `content` text COLLATE utf8mb4_bin COMMENT '回复内容',
  `user_id` int(11) DEFAULT NULL COMMENT '回复用户ID',
  `likes` int(11) DEFAULT NULL COMMENT '认同',
  `unlikes` int(11) DEFAULT NULL COMMENT '不认同',
  `addtime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '入库时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of course_topic_reply
-- ----------------------------
