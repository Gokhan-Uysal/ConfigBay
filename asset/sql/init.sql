DROP TABLE IF EXISTS projects CASCADE;
CREATE TABLE projects (
    id UUID DEFAULT gen_random_uuid(),
    title VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
);

DROP TABLE IF EXISTS secrets CASCADE;
CREATE TABLE secrets (
    id UUID DEFAULT gen_random_uuid(),
    key VARCHAR(255) NOT NULL,
    value TEXT NOT NULL,
    version INT DEFAULT 1,
    project_id UUID NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (project_id, key),
    FOREIGN KEY (project_id) REFERENCES projects(id)
);

DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE users (
    id UUID DEFAULT gen_random_uuid(),
    username VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL,
    active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
);

DROP TABLE IF EXISTS roles CASCADE;
CREATE TABLE roles (
    name varchar(100),
    PRIMARY KEY (name),
    UNIQUE (name)
);

DROP TABLE IF EXISTS permissions CASCADE;
CREATE TABLE permissions (
    name VARCHAR(255),
    PRIMARY KEY (name)
);

DROP TABLE IF EXISTS groups CASCADE;
CREATE TABLE groups (
    id UUID DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    project_id UUID NOT NULL,
    role VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (project_id) REFERENCES projects(id),
    FOREIGN KEY (role) REFERENCES roles(name)
);

-- Many to many relations
DROP TABLE IF EXISTS role_permissions CASCADE;
CREATE TABLE role_permissions (
    role VARCHAR(255) NOT NULL,
    permission VARCHAR(255) NOT NULL,
    PRIMARY KEY (role, permission),
    FOREIGN KEY (role) REFERENCES roles(name),
    FOREIGN KEY (permission) REFERENCES permissions(name)
);

DROP TABLE IF EXISTS group_users CASCADE;
CREATE TABLE group_users (
    group_id UUID NOT NULL,
    user_id UUID NOT NULL,
    PRIMARY KEY (group_id, user_id),
    FOREIGN KEY (group_id) REFERENCES groups(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

DROP TABLE IF EXISTS user_tokens CASCADE;
CREATE TABLE user_tokens (
    user_id UUID NOT NULL,
    access_token TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Init permissions and roles
BEGIN;

INSERT INTO roles(name) VALUES
    ('Admin'),
    ('Team Leader'),
    ('Product Owner'),
    ('Developer')
ON CONFLICT DO NOTHING;

INSERT INTO permissions(name) VALUES
    ('read-project'),
    ('manage-users'),
    ('manage-groups'),
    ('read-secrets'),
    ('write-secrets'),
    ('delete-secrets')
ON CONFLICT DO NOTHING;

INSERT INTO role_permissions(role, permission) VALUES
    ('Admin', 'read-project'),
    ('Admin', 'manage-users'),
    ('Admin', 'manage-groups'),
    ('Admin', 'read-secrets'),
    ('Admin', 'write-secrets'),
    ('Admin', 'delete-secrets')
ON CONFLICT DO NOTHING;

INSERT INTO role_permissions(role, permission) VALUES
    ('Team Leader', 'manage-groups'),
    ('Team Leader', 'manage-users')
ON CONFLICT DO NOTHING;

INSERT INTO role_permissions(role, permission) VALUES
    ('Product Owner', 'read-secrets'),
    ('Product Owner', 'write-secrets'),
    ('Product Owner', 'delete-secrets')
ON CONFLICT DO NOTHING;

INSERT INTO role_permissions(role, permission) VALUES
    ('Developer', 'read-project'),
    ('Developer', 'read-secrets')
ON CONFLICT DO NOTHING;

COMMIT;


BEGIN;

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

COMMIT;


BEGIN;

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

COMMIT;