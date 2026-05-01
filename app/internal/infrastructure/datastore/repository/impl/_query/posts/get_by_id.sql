SELECT id, user_id, title, content, created_at, updated_at, tags, version
FROM posts
WHERE id = $1
