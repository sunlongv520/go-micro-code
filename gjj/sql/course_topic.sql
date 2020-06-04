/*
Navicat MySQL Data Transfer

Source Server         : 192.168.1.100
Source Server Version : 50553
Source Host           : 192.168.1.101:3306
Source Database       : test

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2020-06-03 22:33:22
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for course_topic
-- ----------------------------
DROP TABLE IF EXISTS `course_topic`;
CREATE TABLE `course_topic` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `course_id` int(11) NOT NULL DEFAULT '0' COMMENT '课程ID',
  `course_did` int(11) NOT NULL DEFAULT '0' COMMENT '课时ID',
  `likes` int(11) DEFAULT NULL COMMENT '点赞数',
  `unlikes` int(11) DEFAULT NULL COMMENT '不认同数',
  `title` varchar(500) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '标题',
  `content` text COLLATE utf8mb4_bin COMMENT '评论内容',
  `user_id` int(11) DEFAULT NULL COMMENT '发表者ID',
  `addtime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '发布时间',
  `updatetime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最后修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of course_topic
-- ----------------------------
