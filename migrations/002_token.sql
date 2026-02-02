CREATE TABLE tokens (
			id BIGSERIAL PRIMARY KEY,
			user_id integer REFERENCES users (id), 
			value varchar(255) NOT NULL UNIQUE,
			created timestamp NOT NULL DEFAULT now(),
			updated timestamp
		);

CREATE TRIGGER tokens_set_updated_at
BEFORE UPDATE ON tokens
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();