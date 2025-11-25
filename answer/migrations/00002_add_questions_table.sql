-- +goose Up
INSERT INTO questions (text) VALUES
    ('Вопрос 1?'),
    ('Вопрос 2?'),
    ('Вопрос 3?');

-- +goose Down
DELETE FROM questions
WHERE text IN ('Вопрос 1?', 'Вопрос 2?', 'Вопрос 3?');