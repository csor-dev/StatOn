package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/trycourier/courier-go/v2"
)

type StrapiEvent struct {
	Event     string   `json:"event"`
	CreatedAt string   `json:"created_at"`
	Model     string   `json:"model"`
	Entry     Incident `json:"entry"`
}

type Incident struct {
	Issue_name           string `json:"issue_name"`
	Issue_start          string `json:"issue_start"`
	Issue_end            string `json:"issue_end"`
	Status               string `json:"status"`
	Incident_effect      string `json:"incident_effect"`
	Incident_description string `json:"incident_description"`
}

func sendmail() {
	API_KEY := os.Getenv("API_KEY")

	client := courier.CreateClient(API_KEY, nil)

	requestID, err := client.SendMessage(
		context.Background(),
		courier.SendMessageRequestBody{
			Message: map[string]interface{}{
				"to": map[string]string{
					"email": "imjuststarting@gmail.com",
				},
				"template": "8BNDDTAXAQ4GJMQ87BBA30FRHZAE",
				"data": map[string]string{
					"recipientName": "Courier Boy",
				},
			},
		},
	)

	if err != nil {
		log.Fatal(err)
	}
	log.Println(requestID)
}

func main() {
	app := fiber.New()

	app.Post("/webhook", func(c *fiber.Ctx) error {
		strapiEvent := new(StrapiEvent)
		c.BodyParser(strapiEvent)

		log.Println("----------------------------")

		log.Println("Strapi Event received")
		log.Println("Event => ", strapiEvent.Event)
		log.Println("Created at => ", strapiEvent.CreatedAt)
		log.Println("Model => ", strapiEvent.Model)
		if strapiEvent.Model == "incident" {
			sendmail()
			// log.Println("Entry Name => ", strapiEvent.Entry.Issue_name)
			// log.Println("Entry Started => ", strapiEvent.Entry.Issue_start)
			// log.Println("Entry Ended => ", strapiEvent.Entry.Issue_end)
			// log.Println("Entry Status => ", strapiEvent.Entry.Status)
			// log.Println("Entry Description => ", strapiEvent.Entry.Incident_description)
			// log.Println("Entry Effect => ", strapiEvent.Entry.Incident_effect)

		}
		log.Println("-----------------------------")
		return c.SendString("OK!")

	})

	log.Fatal(app.Listen(":5001"))

}
