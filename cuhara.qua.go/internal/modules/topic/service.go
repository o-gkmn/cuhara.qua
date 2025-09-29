package topic

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"cuhara.qua.go/internal/api/httperrors"
	"cuhara.qua.go/internal/config"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/models"
	"cuhara.qua.go/internal/util"
	"github.com/aarondl/null/v8"
	"github.com/aarondl/sqlboiler/v4/boil"
)

type Service struct {
	db     *sql.DB
	config config.Server
}

func NewService(config config.Server, db *sql.DB) *Service {
	return &Service{
		config: config,
		db:     db,
	}
}

func (s *Service) GetAll(ctx context.Context) ([]dto.TopicDTO, error) {
	log := util.LogFromContext(ctx).With().Str("function", "GetTopics").Logger()

	topics, err := models.Topics().All(ctx, s.db)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get topics")
		return nil, err
	}

	topicDTOs := make([]dto.TopicDTO, len(topics))
	for i, topic := range topics {
		topicDTOs[i] = dto.TopicDTO{
			ID:   topic.ID,
			Name: topic.Name,
		}
	}

	return topicDTOs, nil
}

func (s *Service) Create(ctx context.Context, request dto.CreateTopicRequest) (dto.CreateTopicResponse, error) {
	log := util.LogFromContext(ctx).With().Str("function", "CreateTopic").Logger()

	topic := models.Topic{
		Name: request.Name,
	}

	err := topic.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		log.Error().Err(err).Msg("Failed to create topic")
		return dto.CreateTopicResponse{}, err
	}

	return dto.CreateTopicResponse{ID: topic.ID}, nil
}

func (s *Service) Update(ctx context.Context, request dto.UpdateTopicRequest) (dto.UpdateTopicResponse, error) {
	log := util.LogFromContext(ctx).With().Str("function", "UpdateTopic").Logger()

	topic, err := models.FindTopic(ctx, s.db, request.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error().Err(err).Msg("Topic not found")
			return dto.UpdateTopicResponse{}, httperrors.ErrTopicNotFound
		}

		log.Error().Err(err).Msg("Failed to find topic")
		return dto.UpdateTopicResponse{}, err
	}

	changed := false
	if request.Name != nil && topic.Name != *request.Name {
		topic.Name = *request.Name
		changed = true
	}

	if !changed {
		return dto.UpdateTopicResponse{ID: topic.ID}, nil
	}

	topic.UpdatedAt = null.TimeFrom(time.Now().UTC())
	_, err = topic.Update(ctx, s.db, boil.Whitelist(
		models.TopicColumns.Name,
		models.TopicColumns.UpdatedAt,
	))
	if err != nil {
		log.Error().Err(err).Msg("Failed to update topic")
		return dto.UpdateTopicResponse{}, err
	}

	return dto.UpdateTopicResponse{ID: topic.ID}, nil
}

func (s *Service) Delete(ctx context.Context, request dto.DeleteTopicRequest) (dto.DeleteTopicResponse, error) {
	log := util.LogFromContext(ctx).With().Str("function", "DeleteTopic").Logger()

	topic, err := models.FindTopic(ctx, s.db, request.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error().Err(err).Msg("Topic not found")
			return dto.DeleteTopicResponse{}, httperrors.ErrTopicNotFound
		}

		log.Error().Err(err).Msg("Failed to find topic")
		return dto.DeleteTopicResponse{}, err
	}

	_, err = topic.Delete(ctx, s.db)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete topic")
		return dto.DeleteTopicResponse{}, err
	}

	return dto.DeleteTopicResponse{ID: topic.ID}, nil
}

func (s *Service) GetSubTopics(ctx context.Context, request dto.GetSubTopicsRequest) (dto.SubTopicDTO, error) {
	panic("not implemented")
}

func (s *Service) CreateSubTopic(ctx context.Context, request dto.CreateSubTopicRequest) (dto.CreateSubTopicResponse, error) {
	panic("not implemented")
}

func (s *Service) UpdateSubTopic(ctx context.Context, request dto.UpdateSubTopicRequest) (dto.UpdateSubTopicResponse, error) {
	panic("not implemented")
}

func (s *Service) DeleteSubTopic(ctx context.Context, request dto.DeleteSubTopicRequest) (dto.DeleteSubTopicResponse, error) {
	panic("not implemented")
}