BEGIN;

ALTER TABLE users 
ADD COLUMN username varchar(255) NOT NULL;

COMMIT;