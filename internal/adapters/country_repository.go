package adapters

import (
	"context"

	"github.com/FredySosa/hexArch/internal/core/domain"
)

type CountryRepository struct {
	countries map[string]domain.Country
}

func NewCountryRepository(initialData ...map[string]domain.Country) CountryRepository {
	data := make(map[string]domain.Country, 0)

	if initialData != nil {
		for key, value := range initialData[0] {
			data[key] = value
		}
	}
	return CountryRepository{
		countries: data,
	}
}

func (c CountryRepository) GetCountries(ctx context.Context, data domain.QueryData) ([]domain.Country, int, error) {
	countries := make([]domain.Country, 0)

	for _, country := range c.countries {
		countries = append(countries, country)
	}

	return countries, 0, nil
}
