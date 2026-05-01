SELECT c.id, c.content, c.post_id, c.user_id, c.created_at, c.updated_at, u.id, u.username
FROM comments c
JOIN users u ON u.id = c.user_id
WHERE c.post_id = $1
ORDER BY c.created_at ASC
