-- +goose Up

CREATE TABLE messages(
    id UUID PRIMARY KEY,
    body TEXT NOT NULL,
    posted_at TIMESTAMPTZ NOT NULL,
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE SET NULL,
    project_id UUID NOT NULL REFERENCES projects(project_id) ON DELETE CASCADE
);

-- +goose Down

DROP TABLE messages;