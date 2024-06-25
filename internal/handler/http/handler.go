package http_handler

import (
	"github.com/gofiber/fiber/v2"
	"international-trading-test/internal/config"
	"international-trading-test/internal/pkg/fiber_errors"
	"international-trading-test/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(cfg *config.Config, service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(app *fiber.App) {

	app.Get("/hasher_client/:value", h.getInstrument)

}

// @Summary Get some value
// @Tags Some Client
// @Accept       json
// @Produce      json
// @Param        value   path      string  true  "Price alert external ID"
// @Success      200  {object}   service.HashedValue
// @Failure      400  {object}  fiber_errors.Error
// @Failure      500  {object}  fiber_errors.Error
// @Router       /hasher_client/{value} [get]
func (h *Handler) getInstrument(ctx *fiber.Ctx) error {
	params := ctx.AllParams()
	value, ok := params["value"]
	if !ok {
		return fiber_errors.BadRequest("value is not presented")
	}
	instruments, err := h.service.GetInfoFromSomeClient(ctx.Context(), value)
	if err != nil {
		return fiber_errors.InternalServerError(err)
	}
	return ctx.JSON(instruments)
}
