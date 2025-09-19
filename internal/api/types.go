package api

type GenerateRequest struct {
	Template string `json:"template" binding:"required"`
	Input    string `json:"input" binding:"required"`
}

type GenerateResponse struct {
	Code string `json:"code"`
	Error string `json:"error,omitempty"`
}

type GeneratorInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ListGeneratorsResponse struct {
	Generators []GeneratorInfo `json:"generators"`
}
