SELECT 
	p.id, p.user_id, p.title, p.content, p.created_at, p.version, p.tags,
	u.username,
	COUNT(c.id) AS comments_count
FROM posts p
LEFT JOIN comments c ON c.post_id = p.id
LEFT JOIN users u ON p.user_id = u.id
JOIN followers f ON f.follower_id = p.user_id OR p.user_id = $1
WHERE 
	f.user_id = $1 AND
	(p.title ILIKE '%' || $4 || '%' OR p.content ILIKE '%' || $4 || '%') AND
	(p.tags @> $5 OR $5 = '{}')
GROUP BY p.id, u.username
ORDER BY p.created_at {{SORT}}
LIMIT $2 OFFSET $3
