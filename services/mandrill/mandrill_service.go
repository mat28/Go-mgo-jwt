package mandrill_service

import (
	m "api-uptoo-jwt/core/mandrill"
	mc "github.com/keighl/mandrill"
)

func SendTemplate(emailAddress string,lastName string,firstName string,templateName string, templateContent map[string]interface{},fromEmail string, fromName string, subject string ) ([]*mc.Response) {
	client := m.InitMandrill()
	message := m.GenerateMessage(emailAddress,lastName,firstName,templateContent,fromEmail,fromName,subject)
	responses := m.SendTemplate(client,message,templateName)
	return responses
}
