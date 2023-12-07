USE `webapi`;

TRUNCATE TABLE `posts`;

INSERT INTO `posts` (`content`, `created_at`) VALUES ('post1', NOW());
INSERT INTO `posts` (`content`, `created_at`) VALUES ('post2', NOW());
INSERT INTO `posts` (`content`, `created_at`) VALUES ('post3', NOW());
INSERT INTO `posts` (`content`, `created_at`) VALUES ('post4', NOW());
