package storage

import (
	"GymMembership-api/internal/models"
	"context"
)

func (s *Storage) GetPlanById(ctx context.Context, id int64) (*models.MembershipPlan, error) {
	query := `SELECT id, name, price, duration_days, visits_limit,is_active , created_at FROM membership_plans WHERE id = $1`
	plan := &models.MembershipPlan{}
	err := s.pool.QueryRow(ctx, query, id).Scan(
		&plan.Id,
		&plan.Name,
		&plan.Price,
		&plan.DurationDays,
		&plan.VisitsLimit,
		&plan.IsActive,
		&plan.CreatedAt,
	)
	return plan, err
}

func (s *Storage) GetAllPlans(ctx context.Context) ([]*models.MembershipPlan, error) {
	query := `SELECT id, name, price, duration_days, visits_limit,is_active , created_at FROM membership_plans`

	rows, err := s.pool.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	plans := []*models.MembershipPlan{}
	for rows.Next() {
		plan := &models.MembershipPlan{}
		err := rows.Scan(
			&plan.Id,
			&plan.Name,
			&plan.Price,
			&plan.DurationDays,
			&plan.VisitsLimit,
			&plan.IsActive,
			&plan.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		plans = append(plans, plan)
	}

	return plans, err

}

func (s *Storage) CreatePlan(ctx context.Context, plan *models.MembershipPlan) error {
	query := `INSERT INTO membership_plans (name, price, duration_days, visits_limit,is_active) VALUES ($1 , $2 , $3 , $4 , $5)`
	_, err := s.pool.Exec(ctx, query,
		plan.Name,
		plan.Price,
		plan.DurationDays,
		plan.VisitsLimit,
		plan.IsActive,
	)
	return err
}
