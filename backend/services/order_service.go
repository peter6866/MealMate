package services

import (
	"context"
	"time"

	custom_errors "github.com/peter6866/foodie/custom-errors"
	"github.com/peter6866/foodie/models"
	"github.com/peter6866/foodie/repositories"
	"github.com/peter6866/foodie/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderService struct {
	userRepo  *repositories.UserRepository
	orderRepo *repositories.OrderRepository
}

func NewOrderService(userRepo *repositories.UserRepository, orderRepo *repositories.OrderRepository) *OrderService {
	return &OrderService{
		userRepo:  userRepo,
		orderRepo: orderRepo,
	}
}

// Create a new order
func (s *OrderService) CreateOrder(ctx context.Context, userID string, order *models.Order) error {
	userObjectId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	order.CreatedBy = userObjectId
	order.Status = models.OrderStatusStarted
	order.OrderDate = time.Now()
	user, err := s.userRepo.FindByID(ctx, userObjectId)
	if err != nil {
		return err
	}

	if !user.IsChef {
		chefUser, err := s.userRepo.FindByEmail(ctx, user.PartnerEmail)
		if err != nil {
			return err
		}
		order.SendTo = chefUser.ID

		utilsItem := make([]utils.Item, len(order.Items))
		for i, item := range order.Items {
			utilsItem[i] = utils.Item{
				ID:       item.ID.Hex(),
				Name:     item.Name,
				ImageUrl: item.ImageURL,
			}
		}

		orderDetails := utils.OrderDetails{
			CustomerName: user.Name,
			Items:        utilsItem,
			OrderTime:    time.Now().Format("2006-01-02 15:04"),
		}

		err = s.orderRepo.Create(ctx, order)
		if err != nil {
			return err
		}

		err = utils.SendConfirmationEmail(chefUser.Email, orderDetails)
		if err != nil {
			return err
		}

	} else {
		order.SendTo = user.ID

		err = s.orderRepo.Create(ctx, order)
		if err != nil {
			return err
		}
	}

	return nil
}

// get all orders
func (s *OrderService) GetAllOrders(ctx context.Context, userID string) ([]*models.Order, error) {
	userObjectId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.FindByID(ctx, userObjectId)
	if err != nil {
		return nil, err
	}

	if user.IsChef {
		return s.orderRepo.GetAllForChef(ctx, userObjectId)
	} else {
		return s.orderRepo.GetAllForUser(ctx, userObjectId)
	}
}

func (s *OrderService) CompleteOrder(ctx context.Context, userID string, orderID string) error {
	userObjectId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_errors.ErrInvalidObjectID
	}

	orderObjectId, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return custom_errors.ErrInvalidObjectID
	}

	user, err := s.userRepo.FindByID(ctx, userObjectId)
	if err != nil {
		return err
	}

	if !user.IsChef {
		return custom_errors.ErrUnauthorized
	}

	order, err := s.orderRepo.FindByID(ctx, orderObjectId)
	if err != nil {
		return custom_errors.ErrOrderNotFound
	}

	if order.SendTo != userObjectId {
		return custom_errors.ErrUnauthorized
	}

	if order.Status != models.OrderStatusStarted {
		return custom_errors.ErrOrderCompleted
	}

	return s.orderRepo.CompleteOrder(ctx, orderObjectId, userObjectId)
}
