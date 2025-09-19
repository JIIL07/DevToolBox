package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/JIIL07/devtoolbox/internal/core"
)

type Handler struct {
	registry *core.GeneratorRegistry
}

func NewHandler(registry *core.GeneratorRegistry) *Handler {
	return &Handler{
		registry: registry,
	}
}

func (h *Handler) Generate(c *gin.Context) {
	var req GenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, GenerateResponse{
			Error: "Invalid request format: " + err.Error(),
		})
		return
	}

	generator, exists := h.registry.Get(req.Template)
	if !exists {
		c.JSON(http.StatusNotFound, GenerateResponse{
			Error: "Template not found: " + req.Template,
		})
		return
	}

	code, err := generator.Generate(req.Input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenerateResponse{
			Error: "Generation failed: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, GenerateResponse{
		Code: code,
	})
}

func (h *Handler) ListGenerators(c *gin.Context) {
	generators := h.registry.List()
	generatorInfos := make([]GeneratorInfo, len(generators))
	
	for i, gen := range generators {
		generatorInfos[i] = GeneratorInfo{
			Name:        gen.GetName(),
			Description: gen.GetDescription(),
		}
	}

	c.JSON(http.StatusOK, ListGeneratorsResponse{
		Generators: generatorInfos,
	})
}

func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"service": "devtoolbox-api",
	})
}
