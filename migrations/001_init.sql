CREATE TABLE users (
			id BIGSERIAL PRIMARY KEY,
			email varchar(255) NOT NULL UNIQUE,
			password text NOT NULL,
			created timestamp NOT NULL DEFAULT now(),
			updated timestamp
		);

CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN 
    NEW.updated = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER users_set_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();