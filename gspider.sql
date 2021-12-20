-- MySQL dump 10.13  Distrib 8.0.25, for macos10.14 (x86_64)
--
-- Host: 42.192.203.133    Database: gspider
-- ------------------------------------------------------
-- Server version	5.7.36

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `cxytiandi`
--

DROP TABLE IF EXISTS `cxytiandi`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cxytiandi` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL COMMENT '标题名称',
  `cdate` varchar(30) NOT NULL COMMENT '发布日期',
  `visit` varchar(30) NOT NULL DEFAULT '' COMMENT '访问数',
  `author` varchar(20) NOT NULL DEFAULT '' COMMENT '文章作者',
  `content` text COMMENT '内容',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_title` (`title`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_updated_at` (`updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='抓取猿天地信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cxytiandi`
--

LOCK TABLES `cxytiandi` WRITE;
/*!40000 ALTER TABLE `cxytiandi` DISABLE KEYS */;
/*!40000 ALTER TABLE `cxytiandi` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `task`
--

DROP TABLE IF EXISTS `task`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `task` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `task_name` varchar(64) NOT NULL COMMENT '任务名称',
  `task_rule_name` varchar(64) NOT NULL COMMENT '任务规则名',
  `task_desc` varchar(512) NOT NULL DEFAULT '' COMMENT '任务名称',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '任务状态 0 未知状态 1 运行中 2 停止 3 完成 4 异常终止',
  `counts` int(11) NOT NULL DEFAULT '0' COMMENT '已执行次数',
  `cron_spec` varchar(64) NOT NULL DEFAULT '' COMMENT '任务执行表达式',
  `opt_user_agent` varchar(128) NOT NULL DEFAULT '' COMMENT '请求UA',
  `opt_max_depth` int(11) NOT NULL DEFAULT '0' COMMENT '最大爬取深度',
  `opt_allowed_domains` varchar(512) NOT NULL DEFAULT '' COMMENT '允许的域名 多个用,分割',
  `opt_url_filters` varchar(512) NOT NULL DEFAULT '' COMMENT 'url过滤规则',
  `opt_max_body_size` int(11) NOT NULL DEFAULT '0' COMMENT 'Body最大大小',
  `opt_request_timeout` int(11) NOT NULL DEFAULT '10' COMMENT '请求超时时间',
  `limit_enable` tinyint(4) NOT NULL DEFAULT '0',
  `limit_domain_regexp` varchar(128) NOT NULL DEFAULT '',
  `limit_domain_glob` varchar(128) NOT NULL DEFAULT '',
  `limit_delay` int(11) NOT NULL DEFAULT '0',
  `limit_random_delay` int(11) NOT NULL DEFAULT '0',
  `limit_parallelism` int(11) NOT NULL DEFAULT '0',
  `proxy_urls` varchar(2048) NOT NULL DEFAULT '' COMMENT '代理url',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_task_name` (`task_name`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='任务表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `task`
--

LOCK TABLES `task` WRITE;
/*!40000 ALTER TABLE `task` DISABLE KEYS */;
INSERT INTO `task` VALUES (1,'猿天地','程序员天地','猿天地技术文章',2,119,'*/5 * * * * *','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36',0,'','',0,10,0,'','',0,0,0,'','2021-11-26 02:01:21','2021-12-20 03:06:52'),(2,'网易新闻','网易科技新闻','科技新闻',2,21,'*/15 * * * * *','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36',0,'','',0,10,0,'','',0,0,0,'','2021-11-26 02:01:21','2021-12-17 08:46:03'),(4,'测试添加1','程序员天地','这是个测试的任务',2,6,'*/20 * * * * *','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36',1,'','',30000,10,1,'','*',3,2,1,'','2021-12-18 17:23:27','2021-12-20 03:08:29'),(8,'网易科技抓取2','网易科技新闻','这只是个测试的例子',2,1,'*/20 * * * * *','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36',1,'','',0,10,0,'','',0,0,1,'','2021-12-20 03:20:33','2021-12-20 03:21:20');
/*!40000 ALTER TABLE `task` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(32) NOT NULL COMMENT '用户名',
  `password` varchar(128) NOT NULL COMMENT '密码',
  `roles` varchar(128) NOT NULL DEFAULT '' COMMENT '角色',
  `introduction` varchar(128) NOT NULL DEFAULT '' COMMENT '描述',
  `avatar` varchar(256) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_uname` (`username`),
  KEY `idx_updated_at` (`updated_at`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'admin','$2a$10$qPIBmHQnPMF869iRp5/RpuyJEJ/LEuMz639gukegEVZMofv/5AuXm','admin','admin user','/admin/gspider.jpeg','2021-12-13 01:28:04','2021-12-20 08:11:06');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wangyi_tech_news`
--

DROP TABLE IF EXISTS `wangyi_tech_news`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `wangyi_tech_news` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL COMMENT '标题名称',
  `cdate` varchar(30) NOT NULL COMMENT '发布日期',
  `content` text COMMENT '内容',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_title` (`title`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_updated_at` (`updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='网友科技新闻';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wangyi_tech_news`
--

LOCK TABLES `wangyi_tech_news` WRITE;
/*!40000 ALTER TABLE `wangyi_tech_news` DISABLE KEYS */;
/*!40000 ALTER TABLE `wangyi_tech_news` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'gspider'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-12-20 16:59:39
