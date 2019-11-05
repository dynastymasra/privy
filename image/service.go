package image

import (
	"context"

	"github.com/dynastymasra/privy/config"
	"github.com/dynastymasra/privy/domain"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Repository domain.ImageRepository
}

func NewService(repo domain.ImageRepository) Service {
	return Service{Repository: repo}
}

func (s Service) Create(ctx context.Context, image domain.Image) (*domain.Image, error) {
	log := logrus.WithFields(logrus.Fields{
		config.RequestID: ctx.Value(config.HeaderRequestID),
		"image":          image,
	})

	res, err := s.Repository.Create(ctx, image)
	if err != nil {
		log.WithError(err).Errorln("Failed create new image")

		return nil, err
	}

	result, err := s.Repository.FindByID(ctx, res.ID)
	if err != nil {
		log.WithError(err).Errorln("Failed get image by id")
		return nil, err
	}

	return result, nil
}

func (s Service) FindByID(ctx context.Context, id int) (*domain.Image, error) {
	log := logrus.WithFields(logrus.Fields{
		config.RequestID: ctx.Value(config.HeaderRequestID),
		"id":             id,
	})

	image, err := s.Repository.FindByID(ctx, id)
	if err != nil {
		log.WithError(err).Errorln("Failed get image by id")
		return nil, err
	}

	return image, nil
}

func (s Service) Fetch(ctx context.Context, from, size int) ([]domain.Image, error) {
	log := logrus.WithFields(logrus.Fields{
		config.RequestID: ctx.Value(config.HeaderRequestID),
		"from":           from,
		"size":           size,
	})

	images, err := s.Repository.Fetch(ctx, from, size)
	if err != nil {
		log.WithError(err).Errorln("Failed fetch image")
		return nil, err
	}

	return images, nil
}

func (s Service) Update(ctx context.Context, id int, image domain.Image) (*domain.Image, error) {
	log := logrus.WithFields(logrus.Fields{
		config.RequestID: ctx.Value(config.HeaderRequestID),
		"after":          image,
		"id":             id,
	})

	image.ID = id
	if err := s.Repository.Update(ctx, image); err != nil {
		log.WithError(err).Errorln("Failed update image")
		return nil, err
	}

	result, err := s.Repository.FindByID(ctx, id)
	if err != nil {
		log.WithError(err).Errorln("Failed get image by id")
		return nil, err
	}

	return result, nil
}

func (s Service) Delete(ctx context.Context, id int) error {
	log := logrus.WithFields(logrus.Fields{
		config.RequestID: ctx.Value(config.HeaderRequestID),
		"id":             id,
	})

	image, err := s.Repository.FindByID(ctx, id)
	if err != nil {
		log.WithError(err).Errorln("Failed get image")
		return err
	}

	if err := s.Repository.Delete(ctx, *image); err != nil {
		log.WithError(err).Errorln("Failed delete image")
		return err
	}

	return nil
}
