-- +goose Up
CREATE TABLE questions (
    id BIGSERIAL PRIMARY KEY,
    text TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_questions_created_at ON questions(created_at);

-- +goose Down
DROP TABLE questions;

