
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO `permissions`(router) VALUES("/api/sendmailhandle");

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM `permissions` WHERE route = "/api/sendmailhandle";
