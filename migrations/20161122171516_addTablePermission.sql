
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `permissions` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `router` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8;

--
-- Dumping data for table `permissions`
--

INSERT INTO `permissions` (`id`, `router`, `description`, `created_at`, `updated_at`) VALUES
(1, '/user/firstlogin/', '/user/firstlogin/', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(2, '/user/setting', '/setting', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(3, '/user/signout', '/signout', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(4, '/user/forgetpassword', '/forgetpassword', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(5, '/user/newpassword', '/newpassword', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(6, '/user/edit', '/edit', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(7, '/user/add', '/add', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(8, '/user/delete', 'delete user\r\n', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(9, '/object/edit', 'edit object\r\n', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(10, '/object/review', 'review object\r\n', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(11, '/version/manage', 'manage version', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(12, '/company/news', 'manage company news', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(13, '/right/addadmin', '/right/addadmin', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(14, '/right/addgod', '/right/addgod', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(15, '/api/getteam', '/api/getteam', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(16, '/api/getteambyparent', '/api/getteambyparent', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(17, '/api/getuserteam', '/api/getuserteam', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(18, '/api/getdatabycompany', '/api/getdatabycompany', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(19, '/api/getrequestdata', '/api/getrequestdata', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(20, '/api/getuserbyid', '/api/getuserbyid', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(21, '/api/getpopup', '/api/getpopup', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(22, '/api/getuserbyname', '/api/getuserbyname', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(23, '/api/getallevaluatee', '/api/getallevaluatee', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(24, '/api/import', '/api/import', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(25, '/api/confirmuser', '/api/confirmuser', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(26, '/api/export', '/api/export', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(27, '/api/download', '/api/download', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(28, '/api/getcompanybasicinfo', '/api/getcompanybasicinfo', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(29, '/api/getversioninfo', '/api/getversioninfo', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(30, '/api/updatecompanyteam', '/api/updatecompanyteam', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(31, '/news/update', '/news/update', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(32, '/news/delete', '/news/delete', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(33, '/review/getallotherreview', '/review/getallotherreview', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(34, '/error', '/error', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(35, '/dashboard', '/dashboard', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(36, '/registmembertop', '/registmembertop', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(37, '/matchingtop', '/matchingtop', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(38, '/matching', '/matching', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(39, '/getmatching', '/getmatching', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(40, '/company/basicinfo', '/companybasic', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(41, '/search', '/search', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(42, '/evaluatemanage', 'evaluate manage', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(43, '/contact', 'contact page', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(44, '/help', 'help page ', '2016-08-09 11:44:22', '2016-08-09 11:44:22'),
(46, '/api/getuserbyteam', ' get user by team', '2016-08-09 11:44:22', '2016-08-09 11:44:22');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `permissions`;
