SELECT users.id, username, email, password, created_at, roles.*
FROM users
JOIN roles ON (users.role_id = roles.id)
WHERE users.id = $1 AND is_active = true
