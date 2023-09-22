DROP TABLE IF EXISTS projects CASCADE;
CREATE TABLE projects (
    id UUID PRIMARY KEY NOT NULL,
    title VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

DROP TABLE IF EXISTS secrets CASCADE;
CREATE TABLE secrets (
    id UUID PRIMARY KEY,
    key VARCHAR(255) NOT NULL,
    value TEXT NOT NULL,
    version INT DEFAULT 1,
    project_id UUID NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES projects(id)
);

DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL,
    active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

DROP TABLE IF EXISTS groups CASCADE;
CREATE TABLE groups (
    id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    project_id UUID NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES projects(id)
);

DROP TABLE IF EXISTS roles CASCADE;
CREATE TABLE roles (
    name VARCHAR(255) UNIQUE
);

-- Many to many relations
DROP TABLE IF EXISTS group_roles;
CREATE TABLE group_roles (
    group_id UUID NOT NULL,
    role VARCHAR(255) NOT NULL,
    PRIMARY KEY (group_id, role),
    FOREIGN KEY (group_id) REFERENCES groups(id),
    FOREIGN KEY (role) REFERENCES roles(name)
);

DROP TABLE IF EXISTS group_users;
CREATE TABLE group_users (
    group_id UUID NOT NULL,
    user_id UUID NOT NULL,
    PRIMARY KEY (group_id, user_id),
    FOREIGN KEY (group_id) REFERENCES groups(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

DROP TABLE IF EXISTS user_tokens;
CREATE TABLE user_tokens (
    user_id UUID NOT NULL,
    access_token TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

INSERT INTO roles (name) VALUES
     ('manage-users'),
     ('manage-groups'),
     ('read-secrets'),
     ('write-secrets'),
     ('delete-secrets');

CREATE OR REPLACE FUNCTION manage_timestamps()
    RETURNS TRIGGER AS $$
    BEGIN
        IF (TG_OP = 'INSERT') THEN
            NEW.created_at := current_timestamp;
            NEW.updated_at := current_timestamp;
        ELSIF (TG_OP = 'UPDATE') THEN
            NEW.updated_at := current_timestamp;
        END IF;
        RETURN NEW;
    END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION manage_versions()
    RETURNS TRIGGER AS $secrets$
BEGIN
    IF (TG_OP = 'INSERT') THEN
        NEW.version := 1;
    ELSIF (TG_OP = 'UPDATE') THEN
        NEW.version := OLD.version+1;
    END IF;
    RETURN NEW;
END;
$secrets$ LANGUAGE plpgsql;

CREATE TRIGGER secret_version_trigger
    BEFORE INSERT OR UPDATE ON secrets
    FOR EACH ROW
    EXECUTE PROCEDURE manage_versions();

CREATE TRIGGER user_time_trigger
    BEFORE INSERT OR UPDATE ON users
    FOR EACH ROW
    EXECUTE PROCEDURE manage_timestamps();

CREATE TRIGGER project_time_trigger
    BEFORE INSERT OR UPDATE ON projects
    FOR EACH ROW
EXECUTE PROCEDURE manage_timestamps();

CREATE TRIGGER secret_time_trigger
    BEFORE INSERT OR UPDATE ON secrets
    FOR EACH ROW
EXECUTE PROCEDURE manage_timestamps();

CREATE TRIGGER group_time_trigger
    BEFORE INSERT OR UPDATE ON groups
    FOR EACH ROW
EXECUTE PROCEDURE manage_timestamps();

