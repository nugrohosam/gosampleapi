BEGIN;


ALTER TABLE users 
DROP COLUMN email;

ALTER TABLE users 
DROP COLUMN password;

COMMIT;