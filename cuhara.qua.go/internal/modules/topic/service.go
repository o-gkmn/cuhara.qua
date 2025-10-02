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
	"github.com/aarondl/sqlboiler/v4/queries/qm"
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

	tenantID, err := util.TenantIDFromContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tenant id from context")
		return nil, err
	}

	topics, err := models.Topics(
		models.TopicWhere.TenantID.EQ(tenantID),
	).All(ctx, s.db)
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

	tenantID, err := util.TenantIDFromContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tenant id from context")
		return dto.CreateTopicResponse{}, err
	}

	exists, err := models.Topics(
		models.TopicWhere.Name.EQ(request.Name),
		models.TopicWhere.TenantID.EQ(tenantID),
	).Exists(ctx, s.db)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check whether topic exists")
		return dto.CreateTopicResponse{}, err
	}

	if exists {
		log.Debug().Str("name", request.Name).Msg("Topic already exists")
		return dto.CreateTopicResponse{}, httperrors.ErrConflictTopicAlreadyExists
	}

	topic := models.Topic{
		Name:      request.Name,
		TenantID:  tenantID,
	}

	err = topic.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		log.Error().Err(err).Msg("Failed to create topic")
		return dto.CreateTopicResponse{}, err
	}

	return dto.CreateTopicResponse{ID: topic.ID}, nil
}

func (s *Service) Update(ctx context.Context, request dto.UpdateTopicRequest) (dto.UpdateTopicResponse, error) {
	log := util.LogFromContext(ctx).With().Str("function", "UpdateTopic").Logger()

	tenantID, err := util.TenantIDFromContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tenant id from context")
		return dto.UpdateTopicResponse{}, err
	}

	topic, err := models.Topics(
		models.TopicWhere.ID.EQ(request.ID),
		models.TopicWhere.TenantID.EQ(tenantID),
	).One(ctx, s.db)
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
		exists, err := models.Topics(
			models.TopicWhere.Name.EQ(*request.Name),
			models.TopicWhere.TenantID.EQ(tenantID),
			models.TopicWhere.ID.NEQ(request.ID),
		).Exists(ctx, s.db)
		if err != nil {
			log.Error().Err(err).Msg("Failed to check whether topic exists")
			return dto.UpdateTopicResponse{}, err
		}

		if exists {
			log.Debug().Str("name", *request.Name).Msg("Topic already exists")
			return dto.UpdateTopicResponse{}, httperrors.ErrConflictTopicAlreadyExists
		}

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

	tenantID, err := util.TenantIDFromContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tenant id from context")
		return dto.DeleteTopicResponse{}, err
	}

	topic, err := models.Topics(
		models.TopicWhere.ID.EQ(request.ID),
		models.TopicWhere.TenantID.EQ(tenantID),
	).One(ctx, s.db)
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

func (s *Service) GetSubTopics(ctx context.Context, request dto.GetSubTopicsRequest) ([]dto.SubTopicDTO, error) {
	log := util.LogFromContext(ctx).With().Str("function", "GetSubTopics").Logger()

	tenantID, err := util.TenantIDFromContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tenant id from context")
		return nil, err
	}

	subTopics, err := models.SubTopics(
		models.SubTopicWhere.TopicID.EQ(request.TopicID),
		models.SubTopicWhere.TenantID.EQ(tenantID),
		qm.Load(models.SubTopicRels.Topic),
	).All(ctx, s.db)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get sub topics")
		return nil, err
	}

	subTopicDTOs := make([]dto.SubTopicDTO, len(subTopics))
	for i, subTopic := range subTopics {
		subTopicDTOs[i] = dto.SubTopicDTO{
			ID:   subTopic.ID,
			Name: subTopic.Name,
			Topic: dto.TopicDTO{
				ID:   subTopic.TopicID,
				Name: subTopic.R.Topic.Name,
			},
		}
	}

	return subTopicDTOs, nil
}

