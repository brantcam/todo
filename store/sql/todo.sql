-- Queries --

-- name: get-all
SELECT * FROM list;

-- name: add-item
INSERT INTO list(todo, done) VALUES ($1, $2);