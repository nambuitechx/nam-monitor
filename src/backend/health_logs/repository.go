package healthlogs

import (
	"github.com/nambuitechx/go-monitor/backend/configs"
)

type HealthLogRepository struct {
	DB *configs.DBConnection
}

func NewHealthLogRepository(db *configs.DBConnection) *HealthLogRepository {
	return &HealthLogRepository{
		DB: db,
	}
}

func (r *HealthLogRepository) SelectHealthLogs() ([]HealthLog, error) {
	var healthlogs []HealthLog
	q := r.DB.Session.Query(HealthLogTable.Get())

	if err := q.GetRelease(&healthlogs); err != nil {
		return nil, err
	}

	return healthlogs, nil
}

func (r *HealthLogRepository) InsertHealthLog(healthLog *HealthLog) (*HealthLog, error) {
	q := r.DB.Session.Query(HealthLogTable.Insert()).BindStruct(healthLog)

	if err := q.ExecRelease(); err != nil {
		return nil, err
	}

	return healthLog, nil
}
