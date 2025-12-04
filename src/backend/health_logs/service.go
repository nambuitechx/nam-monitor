package healthlogs

import (
	"time"

	"github.com/gocql/gocql"
)

type HealthLogService struct {
	Repo *HealthLogRepository
}

func NewHealthLogService(repo *HealthLogRepository) *HealthLogService {
	return &HealthLogService{
		Repo: repo,
	}
}

func (s *HealthLogService) GetAll() ([]HealthLog, error) {
	return s.Repo.SelectHealthLogs()
}

func (s *HealthLogService) Create(status string) (*HealthLog, error) {
	healthLog := &HealthLog{
		ID: gocql.TimeUUID(),
		Status: status,
		CreatedAt: time.Now(),
	}

	return s.Repo.InsertHealthLog(healthLog)
}
