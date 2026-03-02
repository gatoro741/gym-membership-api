package service

import (
	"GymMembership-api/internal/models"
	"context"
	"log"
	"time"
)

func (s *Service) BuyMembership(ctx context.Context, planId int, userId int64) error {
	plan, err := s.storage.GetPlanById(ctx, planId)
	if err != nil {
		log.Printf("BuyMembership: failed to get plan %d: %v", planId, err)
		return err
	}

	endDate := time.Now().Add(time.Duration(plan.DurationDays) * 24 * time.Hour)
	userMembership := models.UserMembership{
		UserId:  userId,
		PlanId:  planId,
		EndDate: endDate,
	}

	err = s.storage.CreateUserMembershipPlan(ctx, userMembership)
	if err != nil {
		log.Printf("BuyMembership: failed to create usermembership: %v", err)
		return err
	}
	return nil
}

func (s *Service) GetMyMembership(ctx context.Context, userId int64) (*models.UserMembership, error) {
	membership, err := s.storage.GetMembershipByUserId(ctx, userId)
	if err != nil {
		log.Printf("GetMyMembership: failed to get membership")
		return nil, err
	}
	return membership, nil

}
