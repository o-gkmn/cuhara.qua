package dto

type TopicDTO struct {
	ID   int64
	Name string
}

type CreateTopicRequest struct {
	Name string
}

type CreateTopicResponse struct {
	ID int64
}

type UpdateTopicRequest struct {
	ID   int64
	Name *string
}

type UpdateTopicResponse struct {
	ID int64
}

type DeleteTopicRequest struct {
	ID int64
}

type DeleteTopicResponse struct {
	ID int64
}

type SubTopicDTO struct {
	ID    int64
	Name  string
	Topic TopicDTO
}

type GetSubTopicsRequest struct {
	TopicID int64
}

type CreateSubTopicRequest struct {
	TopicID int64
	Name    string
}

type CreateSubTopicResponse struct {
	ID int64
}

type UpdateSubTopicRequest struct {
	ID   int64
	Name *string
}

type UpdateSubTopicResponse struct {
	ID int64
}

type DeleteSubTopicRequest struct {
	ID int64
}

type DeleteSubTopicResponse struct {
	ID int64
}
