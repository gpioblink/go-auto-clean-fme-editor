package fme

import fmeDomain "github.com/gpioblink/go-auto-clean-fme-editor/pkg/fme/converterDomain/fme"

type MemoryRepository struct {
	fme fmeDomain.Fme
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{fmeDomain.Fme{}}
}

func (m *MemoryRepository) Save(fmeToSave *fmeDomain.Fme) error {
	m.fme = *fmeToSave
	return nil
}

func (m *MemoryRepository) Get() (*fmeDomain.Fme, error) {
	return &m.fme, nil
}
