-- +migrate Up
CREATE TABLE IF NOT EXISTS items (
    id bigint AUTO_INCREMENT NOT NULL,
    title VARCHAR(255),
    isbn VARCHAR(255),
    publisher_name VARCHAR(255),
    sales_date VARCHAR(255),
    content_type int,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
 
-- +migrate Down
DROP TABLE IF EXISTS items;
