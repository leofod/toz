package handler

// service "toz/pkg/service"
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FormData struct {
	Url string `form:"url" binding:"required"`
}

func (h *HandlerUrl) GetFull(c *gin.Context) {
	if answer, err := h.services.Url.GetFull(c.Param("url")); err != nil {
		switch answer {
		case "1":
			c.JSON(400, "Invalid URL")
		default:
			c.JSON(404, "NO such URL.")
		}
	} else {
		c.JSON(200, answer)
	}
}

func (h *HandlerUrl) Create(c *gin.Context) {
	var formData FormData
	if err := c.Bind(&formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if answer, err := h.services.Url.Create(formData.Url); err != nil {
		c.JSON(400, "Invalid URL.")
	} else {
		c.JSON(200, answer)
	}
}
