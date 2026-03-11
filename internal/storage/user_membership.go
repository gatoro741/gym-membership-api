package storage

import (
	"GymMembership-api/internal/models"
	"context"
)

func (s *Storage) CreateUserMembershipPlan(ctx context.Context, userMembership models.UserMembership) error {
	query := `INSERT INTO user_memberships (user_id, plan_id, start_date, end_date, visits_left, is_active) VALUES ($1 ,$2, $3 ,$4, $5 ,$6)`

	_, err := s.pool.Exec(ctx, query,
		userMembership.UserId,
		userMembership.PlanId,
		userMembership.StartDate,
		userMembership.EndDate,
		userMembership.VisitsLeft,
		userMembership.IsActive,
	)
	return err

}

func (s *Storage) GetMembershipByUserId(ctx context.Context, userId int64) (*models.UserMembership, error) {
	query := `SELECT id , user_id, plan_id, start_date, end_date, visits_left, is_active , created_at FROM user_memberships WHERE user_id = $1`
	userMembership := &models.UserMembership{}
	err := s.pool.QueryRow(ctx, query, userId).Scan(
		&userMembership.Id,
		&userMembership.UserId,
		&userMembership.PlanId,
		&userMembership.StartDate,
		&userMembership.EndDate,
		&userMembership.VisitsLeft,
		&userMembership.IsActive,
		&userMembership.CreatedAt,
	)
	return userMembership, err
}

func (s *Storage) DeactivateExpiredMemberships(ctx context.Context) error {
	query := `UPDATE user_memberships SET is_active = false WHERE end_date < NOW() AND is_active = true`
	_, err := s.pool.Exec(ctx, query)
	return err
}
