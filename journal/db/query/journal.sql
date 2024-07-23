-- name: CreateJournal :one
INSERT INTO journal (
    username, title, content, status
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetJournal :one
SELECT * FROM journal
WHERE id = $1 LIMIT 1;

-- name: GetJournalFromUserInTime :many
SELECT * FROM journal
WHERE username = $1 AND updated_at >= $2 AND updated_at <= $3;

-- name: UpdateJournal :one
UPDATE journal
SET
    title = COALESCE(sqlc.narg(title), title),
    content = COALESCE(sqlc.arg(content), content),
    status = COALESCE(sqlc.narg(status), status),
    is_embedded = COALESCE(sqlc.arg(is_embedded), is_embedded),
    updated_at = now()
WHERE id = sqlc.arg(id)
RETURNING *;







