-- name: CreateImage :one
INSERT INTO image (
    journal_id, content
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetImageOfJournal :many
SELECT * 
FROM image
WHERE journal_id = $1
ORDER BY created_at DESC;

-- name: UpdateImageEmbeddingStatus :one
UPDATE image
SET is_embedded = $1
WHERE id = $2
RETURNING *;

-- name: DeleteImage :exec
DELETE FROM image
WHERE id = $1;

-- name: DeleteImageOfJournal :exec
DELETE FROM image
WHERE journal_id = $1;