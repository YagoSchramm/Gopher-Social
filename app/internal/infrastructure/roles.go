package infrastructure

import (
	"context"
	"database/sql"

	"github.com/YagoSchramm/gopher-social/internal/domain"
)

type RoleStore struct {
	db *sql.DB
}

func (s *RoleStore) GetByName(ctx context.Context, slug string) (*domain.Role, error) {
	query := `SELECT id, name, description, level FROM roles WHERE name = $1`

	role := &domain.Role{}
	err := s.db.QueryRowContext(ctx, query, slug).Scan(&role.ID, &role.Name, &role.Description, &role.Level)
	if err != nil {
		return nil, err
	}

	return role, nil
}
