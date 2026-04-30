SELECT u.id, u.username, u.email, u.created_at, u.is_active
FROM users u
JOIN user_invitations ui ON u.id = ui.user_id
WHERE ui.token = $1 AND ui.expiry > $2
