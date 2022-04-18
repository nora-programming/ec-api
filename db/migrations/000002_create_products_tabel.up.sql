CREATE TABLE IF NOT EXISTS products(
   `id`             bigint(20) unsigned AUTO_INCREMENT NOT NULL COMMENT 'ID',
   `title`          varchar(255) NOT NULL COMMENT 'タイトル',
   `description`    varchar(255) NOT NULL COMMENT 'ディスクリプション',
   `price`          bigint(20) NOT NULL COMMENT '値段',
   `creater_id`     bigint(20) unsigned NOT NULL COMMENT 'user_id',

   PRIMARY KEY (`id`),
   CONSTRAINT `products_fk_creater_id` FOREIGN KEY (`creater_id`) REFERENCES `users` (`id`)
);
