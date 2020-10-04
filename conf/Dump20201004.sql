-- MySQL dump 10.13  Distrib 5.7.31, for Linux (x86_64)
--
-- Host: localhost    Database: ipfs_fileinfo
-- ------------------------------------------------------
-- Server version	5.7.31-0ubuntu0.16.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `file_info`
--

DROP TABLE IF EXISTS `file_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `file_info` (
  `fhash` varchar(100) COLLATE utf8_danish_ci NOT NULL,
  `title` varchar(100) COLLATE utf8_danish_ci DEFAULT NULL,
  `owners` varchar(10000) COLLATE utf8_danish_ci DEFAULT NULL,
  `uploader` varchar(100) COLLATE utf8_danish_ci DEFAULT NULL,
  `size` int(11) DEFAULT NULL,
  `authority_code` int(11) DEFAULT NULL,
  `note` varchar(1000) COLLATE utf8_danish_ci DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp(6) NULL DEFAULT NULL,
  PRIMARY KEY (`fhash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_danish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `file_info`
--

LOCK TABLES `file_info` WRITE;
/*!40000 ALTER TABLE `file_info` DISABLE KEYS */;
INSERT INTO `file_info` VALUES ('Qmasdi','标题','{\"QmfEgNcEJZdxZPeYPZSBxTzXXTDVSYKf8ZjREyBgYSceaP\":{\"dhash\":\"QmfEgNcEJZdxZPeYPZSBxTzXXTDVSYKf8ZjREyBgYSceaP\",\"status\":3,\"ip\":\"127.0.0.1\",\"port\":8435,\"capacity\":400,\"remain\":500,\"lastpingpongtime\":\"0001-01-01T00:00:00Z\"}}','MyDhash',100,7,'注释','2020-10-04 05:11:47',NULL,NULL),('QmfEgNcEJZdxZPeYPZSBxTzXXTDVSYKf8ZjREyBgYSceaP','','{\"Qma6wFrmDQ48u7GrZvvzrj5RkMAFBCNepNx62fVe5134GL\":{\"dhash\":\"Qma6wFrmDQ48u7GrZvvzrj5RkMAFBCNepNx62fVe5134GL\",\"status\":3,\"ip\":\"127.0.0.1\",\"port\":8435,\"capacity\":100,\"remain\":100,\"lastpingpongtime\":\"0001-01-01T00:00:00Z\"}}','',62,7,'','2020-10-04 05:12:58',NULL,NULL);
/*!40000 ALTER TABLE `file_info` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-10-04 13:29:17
