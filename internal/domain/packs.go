package domain

// Packs is the representation of available pack sizes
type Packs []int

// PacksGetter is an interface for the method that will get Packs sizes available
type PacksGetter interface {
	GetPacks() Packs
}
