BEGIN;

CREATE TABLE IF NOT EXISTS users (
    id bigint PRIMARY KEY AUTO_INCREMENT,
    name varchar(255) NULL
);

COMMIT;