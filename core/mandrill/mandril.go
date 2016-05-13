package mandrill



import (
	m "github.com/keighl/mandrill"
	"log"
	"fmt"
)

type Response struct {
	*m.Response
}

func InitMandrill() *m.Client{
	//client := m.ClientWithKey(os.Getenv("KEY_MANDRILL"))
	client := m.ClientWithKey("k_tsWyvMZNWoerVuF1IgVA")
	return client
}

func GenerateMessage(emailAddress string,lastName string,firstName string, templateContent map[string]interface{},fromEmail string, fromName string, subject string) *m.Message{
	message := &m.Message{}
	message.AddRecipient(emailAddress, firstName + " " + lastName, "to")
	message.FromEmail = fromEmail
	message.FromName = fromName
	message.Subject = subject
	message.GlobalMergeVars = m.MapToVars(templateContent)

	return message
}

func SendTemplate(client *m.Client, message *m.Message, templateName string) []*m.Response{

	response, err := client.MessagesSendTemplate(message,templateName,"")
	if err != nil {
		log.Print(fmt.Println(err.Error()))
	}
	return response
}
