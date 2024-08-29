package repository

import "github.com/halalala222/GoBoilder/internal/template/repository"

var _ Template = &MySQLRepositoryTemplate{}

// MySQLRepositoryTemplate is the template for MySQL repository
type MySQLRepositoryTemplate struct {
}

// Build returns the template for MySQL repository
func (t *MySQLRepositoryTemplate) Build() []byte {
	return repository.MysqlTemplate
}
