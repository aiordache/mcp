package client

import (
	"fmt"
)

type SOError struct {
	Id      int    `json:"error_id"`
	Message string `json:"error_message"`
	Name    string `json:"error_name"`
}

func (soErr *SOError) Error() string {
	return fmt.Sprintf("so error: id: %d; message: %s; name %s", soErr.Id, soErr.Message, soErr.Name)
}

type SOCollectiveExternalLink struct {
	Link string `json:"link"`
	Type string `json:"type"`
}

type SOCollective struct {
	Description   string                     `json:"description"`
	ExternalLinks []SOCollectiveExternalLink `json:"external_links"`
	Link          string                     `json:"link"`
	Name          string                     `json:"name"`
	Slug          string                     `json:"slug"`
	Tags          []string                   `json:"tag"`
}

type TagItem struct {
	Collectives      *[]SOCollective `json:"collectives"`
	Count            int             `json:"count"`
	HasSynonyms      bool            `json:"has_synonyms"`
	IsModeratorOnly  bool            `json:"is_moderator_only"`
	IsRequired       bool            `json:"is_required"`
	LastActivityDate *uint64         `json:"last_activity_date"`
	Synonyms         []string        `json:"synonyms"`
	UserId           string          `json:"user_id"`
	Name             string          `json:"name"`
}

type Response struct {
	HasMore        bool `json:"has_more"`
	QuotaMax       int  `json:"quota_max"`
	QuotaRemaining int  `json:"quota_remaining"`
}

type SearchResponse struct {
	Response
	Items []SearchItem `json:"items"`
}

type SearchItem struct {
	Tags             []string `json:"tags"`
	Owner            Owner    `json:"owner"`
	IsAnswered       bool     `json:"is_answered"`
	ViewCount        int      `json:"view_count"`
	AcceptedAnswerID int      `json:"accepted_answer_id"`
	AnswerCount      int      `json:"answer_count"`
	Score            int      `json:"score"`
	LastActivityDate int64    `json:"last_activity_date"`
	CreationDate     int64    `json:"creation_date"`
	QuestionID       int      `json:"question_id"`
	Link             string   `json:"link"`
	Title            string   `json:"title"`
}

type Owner struct {
	AccountID int `json:"account_id"`
	// Reputation   int    `json:"reputation"`
	// UserID       int    `json:"user_id"`
	UserType string `json:"user_type"`
	// ProfileImage string `json:"profile_image"`
	DisplayName string `json:"display_name"`
	Link        string `json:"link"`
}

// Questions response

type AnswerItem struct {
	AnswerID int   `json:"answer_id"`
	Owner    Owner `json:"owner"`
}

type AnswerIDsResponse struct {
	Items []AnswerItem `json:"items"`
}

// Answer Response
type Answer struct {
	AnswerID int       `json:"answer_id"`
	Body     string    `json:"body"`
	Owner    Owner     `json:"owner"`
	Comments []Comment `json:"comments"`
}

type AnswerResponse struct {
	Items []Answer `json:"items"`
}

type Comment struct {
	// Owner         Owner  `json:"owner"`
	Score int `json:"score"`
	// CreationDate  int64         `json:"creation_date"`
	// PostID        int           `json:"post_id"`
	CommentID int    `json:"comment_id"`
	Body      string `json:"body"`
}
