/*
Navicat MySQL Data Transfer

Source Server         : 127.0.0.1
Source Server Version : 50553
Source Host           : 127.0.0.1:3306
Source Database       : video

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2019-09-01 11:18:23
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `video_id` int(11) DEFAULT NULL,
  `author_id` int(11) DEFAULT NULL,
  `content` varchar(255) DEFAULT NULL,
  `time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of comments
-- ----------------------------

-- ----------------------------
-- Table structure for sessions
-- ----------------------------
DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions` (
  `session_id` text NOT NULL,
  `ttl` tinyint(4) DEFAULT NULL,
  `login_name` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sessions
-- ----------------------------
INSERT INTO `sessions` VALUES ('f587a91f-fef2-4570-a7ce-9f774d166e0f', '127', 'bbb');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `login_name` varchar(255) DEFAULT NULL,
  `pwd` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES ('1', 'aaa', 'aaa');
INSERT INTO `users` VALUES ('2', 'bbb', 'aaa');

-- ----------------------------
-- Table structure for video_del_rec
-- ----------------------------
DROP TABLE IF EXISTS `video_del_rec`;
CREATE TABLE `video_del_rec` (
  `video_id` varchar(64) NOT NULL,
  PRIMARY KEY (`video_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of video_del_rec
-- ----------------------------

-- ----------------------------
-- Table structure for video_info
-- ----------------------------
DROP TABLE IF EXISTS `video_info`;
CREATE TABLE `video_info` (
  `id` varchar(11) NOT NULL,
  `author_id` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `display_ctime` text,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of video_info
-- ----------------------------
