package main

import "github.com/m0nadic/moquer/models"

func SampleConfig() *models.Config {
	return &models.Config{
		Version: "1.0",
		Services: map[string]*models.Service{
			"payments" : {
				Port: 9910,
				Routes: []*models.Route{
					{
						Method: "GET",
						Pattern: "/v1/payments",
						Response: &models.Response{
							Type:  "file",
							Value: "payments/list.json",
							Status: 200,
						},
					},
					{
						Method: "GET",
						Pattern: "/v1/payments/:id",
						Response: &models.Response{
							Type:  "file",
							Value: "payments/detail.json",
							Status: 200,
						},
					},
					{
						Method: "GET",
						Pattern: "/v1/ping",
						Response:&models.Response{
							Type:  "string",
							Value: "PONG",
							Status: 200,
						},
					},
				},
			},
			"invoices": {
				Port: 9920,
				Routes: []*models.Route{
					{
						Method: "GET",
						Pattern: "/v1/invoices",
						Response: &models.Response{
							Type: "file",
							Value: "invoices/list.json",
							Status: 200,
						},
					},
					{
						Method: "GET",
						Pattern: "/v1/invoices/:id",
						Response: &models.Response{
							Type: "file",
							Value: "invoices/detail.json",
							Status: 200,
						},
					},
				},
			},
		},
	}
}

