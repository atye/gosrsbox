package api

import (
	"fmt"

	"github.com/atye/gosrsbox/osrsbox/db"
)

type ItemsResponse struct {
	Items []db.Item `json:"_items"`
	Meta  struct {
		Page       int `json:"page"`
		Total      int `json:"total"`
		MaxResults int `json:"max_results"`
	} `json:"_meta"`
	Links struct {
		Parent struct {
			Title string `json:"title"`
			Href  string `json:"href"`
		} `json:"parent"`
		Self struct {
			Title string `json:"title"`
			Href  string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Error *APIError `json:"_error"`
}

type MonstersResponse struct {
	Monsters []db.Monster `json:"_items"`
	Meta     struct {
		Page       int `json:"page"`
		Total      int `json:"total"`
		MaxResults int `json:"max_results"`
	} `json:"_meta"`
	Links struct {
		Parent struct {
			Title string `json:"title"`
			Href  string `json:"href"`
		} `json:"parent"`
		Self struct {
			Title string `json:"title"`
			Href  string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Error *APIError `json:"_error"`
}

type PrayersResponse struct {
	Prayers []db.Prayer `json:"_items"`
	Meta    struct {
		Page       int `json:"page"`
		Total      int `json:"total"`
		MaxResults int `json:"max_results"`
	} `json:"_meta"`
	Links struct {
		Parent struct {
			Title string `json:"title"`
			Href  string `json:"href"`
		} `json:"parent"`
		Self struct {
			Title string `json:"title"`
			Href  string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Error *APIError `json:"_error"`
}

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err APIError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", err.Code, err.Message)
}
