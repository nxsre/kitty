/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50547
 Source Host           : localhost
 Source Database       : kitty

 Target Server Type    : MySQL
 Target Server Version : 50547
 File Encoding         : utf-8

 Date: 09/12/2016 22:20:22 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `job_info`
-- ----------------------------
DROP TABLE IF EXISTS `job_info`;
CREATE TABLE `job_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `job_name` varchar(255) NOT NULL COMMENT '任务名称',
  `job_group` varchar(255) NOT NULL COMMENT '任务组',
  `state` int(11) NOT NULL DEFAULT '0' COMMENT '状态  0正常  1 删除',
  `params` varchar(255) DEFAULT NULL COMMENT '参数',
  `cron` varchar(150) NOT NULL COMMENT 'cron表达式',
  `url` varchar(255) NOT NULL COMMENT '具体执行任务的url地址 多个用,分割',
  `phone` varchar(32) DEFAULT NULL COMMENT '负责人的手机号码',
  `active` int(11) DEFAULT '0' COMMENT '是否激活  0 不激活  1 激活',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` text COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `job_info_history`
-- ----------------------------
DROP TABLE IF EXISTS `job_info_history`;
CREATE TABLE `job_info_history` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `job_name` varchar(255) NOT NULL COMMENT '任务名称',
  `job_group` varchar(255) NOT NULL COMMENT '任务组',
  `state` int(11) NOT NULL DEFAULT '0' COMMENT '状态  0正常  1 删除',
  `params` varchar(255) DEFAULT NULL COMMENT '参数',
  `cron` varchar(150) NOT NULL COMMENT 'cron表达式',
  `url` varchar(255) NOT NULL COMMENT '具体执行任务的url地址 多个用,分割',
  `phone` varchar(32) DEFAULT NULL COMMENT '负责人的手机号码',
  `active` int(11) DEFAULT '0' COMMENT '是否激活  0 不激活  1 激活',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` text COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `job_snapshot`
-- ----------------------------
DROP TABLE IF EXISTS `job_snapshot`;
CREATE TABLE `job_snapshot` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `job_name` varchar(255) NOT NULL COMMENT '任务名称',
  `job_group` varchar(255) NOT NULL COMMENT '任务组',
  `params` varchar(255) DEFAULT NULL COMMENT '参数',
  `cron` varchar(150) NOT NULL COMMENT 'cron表达式',
  `url` varchar(255) NOT NULL COMMENT '具体执行任务的url地址 多个用,分割',
  `phone` varchar(32) DEFAULT NULL COMMENT '负责人的手机号码',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `detail` text COMMENT '执行详情',
  `ip` varchar(32) DEFAULT NULL COMMENT '执行的目标ip',
  `state` int(11) DEFAULT NULL COMMENT '执行状态',
  `result` varchar(255) DEFAULT NULL COMMENT '执行结果',
  `time_consume` bigint(20) DEFAULT '0' COMMENT '任务耗时',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `job_snapshot_history`
-- ----------------------------
DROP TABLE IF EXISTS `job_snapshot_history`;
CREATE TABLE `job_snapshot_history` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `job_name` varchar(255) NOT NULL COMMENT '任务名称',
  `job_group` varchar(255) NOT NULL COMMENT '任务组',
  `params` varchar(255) DEFAULT NULL COMMENT '参数',
  `cron` varchar(150) NOT NULL COMMENT 'cron表达式',
  `url` varchar(255) NOT NULL COMMENT '具体执行任务的url地址 多个用,分割',
  `phone` varchar(32) DEFAULT NULL COMMENT '负责人的手机号码',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `detail` text COMMENT '执行详情',
  `ip` varchar(32) DEFAULT NULL COMMENT '执行的目标ip',
  `state` int(11) DEFAULT NULL COMMENT '执行状态',
  `result` varchar(255) DEFAULT NULL COMMENT '执行结果',
  `time_consume` bigint(20) DEFAULT '0' COMMENT '任务耗时',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
