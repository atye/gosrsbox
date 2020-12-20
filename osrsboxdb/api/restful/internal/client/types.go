package client

import (
	"fmt"

	"github.com/atye/gosrsbox/osrsboxdb"
)

type itemsResponse struct {
	Items []osrsboxdb.Item `json:"_items"`
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
	Error *apiError `json:"_error"`
}

type monstersResponse struct {
	Monsters []osrsboxdb.Monster `json:"_items"`
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
	Error *apiError `json:"_error"`
}

type prayersResponse struct {
	Prayers []osrsboxdb.Prayer `json:"_items"`
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
	Error *apiError `json:"_error"`
}

type apiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err apiError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", err.Code, err.Message)
}
