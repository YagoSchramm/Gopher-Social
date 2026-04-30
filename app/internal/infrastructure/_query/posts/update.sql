UPDATE posts
SET title = $1, content = $2, version = version + 1
WHERE id = $3 AND version = $4
RETURNING version
