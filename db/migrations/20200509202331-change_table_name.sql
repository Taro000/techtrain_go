
-- +migrate Up
RENAME TABLE `techtrain_game`.`user` TO `techtrain_game`.`users`;

RENAME TABLE `techtrain_game`.`character` TO `techtrain_game`.`characters`;
-- +migrate Down
