
-- +migrate Up
CREATE TABLE IF NOT EXISTS `techtrain_game`.`character` (
    `id` CHAR(36) NOT NULL,
    `name` VARCHAR(45) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `modified_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE)
ENGINE = InnoDB;
-- +migrate Down
