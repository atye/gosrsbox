package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/atye/gosrsbox/osrsboxapi/api/internal/client/openapi"
)

// "github.com/atye/gosrsbox/osrsboxapi/api/internal/client/openapi-client"
func main() {
	conf := &openapi.Configuration{
		Scheme:     "https",
		HTTPClient: http.DefaultClient,
		Servers: []openapi.ServerConfiguration{
			{
				URL: "api.osrsbox.com",
			},
		},
	}
	api := openapi.NewAPIClient(conf)
	_, resp, err := api.ItemApi.Getitems(context.Background()).Where(`name=="Abyssal whip"}`).Execute()
	defer resp.Body.Close()
	if err.Error() != "" {
		log.Println(string(err.Body()))
		//fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.Getitems``: %v\n", err)
		//fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
		var apiErr openapi.Error
		err := json.NewDecoder(resp.Body).Decode(&apiErr)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("code: %d, message: %s", apiErr.Error.GetCode(), apiErr.Error.GetMessage())
	}
}
