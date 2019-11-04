package product

import (
	"context"

	"github.com/dynastymasra/privy/config"
	"github.com/dynastymasra/privy/domain"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Repository domain.ProductRepository
}

func NewService(repo domain.ProductRepository) Service {
	return Service{Repository: repo}
}

func (s Service) Create(ctx context.Context, product domain.Product) (*domain.Product, error) {
	log := logrus.WithFields(logrus.Fields{
		config.RequestID: ctx.Value(config.HeaderRequestID),
		"product":        product,
	})

	res, err := s.Repository.Create(ctx, product)
	if err != nil {
		log.WithError(err).Errorln("Failed create new product")

		return nil, err
	}

	result, err := s.Repository.FindByID(ctx, res.ID)
	if err != nil {
		log.WithError(err).Errorln("Failed get product by id")
		return nil, err
	}

	return result, nil
}

func (s Service) FindByID(ctx context.Context, id int) (*domain.Product, error) {
	log := logrus.WithFields(logrus.Fields{
		config.RequestID: ctx.Value(config.HeaderRequestID),
		"id":             id,
	})

	product, err := s.Repository.FindByID(ctx, id)
	if err != nil {
		log.WithError(err).Errorln("Failed get product by id")
		return nil, err
	}

	return product, nil
}

func (s Service) Fetch(ctx context.Context, from, size int) ([]domain.Product, error) {
	log := logrus.WithFields(logrus.Fields{
		config.RequestID: ctx.Value(config.HeaderRequestID),
		"from":           from,
		"size":           size,
	})

	products, err := s.Repository.Fetch(ctx, from, size)
	if err != nil {
		log.WithError(err).Errorln("Failed fetch product")
		return nil, err
	}

	return products, nil
}

func (s Service) Update(ctx context.Context, id int, product domain.Product) (*domain.Product, error) {
	log := logrus.WithFields(logrus.Fields{
		config.RequestID: ctx.Value(config.HeaderRequestID),
		"after":          product,
		"id":             id,
	})

	prod, err := s.Repository.FindByID(ctx, id)
	if err != nil {
		log.WithError(err).Errorln("Failed get product by id")
		return nil, err
	}

	product.ID = id
	if err := s.Repository.Update(ctx, product); err != nil {
		log.WithField("before", prod).WithError(err).Errorln("Failed update product")
		return nil, err
	}

	result, err := s.Repository.FindByID(ctx, id)
	if err != nil {
		log.WithError(err).Errorln("Failed get product by id")
		return nil, err
	}

	return result, nil
}

func (s Service) Delete(ctx context.Context, id int) error {
	log := logrus.WithFields(logrus.Fields{
		config.RequestID: ctx.Value(config.HeaderRequestID),
		"id":             id,
	})

	product, err := s.Repository.FindByID(ctx, id)
	if err != nil {
		log.WithError(err).Errorln("Failed get product")
		return err
	}

	if err := s.Repository.Delete(ctx, *product); err != nil {
		log.WithError(err).Errorln("Failed delete product")
		return err
	}

	return nil
}
