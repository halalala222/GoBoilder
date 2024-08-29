package repository

import "github.com/halalala222/GoBoilder/internal/template/repository"

var _ Template = &MongoDBRepositoryTemplate{}

// MongoDBRepositoryTemplate is the template for MongoDB repository
type MongoDBRepositoryTemplate struct {
}

// Build returns the template for MongoDB repository
func (t *MongoDBRepositoryTemplate) Build() []byte {
	return repository.MongodbTemplate
}
