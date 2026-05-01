INSERT INTO comments (content, post_id, user_id)
VALUES ($1, $2, $3)
RETURNING id, created_at, updated_at
