package dto

type TopicDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CreateTopicRequest struct {
	Name string `json:"name"`
}

type CreateTopicResponse struct {
	ID int64 `json:"id"`
}

type UpdateTopicRequest struct {
	ID   int64   `json:"id"`
	Name *string `json:"name"`
}

type UpdateTopicResponse struct {
	ID int64 `json:"id"`
}

type DeleteTopicRequest struct {
	ID int64 `json:"id"`
}

type DeleteTopicResponse struct {
	ID int64 `json:"id"`
}

type SubTopicDTO struct {
	ID    int64    `json:"id"`
	Name  string   `json:"name"`
	Topic TopicDTO `json:"topic"`
}

type GetSubTopicsRequest struct {
	TopicID int64 `json:"topicId"`
}

type CreateSubTopicRequest struct {
	TopicID int64  `json:"topicId"`
	Name    string `json:"name"`
}

type CreateSubTopicResponse struct {
	ID int64 `json:"id"`
}

type UpdateSubTopicRequest struct {
	ID      int64   `json:"id"`
	TopicID int64   `json:"topicId"`
	Name    *string `json:"name"`
}

type UpdateSubTopicResponse struct {
	ID int64 `json:"id"`
}

type DeleteSubTopicRequest struct {
	ID      int64 `json:"id"`
	TopicID int64 `json:"topicId"`
}

type DeleteSubTopicResponse struct {
	ID int64 `json:"id"`
}
