package storage

import (
	"GymMembership-api/internal/models"
	"context"
)

func (s *Storage) GetClassById(ctx context.Context, id int64) (*models.Class, error) {
	query := `SELECT id, title, start_time, trainer_name, capacity , occupied FROM classes WHERE id = $1`
	class := &models.Class{}
	err := s.pool.QueryRow(ctx, query, id).Scan(
		&class.Id,
		&class.Title,
		&class.StartTime,
		&class.TrainerName,
		&class.Capacity,
		&class.Occupied,
	)
	return class, err
}

func (s *Storage) GetAllClasses(ctx context.Context) ([]*models.Class, error) {
	query := `SELECT id, title, start_time, trainer_name, capacity , occupied FROM classes`

	rows, err := s.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	classes := []*models.Class{}
	for rows.Next() {
		class := &models.Class{}
		err := rows.Scan(
			&class.Id,
			&class.Title,
			&class.StartTime,
			&class.TrainerName,
			&class.Capacity,
			&class.Occupied,
		)
		if err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	return classes, err

}

func (s *Storage) CreateClass(ctx context.Context, class *models.Class) error {
	query := `INSERT INTO classes (title, start_time, trainer_name, capacity) VALUES ($1 , $2 , $3 , $4 )`
	_, err := s.pool.Exec(ctx, query,
		class.Title,
		class.StartTime,
		class.TrainerName,
		class.Capacity,
	)
	return err
}

func (s *Storage) IncrementOccupied(ctx context.Context, classId int64) error {
	query := `UPDATE classes SET occupied = occupied + 1 WHERE id = $1`
	_, err := s.pool.Exec(ctx, query, classId)
	return err
}
