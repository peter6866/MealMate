package services

import (
	"context"

	"github.com/peter6866/foodie/models"
	"github.com/peter6866/foodie/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(ctx context.Context, category *models.Category) error {
	return s.repo.Create(ctx, category)
}

func (s *CategoryService) GetAllCategories(ctx context.Context) ([]*models.Category, error) {
	return s.repo.GetAll(ctx)
}
