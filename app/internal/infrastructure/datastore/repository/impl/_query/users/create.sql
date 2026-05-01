INSERT INTO users (username, password, email, role_id) VALUES 
($1, $2, $3, (SELECT id FROM roles WHERE name = $4))
RETURNING id, created_at
