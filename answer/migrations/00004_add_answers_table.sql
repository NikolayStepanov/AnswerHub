-- +goose Up
INSERT INTO answers (question_id, user_id, text) VALUES
    (1, '11111111-1111-1111-1111-111111111111', 'User’s response 1 to question 1'),
    (1, '22222222-2222-2222-2222-222222222222', 'User’s response 2 to question 1');

INSERT INTO answers (question_id, user_id, text) VALUES
    (2, '33333333-3333-3333-3333-333333333333', 'User’s response 3 to question 2');

-- +goose Down
DELETE FROM answers
WHERE question_id IN (1, 2);