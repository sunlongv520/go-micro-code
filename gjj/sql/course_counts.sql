/*
Navicat MySQL Data Transfer

Source Server         : 192.168.1.100
Source Server Version : 50553
Source Host           : 192.168.1.101:3306
Source Database       : test

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2020-06-03 22:33:06
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for course_counts
-- ----------------------------
DROP TABLE IF EXISTS `course_counts`;
CREATE TABLE `course_counts` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `count_id` int(11) NOT NULL DEFAULT '0',
  `course_id` int(11) unsigned NOT NULL DEFAULT '0',
  `count_key` varchar(50) NOT NULL DEFAULT '',
  `count_value` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of course_counts
-- ----------------------------
INSERT INTO `course_counts` VALUES ('1', '1', '2', 'click', '3');
INSERT INTO `course_counts` VALUES ('2', '2', '2', 'fav', '2');
