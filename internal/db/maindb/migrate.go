package maindb

import (
	"github.com/sumup-oss/go-pkgs/errors"
)

func (c *Client) Migrate() error {
	c.logger.Info(
		logMessageRunningMigrations,
		emojiField("💽"),
		dbNameField(c.name),
	)

	_, err := c.db.Exec(`
		CREATE EXTENSION IF NOT EXISTS pgcrypto;

		CREATE TABLE IF NOT EXISTS transactions_cache (
			hash TEXT PRIMARY KEY,
			data JSONB NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS users (
		    id SERIAL PRIMARY KEY,
		    username VARCHAR(255) UNIQUE NOT NULL,
		    hashed_password VARCHAR(255) NOT NULL,
    		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP		    
		);

		CREATE TABLE IF NOT EXISTS user_transactions (
		    id SERIAL PRIMARY KEY,
		    user_id INTEGER NOT NULL,
		    transaction_hash TEXT NOT NULL,
    		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP		    
		);

        CREATE UNIQUE INDEX IF NOT EXISTS user_transactions_unique_user_id_transaction_hash
            ON user_transactions (user_id, transaction_hash);

		INSERT INTO users (username, hashed_password)
		VALUES
		    ('alice', crypt('alice', gen_salt('bf', 8))),
		    ('bob',   crypt('bob',   gen_salt('bf', 8))),
		    ('carol', crypt('carol', gen_salt('bf', 8))),
		    ('dave',  crypt('dave',  gen_salt('bf', 8)))
		ON CONFLICT DO NOTHING;
	`)
	if err != nil {
		return errors.Wrap(err, "failed to create transactions table: %s", err.Error())
	}

	return nil
}