func (s *Service) CreateSubTopic(ctx context.Context, request dto.CreateSubTopicRequest) (dto.CreateSubTopicResponse, error) {
	log := util.LogFromContext(ctx).With().Str("function", "CreateSubTopic").Logger()

	tenantID, err := util.TenantIDFromContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tenant id from context")
		return dto.CreateSubTopicResponse{}, err
	}

	exists, err := models.SubTopics(
		models.SubTopicWhere.Name.EQ(request.Name),
		models.SubTopicWhere.TopicID.EQ(request.TopicID),
		models.SubTopicWhere.TenantID.EQ(tenantID),
	).Exists(ctx, s.db)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check whether sub topic exists")
		return dto.CreateSubTopicResponse{}, err
	}

	if exists {
		log.Debug().Str("name", request.Name).Msg("Sub topic already exists")
		return dto.CreateSubTopicResponse{}, httperrors.ErrConflictSubTopicAlreadyExists
	}

	subTopic := models.SubTopic{
		Name:      request.Name,
		TopicID:   request.TopicID,
		TenantID:  tenantID,
	}

	err = subTopic.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		log.Error().Err(err).Msg("Failed to create sub topic")
		return dto.CreateSubTopicResponse{}, err
	}

	return dto.CreateSubTopicResponse{ID: subTopic.ID}, nil
}

func (s *Service) UpdateSubTopic(ctx context.Context, request dto.UpdateSubTopicRequest) (dto.UpdateSubTopicResponse, error) {
	log := util.LogFromContext(ctx).With().Str("function", "UpdateSubTopic").Logger()

	tenantID, err := util.TenantIDFromContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tenant id from context")
		return dto.UpdateSubTopicResponse{}, err
	}

	subTopic, err := models.SubTopics(
		models.SubTopicWhere.ID.EQ(request.ID),
		models.SubTopicWhere.TopicID.EQ(request.TopicID),
		models.SubTopicWhere.TenantID.EQ(tenantID),
		qm.Load(models.SubTopicRels.Topic),
	).One(ctx, s.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error().Err(err).Msg("Sub topic not found")
			return dto.UpdateSubTopicResponse{}, httperrors.ErrSubTopicNotFound
		}

		log.Error().Err(err).Msg("Failed to find sub topic")
		return dto.UpdateSubTopicResponse{}, err
	}

	changed := false
	if request.Name != nil && subTopic.Name != *request.Name {
		exists, err := models.SubTopics(
			models.SubTopicWhere.Name.EQ(*request.Name),
			models.SubTopicWhere.TopicID.EQ(request.TopicID),
			models.SubTopicWhere.TenantID.EQ(tenantID),
			models.SubTopicWhere.ID.NEQ(request.ID),
		).Exists(ctx, s.db)
		if err != nil {
			log.Error().Err(err).Msg("Failed to check whether sub topic exists")
			return dto.UpdateSubTopicResponse{}, err
		}

		if exists {
			log.Debug().Str("name", *request.Name).Msg("Sub topic already exists")
			return dto.UpdateSubTopicResponse{}, httperrors.ErrConflictSubTopicAlreadyExists
		}

		subTopic.Name = *request.Name
		changed = true
	}

	if !changed {
		return dto.UpdateSubTopicResponse{ID: subTopic.ID}, nil
	}

	subTopic.UpdatedAt = null.TimeFrom(time.Now().UTC())
	_, err = subTopic.Update(ctx, s.db, boil.Whitelist(
		models.SubTopicColumns.Name,
		models.SubTopicColumns.UpdatedAt,
	))
	if err != nil {
		log.Error().Err(err).Msg("Failed to update sub topic")
		return dto.UpdateSubTopicResponse{}, err
	}

	return dto.UpdateSubTopicResponse{ID: subTopic.ID}, nil
}

func (s *Service) DeleteSubTopic(ctx context.Context, request dto.DeleteSubTopicRequest) (dto.DeleteSubTopicResponse, error) {
	log := util.LogFromContext(ctx).With().Str("function", "DeleteSubTopic").Logger()

	tenantID, err := util.TenantIDFromContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tenant id from context")
		return dto.DeleteSubTopicResponse{}, err
	}

	subTopic, err := models.SubTopics(
		models.SubTopicWhere.ID.EQ(request.ID),
		models.SubTopicWhere.TopicID.EQ(request.TopicID),
		models.SubTopicWhere.TenantID.EQ(tenantID),
	).One(ctx, s.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error().Err(err).Msg("Sub topic not found")
			return dto.DeleteSubTopicResponse{}, httperrors.ErrSubTopicNotFound
		}

		log.Error().Err(err).Msg("Failed to find sub topic")
		return dto.DeleteSubTopicResponse{}, err
	}

	_, err = subTopic.Delete(ctx, s.db)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete sub topic")
		return dto.DeleteSubTopicResponse{}, err
	}

	return dto.DeleteSubTopicResponse{ID: subTopic.ID}, nil
}