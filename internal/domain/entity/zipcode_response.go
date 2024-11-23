package entity

type ZipcodeResponse struct {
	Localidade string `json:"localidade"`
	Erro       string `json:"erro"`
}

func NewZipcodeResponse(localidade string, erro string) *ZipcodeResponse {
	return &ZipcodeResponse{Localidade: localidade, Erro: erro}
}
