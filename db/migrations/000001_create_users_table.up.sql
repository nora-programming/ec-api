CREATE TABLE IF NOT EXISTS users(
   `id`         bigint(20) unsigned AUTO_INCREMENT NOT NULL COMMENT 'ID',
   `name`       varchar(255) NOT NULL COMMENT '名前',
   `email`      varchar(255) NOT NULL COMMENT 'メールアドレス',
   `password`   varchar(255) NOT NULL COMMENT 'パスワード',
   `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
   `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,

   PRIMARY KEY (`id`),
   UNIQUE KEY `idx_email` (`email`)
);
