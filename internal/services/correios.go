package services

import (
	"errors"
	"github.com/Gierdiaz/internal/domain/entities"
)

type CorreiosServiceImpl struct {
	BaseURL string
}

func NewCorreiosService(baseURL string) *CorreiosServiceImpl {
	return &CorreiosServiceImpl{BaseURL: baseURL}
}

func (s *CorreiosServiceImpl) CalcularFrete(frete *entities.Frete) (*entities.Frete, error) {
	if frete.Peso <= 0 {
		return nil, errors.New("peso inválido")
	}

	frete.Preco = frete.Peso * 5.0
	frete.PrazoEntrega = 7 // Exemplo: 7 dias úteis

	return frete, nil
}

func (s *CorreiosServiceImpl) ConsultarPrazo(cepOrigem, cepDestino string) (int, error) {
	return 5, nil // Exemplo: 5 dias úteis
}
