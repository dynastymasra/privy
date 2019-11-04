package category

import (
	"context"

	"github.com/dynastymasra/privy/config"
	"github.com/dynastymasra/privy/domain"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Repository domain.CategoryRepository
}

func NewService(repo domain.CategoryRepository) Service {
	return Service{Repository: repo}
}

func (s Service) Create(ctx context.Context, category domain.Category) (*domain.Category, error) {
	log := logrus.WithFields(logrus.Fields{
		config.RequestID: ctx.Value(config.HeaderRequestID),
		"category":       category,
	})

	res, err := s.Repository.Create(ctx, category)
	if err != nil {
		log.WithError(err).Errorln("Failed create new category")

		return nil, err
	}

	return res, nil
}

func (s Service) FindByID(ctx context.Context, id int) (*domain.Category, error) {
	log := logrus.WithFields(logrus.Fields{
		config.RequestID: ctx.Value(config.HeaderRequestID),
		"id":             id,
	})

	category, err := s.Repository.FindByID(ctx, id)
	if err != nil {
		log.WithError(err).Errorln("Failed get category by id")
		return nil, err
	}

	return category, nil
}

func (s Service) Fetch(ctx context.Context, from, size int) ([]domain.Category, error) {
	log := logrus.WithFields(logrus.Fields{
		config.RequestID: ctx.Value(config.HeaderRequestID),
		"from":           from,
		"size":           size,
	})

	categories, err := s.Repository.Fetch(ctx, from, size)
	if err != nil {
		log.WithError(err).Errorln("Failed fetch category")
		return nil, err
	}

	return categories, nil
}

func (s Service) Update(ctx context.Context, id int, category domain.Category) (*domain.Category, error) {
	log := logrus.WithFields(logrus.Fields{
		config.RequestID: ctx.Value(config.HeaderRequestID),
		"after":          category,
		"id":             id,
	})

	prod, err := s.Repository.FindByID(ctx, id)
	if err != nil {
		log.WithError(err).Errorln("Failed get category by id")
		return nil, err
	}

	category.ID = id
	if err := s.Repository.Update(ctx, category); err != nil {
		log.WithField("before", prod).WithError(err).Errorln("Failed update category")
		return nil, err
	}

	return &category, nil
}

func (s Service) Delete(ctx context.Context, id int) error {
	log := logrus.WithFields(logrus.Fields{
		config.RequestID: ctx.Value(config.HeaderRequestID),
		"id":             id,
	})

	category, err := s.Repository.FindByID(ctx, id)
	if err != nil {
		log.WithError(err).Errorln("Failed get category")
		return err
	}

	if err := s.Repository.Delete(ctx, *category); err != nil {
		log.WithError(err).Errorln("Failed delete category")
		return err
	}

	return nil
}
