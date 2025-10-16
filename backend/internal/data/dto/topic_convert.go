package dto

import "cuhara.qua.go/internal/types"

func (t *TopicDTO) ToTypes() *types.TopicResponse {
	return &types.TopicResponse{
		Id:   &t.ID,
		Name: &t.Name,
	}
}

func (c *CreateTopicResponse) ToTypes() *types.CreateTopicResponse {
	return &types.CreateTopicResponse{
		Id: &c.ID,
	}
}

func (u *UpdateTopicResponse) ToTypes() *types.UpdateTopicResponse {
	return &types.UpdateTopicResponse{
		Id: &u.ID,
	}
}

func (d *DeleteTopicResponse) ToTypes() *types.DeleteTopicResponse {
	return &types.DeleteTopicResponse{
		Id: &d.ID,
	}
}

func (s *SubTopicDTO) ToTypes() *types.SubTopicResponse {
	return &types.SubTopicResponse{
		Id:   &s.ID,
		Name: &s.Name,
		Topic: s.Topic.ToTypes(),
	}
}

func (c *CreateSubTopicResponse) ToTypes() *types.CreateSubTopicResponse {
	return &types.CreateSubTopicResponse{
		Id: &c.ID,
	}
}

func (u *UpdateSubTopicResponse) ToTypes() *types.UpdateSubTopicResponse {
	return &types.UpdateSubTopicResponse{
		Id: &u.ID,
	}
}

func (d *DeleteSubTopicResponse) ToTypes() *types.DeleteSubTopicResponse {
	return &types.DeleteSubTopicResponse{
		Id: &d.ID,
	}
}