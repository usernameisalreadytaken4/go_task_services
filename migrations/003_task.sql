CREATE TABLE tasks (
			id BIGSERIAL PRIMARY KEY,
			user_id integer REFERENCES users (id), 
			payload JSONB NOT NULL DEFAULT '{}',
			result JSONB NOT NULL DEFAULT '{}',
            started timestamp,
            finished timestamp,
			created timestamp NOT NULL DEFAULT now(),
			updated timestamp
		);

CREATE TRIGGER tasks_set_updated_at
BEFORE UPDATE ON tasks
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();