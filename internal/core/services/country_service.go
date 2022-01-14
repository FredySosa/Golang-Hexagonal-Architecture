package services

import (
	"context"
	"strconv"

	"github.com/FredySosa/hexArch/internal/core/domain"
	"github.com/FredySosa/hexArch/internal/ports"
)

type CountryService struct {
	countryRepository ports.CountriesRepository
}

func NewCountryService(cr ports.CountriesRepository) CountryService {
	return CountryService{
		countryRepository: cr,
	}
}

func (s CountryService) GetCountries(ctx context.Context, data domain.QueryData) ([]domain.Country, string, error) {
	countries, newOffset, err := s.countryRepository.GetCountries(ctx, data)
	if err != nil {
		return nil, "", err
	}

	return countries, strconv.Itoa(newOffset), nil
}
