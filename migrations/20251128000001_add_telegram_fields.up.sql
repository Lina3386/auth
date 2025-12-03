-- +goose Up
ALTER TABLE users ADD COLUMN IF NOT EXISTS telegram_id BIGINT UNIQUE;
ALTER TABLE users ADD COLUMN IF NOT EXISTS telegram_username VARCHAR(255);
ALTER TABLE users ADD COLUMN IF NOT EXISTS auth_token TEXT;

CREATE INDEX IF NOT EXISTS idx_users_telegram_id ON users(telegram_id);
CREATE INDEX IF NOT EXISTS idx_users_auth_token ON users(auth_token);

-- +goose Down
DROP INDEX IF EXISTS idx_users_auth_token;
DROP INDEX IF EXISTS idx_users_telegram_id;
ALTER TABLE users DROP COLUMN IF EXISTS auth_token;
ALTER TABLE users DROP COLUMN IF EXISTS telegram_username;
ALTER TABLE users DROP COLUMN IF EXISTS telegram_id;


