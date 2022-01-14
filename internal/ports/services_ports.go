package ports

import (
	"context"

	"github.com/FredySosa/hexArch/internal/core/domain"
)

type EventService interface {
	GetCountries(ctx context.Context, limit, offset string) ([]domain.Country, string, error)
}

type CountriesService interface {
	GetCountries(ctx context.Context, data domain.QueryData) ([]domain.Country, string, error)
}

type CountriesRepository interface {
	GetCountries(ctx context.Context, data domain.QueryData) ([]domain.Country, int, error)
}
