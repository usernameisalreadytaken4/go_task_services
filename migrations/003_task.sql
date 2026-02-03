CREATE TABLE tasks (
			id BIGSERIAL PRIMARY KEY,
			user_id integer REFERENCES users (id), 
			value varchar(255) NOT NULL UNIQUE,
            started timestamp,
            finished timestamp,
			created timestamp NOT NULL DEFAULT now(),
			updated timestamp
		);

CREATE TRIGGER tasks_set_updated_at
BEFORE UPDATE ON tasks
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();