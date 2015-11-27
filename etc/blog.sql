-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.5.24-log - MySQL Community Server (GPL)
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  8.3.0.4803
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- 导出 blog 的数据库结构
CREATE DATABASE IF NOT EXISTS `blog` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `blog`;


-- 导出  表 blog.article 结构
CREATE TABLE IF NOT EXISTS `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '文章标题',
  `uri` varchar(255) NOT NULL DEFAULT '' COMMENT 'URL中的文章标题',
  `keywords` varchar(2550) DEFAULT '' COMMENT '关键词',
  `content` LONGTEXT COMMENT '正文',
  `author` varchar(255) NOT NULL DEFAULT '' COMMENT '作者',
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  `count` int(11) NOT NULL DEFAULT '0' COMMENT '阅读次数',
  `status` int(4) NOT NULL DEFAULT '0' COMMENT '状态: 0草稿，1已发布',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uri` (`uri`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章';

-- 数据导出被取消选择。


-- 导出  表 blog.config 结构
CREATE TABLE IF NOT EXISTS `config` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `key` varchar(128) NOT NULL COMMENT 'key',
  `value` text NOT NULL COMMENT 'value',
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `key` (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='配置';

-- 数据导出被取消选择。


-- 导出  表 blog.file 结构
CREATE TABLE IF NOT EXISTS `file` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `filename` varchar(96) NOT NULL COMMENT '文件名',
  `path` varchar(128) NOT NULL COMMENT '路径',
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `store` enum('local','oss') NOT NULL COMMENT '存储类型',
  `mime` varchar(100) DEFAULT NULL COMMENT '文件类型',
  PRIMARY KEY (`id`),
  UNIQUE KEY `path` (`path`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文件';

-- 数据导出被取消选择。


-- 导出  表 blog.project 结构
CREATE TABLE IF NOT EXISTS `project` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL COMMENT '项目名称',
  `icon_url` varchar(256) NOT NULL COMMENT '图标地址',
  `author` varchar(50) NOT NULL COMMENT '作者',
  `description` text NOT NULL COMMENT '描述',
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='项目';

-- 数据导出被取消选择。


-- 导出  表 blog.users 结构
CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(255) NOT NULL DEFAULT '' COMMENT '盐',
  `email` varchar(255) DEFAULT '' COMMENT 'email',
  `varified` enum('Y','N') NOT NULL DEFAULT 'N' COMMENT '邮箱已验证',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户';

-- 数据导出被取消选择。


-- 导出  表 blog.varify 结构
CREATE TABLE IF NOT EXISTS `varify` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL DEFAULT '0' COMMENT '用户名',
  `code` varchar(128) NOT NULL DEFAULT '0' COMMENT '验证码',
  `overdue` datetime NOT NULL COMMENT '过期时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='保密验证码';

-- 数据导出被取消选择。
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */
