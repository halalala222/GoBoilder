package repository

import "github.com/halalala222/GoBoilder/internal/template/repository"

var _ Template = &GormRepositoryTemplate{}

// GormRepositoryTemplate is the template for Gorm repository
type GormRepositoryTemplate struct {
}

// Build returns the template for Gorm repository
func (t *GormRepositoryTemplate) Build() []byte {
	return repository.GormTemplate
}
