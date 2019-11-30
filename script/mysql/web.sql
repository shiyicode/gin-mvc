/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 50720
 Source Host           : localhost:3306
 Source Schema         : web

 Target Server Type    : MySQL
 Target Server Version : 50720
 File Encoding         : 65001

 Date: 24/11/2019 05:00:03
*/

DROP DATABASE IF EXISTS `web`;
CREATE DATABASE `web`;
USE `web`;

-- ----------------------------
-- Table structure for group
-- ----------------------------
DROP TABLE IF EXISTS `group`;
CREATE TABLE `group` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '团队ID',
  `name` varchar(80) DEFAULT NULL COMMENT '团队名称',
  `description` varchar(255) DEFAULT NULL COMMENT '团队描述',
  `users` varchar(255) DEFAULT NULL COMMENT '团队成员',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10001 DEFAULT CHARSET=utf8 COMMENT '团队表';
-- ----------------------------
-- Table structure for user
-- ----------------------------

CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL auto_increment COMMENT '用户ID',
  `email` varchar(256) binary default '' NOT NULL COMMENT '用户邮箱',
  `username` varchar(100) binary default '' NOT NULL COMMENT '用户名',
  `password` varchar(100) default '' NOT NULL COMMENT '用户密码',
  PRIMARY KEY(`id`),
  KEY `email` (`email`),
  KEY `username` (`username`),
  KEY `password` (`password`)
)ENGINE=InnoDB AUTO_INCREMENT=10001 default CHARSET=utf8 COMMENT '用户表';