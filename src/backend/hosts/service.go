package hosts

import (
	"time"

	"github.com/google/uuid"
)

type HostService struct {
	Repo *HostRepository
}

func NewHealthLogService(repo *HostRepository) *HostService {
	return &HostService{
		Repo: repo,
	}
}

func (s *HostService) GetAll() ([]Host, error) {
	return s.Repo.SelectHosts()
}

func (s *HostService) Create(url string) (*Host, error) {
	host := &Host{
		ID: uuid.New().String(),
		Url: url,
		CreatedAt: time.Now(),
	}

	return s.Repo.InsertHost(host)
}

