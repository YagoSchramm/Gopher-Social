package infrastructure

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/YagoSchramm/gopher-social/internal/domain"
)

//go:embed _query/roles/get_by_name.sql
var roleGetByNameQuery string

func NewRoleRepository(db *sql.DB) RoleRepository {
	return &RoleStore{db: db}
}

type RoleStore struct {
	db *sql.DB
}

func (s *RoleStore) GetByName(ctx context.Context, slug string) (*domain.Role, error) {
	role := &domain.Role{}
	err := s.db.QueryRowContext(ctx, roleGetByNameQuery, slug).Scan(&role.ID, &role.Name, &role.Description, &role.Level)
	if err != nil {
		return nil, err
	}

	return role, nil
}
