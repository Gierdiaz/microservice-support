package interfaces

import entities "github.com/Gierdiaz/internal/domain/entities"

type CorreiosInterface interface {
	CalcularFrete(frete *entities.Frete) (*entities.Frete, error)
	ConsultarPrazo(cepOrigem, cepDestino string) (int, error)
}