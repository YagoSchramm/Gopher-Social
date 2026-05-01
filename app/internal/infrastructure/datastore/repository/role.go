package repository

import (
	"context"

	"github.com/YagoSchramm/gopher-social/internal/domain"
)

type RoleRepository interface {
	GetByName(context.Context, string) (*domain.Role, error)
}
