package http

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/mariobgr/pack-shipment-exercise/internal/application/service"
	"github.com/mariobgr/pack-shipment-exercise/internal/infra/logger"
	"github.com/mariobgr/pack-shipment-exercise/internal/infra/marshaling/json"

	"github.com/go-chi/chi/v5"
)

type PackShipmentHandler struct {
	calculatorService *service.PackCalculatorService
	logger            *logger.Logger
}

// NewPacksShipmentHandler creates a new SequenceHandler
func NewPacksShipmentHandler(calculatorService *service.PackCalculatorService, logger *logger.Logger) *PackShipmentHandler {
	return &PackShipmentHandler{
		calculatorService: calculatorService,
		logger:            logger,
	}
}

func (handler *PackShipmentHandler) Routes() chi.Router {
	// Only allow Authorized users
	r := chi.NewRouter().With(PermissionsMiddleware)

	r.Get("/", handler.list)
	r.Post("/calculate", handler.calculate)

	return r
}

// list returns all pack sizes for reference
func (handler *PackShipmentHandler) list(w http.ResponseWriter, r *http.Request) {
	packSizes := handler.calculatorService.GetSizes()
	if len(packSizes) == 0 {
		handler.logger.Warn("no pack sizes found")
		notFound(w, r)
	}

	successResponse(w, r, http.StatusOK, json.NewFromDomainPacks(packSizes))
}

// calculate returns the packs needed for the requested items
func (handler *PackShipmentHandler) calculate(w http.ResponseWriter, r *http.Request) {
	payload := &json.ItemsRequestedPayload{}
	if err := render.DecodeJSON(r.Body, &payload); err != nil {
		badRequest(w, r, handler.logger, err)
		return
	}

	packsToShip := handler.calculatorService.Calculate(payload.Items)

	successResponse(w, r, http.StatusOK, json.NewFromCalculatedShipment(packsToShip))
}
