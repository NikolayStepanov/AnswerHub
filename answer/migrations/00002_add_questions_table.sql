-- +goose Up
INSERT INTO questions (text) VALUES
    ('Question 1?'),
    ('Question 2?'),
    ('Question 3?');

-- +goose Down
DELETE FROM questions
WHERE text IN ('Question 1?', 'Question 2?', 'Question 3?');