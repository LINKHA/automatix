/*
 Navicat Premium Data Transfer

 Source Server         : automatix
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : localhost:33069
 Source Schema         : amx_servermanager

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 05/11/2023 23:23:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for server
-- ----------------------------
DROP TABLE IF EXISTS `server`;
CREATE TABLE `server`  (
  `id` bigint(0) NOT NULL,
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `update_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
  `delete_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `del_state` tinyint(0) NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '服务器名称',
  `server_type` tinyint(0) NOT NULL COMMENT '服务类型(0:运营 1:测试)',
  `switch` tinyint(0) NOT NULL COMMENT '服务器是否开启(0:关闭 1:开启)',
  `start_time` int(0) NOT NULL COMMENT '开服时间',
  `area` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '服务器地区',
  `tags` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标签列表(使用逗号分隔)',
  `max_online` int(0) NOT NULL COMMENT '最大在线人数',
  `max_queue` int(0) NOT NULL COMMENT '最大排队人数',
  `max_sign` int(0) NOT NULL COMMENT '最大注册人数',
  `template_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '自定义参数(json格式)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
