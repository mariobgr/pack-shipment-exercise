package json

import "fmt"

// PackListResponse is the response for API method /list
type PackListResponse struct {
	Packs []string `json:"packs"`
}

// NewFromDomainPacks converts the result to json response PackListResponse
func NewFromDomainPacks(packs []int) PackListResponse {
	var out PackListResponse

	for _, pack := range packs {
		out.Packs = append(out.Packs, fmt.Sprintf("%d items", pack))
	}

	return out
}

// CalculatedShipmentResponse is the response for API method /calculate
type CalculatedShipmentResponse struct {
	NumPacks []string `json:"packages_required"`
}

// NewFromCalculatedShipment converts the result to json response CalculatedShipmentResponse
func NewFromCalculatedShipment(ship map[int]int) CalculatedShipmentResponse {
	var out CalculatedShipmentResponse

	for size, count := range ship {
		out.NumPacks = append(out.NumPacks, fmt.Sprintf("%d x %d", count, size))
	}

	return out
}
