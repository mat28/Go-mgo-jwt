package models



func GenerateTemplateMandrill(template string, variable string) (map[string]interface{}) {

	switch template {
		case "activation" : templateContent := map[string]interface{}{"TOKEN": variable,"CODE": "Test"}
			            return templateContent


	}
	return nil
}