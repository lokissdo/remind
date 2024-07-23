-- name: CreateAudio :one
INSERT INTO audio (
    journal_id, content
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetAudioOfJournal :many
SELECT * 
FROM audio
WHERE journal_id = $1
ORDER BY created_at DESC;

-- name: UpdateAudioEmbeddingStatus :one
UPDATE audio
SET is_embedded = $1
WHERE id = $2
RETURNING *;