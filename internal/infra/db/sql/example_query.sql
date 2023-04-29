-- name: GetMessage :one
SELECT * FROM examples WHERE id = ? ;

-- name: ListMessages :many
SELECT * FROM examples;

-- name: InsertMessage :execresult
INSERT INTO examples (message) VALUES( ?);