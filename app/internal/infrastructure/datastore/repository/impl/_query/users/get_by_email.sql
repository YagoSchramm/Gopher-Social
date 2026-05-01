SELECT id, username, email, password, created_at FROM users
WHERE email = $1 AND is_active = true
