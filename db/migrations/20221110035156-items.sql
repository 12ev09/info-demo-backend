-- +migrate Up
ALTER TABLE items ADD COLUMN author VARCHAR(255);
 
-- +migrate Down
ALTER TABLE items DROP author;
