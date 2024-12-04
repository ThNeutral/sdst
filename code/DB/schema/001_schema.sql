-- +goose Up
CREATE TABLE users(
    user_id UUID PRIMARY KEY,
    token UUID NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(30) UNIQUE NOT NULL,
    password VARCHAR(72) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    last_login TIMESTAMPTZ NOT NULL
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
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE project_metadata (
    project_id UUID PRIMARY KEY,
    metadata JSONB,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    CONSTRAINT fk_metadaata FOREIGN KEY (project_id)
        REFERENCES projects(project_id)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);

CREATE TABLE project_users(
    project_id UUID NOT NULL,
    user_id UUID NOT NULL,
    role VARCHAR(12) NOT NULL,
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

DROP TABLE system_logs;
DROP TABLE project_metadata;
DROP TABLE project_users;
DROP TABLE project_file;
DROP TABLE projects;
DROP TABLE users;
