INSERT INTO posts (content, title, user_id, tags)
VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
