package entities

type Frete struct {
	CEPOrigem      string  `json:"cep_origem"`
	CEPDestino     string  `json:"cep_destino"`
	Peso           float64 `json:"peso"`
	ValorDeclarado float64 `json:"valor_declarado"`
	PrazoEntrega   int     `json:"prazo_entrega"` 
	Preco          float64 `json:"preco"`        
}
