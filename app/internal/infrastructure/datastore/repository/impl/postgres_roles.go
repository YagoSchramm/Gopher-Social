package impl

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"

	"github.com/YagoSchramm/gopher-social/internal/derr"
	"github.com/YagoSchramm/gopher-social/internal/domain"
	"github.com/YagoSchramm/gopher-social/internal/infrastructure/datastore/repository"
)

//go:embed _query/roles/get_by_name.sql
var roleGetByNameQuery string

func NewRoleRepository(db *sql.DB) repository.RoleRepository {
	return &RoleStore{db: db}
}

type RoleStore struct {
	db *sql.DB
}

func (s *RoleStore) GetByName(ctx context.Context, slug string) (*domain.Role, error) {
	role := &domain.Role{}
	err := s.db.QueryRowContext(ctx, roleGetByNameQuery, slug).Scan(&role.ID, &role.Name, &role.Description, &role.Level)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, derr.RoleNotFound
		}
		return nil, err
	}

	return role, nil
}
