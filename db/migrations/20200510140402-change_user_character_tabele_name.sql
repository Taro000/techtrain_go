
-- +migrate Up
RENAME TABLE `techtrain_game`.`user_character` TO `techtrain_game`.`user_characters`;
-- +migrate Down
