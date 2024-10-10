package service

import (
	"slices"
	"sort"

	"github.com/mariobgr/pack-shipment-exercise/internal/domain"
)

type PackCalculatorService struct {
	packsGetter domain.PacksGetter
}

// NewCalculatorService initializes new service for calculating packs
func NewCalculatorService(packsGetter domain.PacksGetter) *PackCalculatorService {
	return &PackCalculatorService{
		packsGetter: packsGetter,
	}
}

// Calculate returns the number of packs required for the requested number of items
func (service *PackCalculatorService) Calculate(items int) map[int]int {
	packsToShip := make(map[int]int)

	packSizes := service.packsGetter.GetPacks()

	// sort the sizes in descending order for optimal traversing
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	for _, packSize := range packSizes {
		// calculate the number of packs of the current size we need
		if items >= packSize {
			numPacks := items / packSize
			packsToShip[packSize] = numPacks
			// reduce the order size by the number of items we've packed
			items -= numPacks * packSize
		}
	}

	// for any remaining items, use the smallest available pack
	if items > 0 {
		smallestPack := packSizes[len(packSizes)-1]
		if _, ok := packsToShip[smallestPack]; ok {
			// check if we can combine two of smallest into one larger - depends on config
			if slices.Contains(packSizes, smallestPack*2) {
				packsToShip[smallestPack*2]++
				// unset the small pack as it was combined with a larger one
				delete(packsToShip, smallestPack)
			}
		} else {
			packsToShip[smallestPack]++
		}
	}

	return packsToShip
}

// GetSizes returns the available packs sizes
func (service *PackCalculatorService) GetSizes() domain.Packs {
	return service.packsGetter.GetPacks()
}
