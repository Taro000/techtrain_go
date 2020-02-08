
-- +migrate Up
CREATE TABLE IF NOT EXISTS `techtrain_game`.`user_character` (
  `user_id` CHAR(36) NOT NULL,
  `character_id` CHAR(36) NOT NULL,
  `user_character_id` VARCHAR(36) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `modified_at` DATETIME NOT NULL,
  PRIMARY KEY (`user_character_id`),
  INDEX `fk_user_has_character_character1_idx` (`character_id` ASC) VISIBLE,
  INDEX `fk_user_has_character_user_idx` (`user_id` ASC) VISIBLE,
  UNIQUE INDEX `id_UNIQUE` (`user_character_id` ASC) VISIBLE,
  CONSTRAINT `fk_user_has_character_user`
    FOREIGN KEY (`user_id`)
    REFERENCES `techtrain_game`.`user` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_user_has_character_character1`
    FOREIGN KEY (`character_id`)
    REFERENCES `techtrain_game`.`character` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
-- +migrate Down
