package services

import "ums/internal/interfaces"

// type HealthCheckService interface {
// 	HealthCheckRepository() (string, error)
// }

type Healthcheck struct {
	HealthCheckRepository interfaces.IHealthCheckRepository
}

func (s *Healthcheck) HealthCheckServices() (string, error) {
	return "serve healthy", nil
}
