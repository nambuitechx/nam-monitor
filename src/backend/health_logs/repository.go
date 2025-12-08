package healthlogs

import (
	"github.com/nambuitechx/go-monitor/backend/configs"
)

type HealthLogRepository struct {
	// DB *configs.ScyllaConnection
	DB *configs.PostgresConnection
}

// func NewHealthLogRepository(db *configs.ScyllaConnection) *HealthLogRepository {
func NewHealthLogRepository(db *configs.PostgresConnection) *HealthLogRepository {
	return &HealthLogRepository{
		DB: db,
	}
}

// func (r *HealthLogRepository) SelectHealthLogs() ([]HealthLog, error) {
// 	var healthlogs []HealthLog
// 	q := r.DB.Session.Query(HealthLogTable.Get())

// 	if err := q.GetRelease(&healthlogs); err != nil {
// 		return nil, err
// 	}

// 	return healthlogs, nil
// }

// func (r *HealthLogRepository) InsertHealthLog(healthLog *HealthLog) (*HealthLog, error) {
// 	q := r.DB.Session.Query(HealthLogTable.Insert()).BindStruct(healthLog)

// 	if err := q.ExecRelease(); err != nil {
// 		return nil, err
// 	}

// 	return healthLog, nil
// }


func (r *HealthLogRepository) SelectHealthLogs() ([]HealthLog, error) {
	var healthlogs []HealthLog
	err := r.DB.DB.Select(&healthlogs, "SELECT * FROM health_logs")

	if err != nil {
		return  nil, err
	}

	return healthlogs, nil
}

func (r *HealthLogRepository) InsertHealthLog(healthLog *HealthLog) (*HealthLog, error) {
	_, err := r.DB.DB.NamedExec("INSERT INTO health_logs(id, status, created_at) VALUES (:id, :status, :created_at)", healthLog)

	if err != nil {
		return nil, err
	}

	return healthLog, nil
}
