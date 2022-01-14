package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/FredySosa/hexArch/internal/core/domain"
	"github.com/FredySosa/hexArch/internal/ports"
)

type EventService struct {
	countriesSvc ports.CountriesService
}

func NewEventService(cs ports.CountriesService) EventService {
	return EventService{
		countriesSvc: cs,
	}
}

func (e EventService) GetCountries(ctx context.Context, limit, offset string) ([]domain.Country, string, error) {
	queryData := domain.QueryData{
		Offset: offset,
		Limit:  limit,
	}

	countries, newOffset, err := e.countriesSvc.GetCountries(ctx, queryData)
	if err != nil {
		if errors.Is(err, domain.ErrCountriesNotFound) {
			return []domain.Country{}, "", nil
		}
		return nil, "", fmt.Errorf("something went wrong: %w", err)
	}

	return countries, newOffset, nil
}
