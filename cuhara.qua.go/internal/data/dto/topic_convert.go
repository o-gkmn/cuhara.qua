package dto

import "cuhara.qua.go/internal/types"

func (t *TopicDTO) ToTypes() *types.TopicResponse {
	return &types.TopicResponse{
		ID:   t.ID,
		Name: t.Name,
	}
}

func (c *CreateTopicResponse) ToTypes() *types.CreateTopicResponse {
	return &types.CreateTopicResponse{
		ID: c.ID,
	}
}

func (u *UpdateTopicResponse) ToTypes() *types.UpdateTopicResponse {
	return &types.UpdateTopicResponse{
		ID: u.ID,
	}
}

func (d *DeleteTopicResponse) ToTypes() *types.DeleteTopicResponse {
	return &types.DeleteTopicResponse{
		ID: d.ID,
	}
}

func (s *SubTopicDTO) ToTypes() *types.SubTopicResponse {
	return &types.SubTopicResponse{
		ID:   s.ID,
		Name: s.Name,
		Topic: s.Topic.ToTypes(),
	}
}

func (c *CreateSubTopicResponse) ToTypes() *types.CreateSubTopicResponse {
	return &types.CreateSubTopicResponse{
		ID: c.ID,
	}
}

func (u *UpdateSubTopicResponse) ToTypes() *types.UpdateSubTopicResponse {
	return &types.UpdateSubTopicResponse{
		ID: u.ID,
	}
}

func (d *DeleteSubTopicResponse) ToTypes() *types.DeleteSubTopicResponse {
	return &types.DeleteSubTopicResponse{
		ID: d.ID,
	}
}