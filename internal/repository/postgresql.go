package repository

import "github.com/halalala222/GoBoilder/internal/template/repository"

var _ Template = &PostgreSQLRepositoryTemplate{}

// PostgreSQLRepositoryTemplate is the template for PostgreSQL repository
type PostgreSQLRepositoryTemplate struct {
}

// Build returns the template for PostgreSQL repository
func (t *PostgreSQLRepositoryTemplate) Build() []byte {
	return repository.PostgresqlTemplate
}
