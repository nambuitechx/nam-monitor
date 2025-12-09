package hosts

import (
	"github.com/nambuitechx/go-monitor/backend/configs"
)

type HostRepository struct {
	DB *configs.PostgresConnection
}

func NewHostRepository(db *configs.PostgresConnection) *HostRepository {
	return &HostRepository{
		DB: db,
	}
}


func (r *HostRepository) SelectHosts() ([]Host, error) {
	var hosts []Host
	err := r.DB.DB.Select(&hosts, "SELECT * FROM hosts")

	if err != nil {
		return  nil, err
	}

	return hosts, nil
}

func (r *HostRepository) InsertHost(host *Host) (*Host, error) {
	_, err := r.DB.DB.NamedExec("INSERT INTO hosts(id, url, created_at) VALUES (:id, :url, :created_at)", host)

	if err != nil {
		return nil, err
	}

	return host, nil
}
