package utils

import (
	"bytes"
	"html/template"

	"github.com/peter6866/foodie/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Item struct {
	ID       string
	Name     string
	ImageUrl string
}

type OrderDetails struct {
	CustomerName string
	Items        []Item
	OrderTime    string
}

func parseTemplate(templatePath string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func SendConfirmationEmail(to string, orderDetails OrderDetails) error {
	from := mail.NewEmail("MealMate", "hjiayu@wustl.edu")
	subject := "Order Received"
	recipient := mail.NewEmail("", to)

	htmlContent, err := parseTemplate("utils/email_templates/order_confirmation.html", orderDetails)
	if err != nil {
		return err
	}

	message := mail.NewSingleEmail(from, subject, recipient, "", htmlContent)
	client := sendgrid.NewSendClient(config.AppConfig.SENDGRID_API_KEY)
	_, err = client.Send(message)

	return err
}
