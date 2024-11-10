-- +goose Up
CREATE TYPE user_roles AS ENUM ('Super Admin', 'Admin', 'User');

CREATE TABLE roles(
    role_id UUID PRIMARY KEY,
    role user_roles NOT NULL,
    role_description VARCHAR(256)
);

CREATE TABLE users(
    user_id UUID PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    password VARCHAR(72) NOT NULL,
    email VARCHAR(30) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    last_login TIMESTAMP,
    role_id UUID NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (role_id)
        REFERENCES roles (role_id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE system_logs(
    log_id UUID NOT NULL PRIMARY KEY,
    message TEXT NOT NULL,
    context JSONB,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    user_id UUID NOT NULL,
    CONSTRAINT fk_logs FOREIGN KEY (user_id)
        REFERENCES users (user_id)
        ON UPDATE CASCADE
);

CREATE TABLE projects (
    project_id UUID PRIMARY KEY,
    p_name VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    owner_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE project_metadata (
    project_id UUID PRIMARY KEY,
    metadata JSONB,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_metadaata FOREIGN KEY (project_id)
        REFERENCES projects(project_id)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);

CREATE TABLE project_users(
    project_id UUID NOT NULL,
    user_id UUID NOT NULL,
    PRIMARY KEY (project_id, user_id),
    CONSTRAINT fk_proj_user FOREIGN KEY (project_id)
        REFERENCES projects (project_id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    CONSTRAINT fk_users_proj FOREIGN KEY (user_id)
        REFERENCES users (user_id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE project_file(
    file_id UUID NOT NULL PRIMARY KEY,
    file_location VARCHAR(256),
    file_content TEXT,
    project_id UUID NOT NULL,
    FOREIGN KEY (project_id) REFERENCES projects (project_id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- +goose Down

DROP TABLE roles;
DROP TABLE users;
DROP TABLE system_logs;
DROP TABLE projects;
DROP TABLE project_metadata;
DROP TABLE project_users;
DROP TABLE project_file;
