package healthlogs

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
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

func (s *HealthLogService) CheckHealth(url string) (*HealthLog, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("HEAD", url, nil)

	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	status := fmt.Sprintf("%d - %s", res.StatusCode, res.Status)

	healthLog := &HealthLog{
		ID: uuid.New().String(),
		HostID: "ID",
		Status: status,
		CreatedAt: time.Now(),
	}

	return s.Repo.InsertHealthLog(healthLog)
}
