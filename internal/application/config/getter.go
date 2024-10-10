package config

import "github.com/mariobgr/pack-shipment-exercise/internal/domain"

type PacksGetter struct {
	packsAvailable *envVars
}

// NewPacksGetter returns a dynamic packs getter
func NewPacksGetter() PacksGetter {
	return PacksGetter{packsAvailable: &EnvConfig}
}

// GetPacks implements domain.PacksGetter
func (getter PacksGetter) GetPacks() domain.Packs {
	return getter.packsAvailable.AvailablePackSizes
}
