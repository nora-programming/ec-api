CREATE TABLE IF NOT EXISTS purchases(
   `id`           bigint(20) unsigned AUTO_INCREMENT NOT NULL COMMENT 'ID',
   `product_id`   bigint(20) unsigned NOT NULL COMMENT 'product_id',
   `buyer_id`     bigint(20) unsigned NOT NULL COMMENT 'user_id',
   `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
   `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,

   PRIMARY KEY (`id`),
   CONSTRAINT `purchases_fk_product_id` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
   CONSTRAINT `purchases_fk_buyer_id` FOREIGN KEY (`buyer_id`) REFERENCES `users` (`id`)
);
