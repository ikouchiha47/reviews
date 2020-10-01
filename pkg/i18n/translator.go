// mocked i18n
package i18n

type Translator struct {
}

func (tr Translator) Message(lang, key string, templateData map[string]interface{}) string {
	// key = fmt.Sprintf("%s_message", key)
	return key
}

func (tr Translator) Title(lang, key string, templateData map[string]interface{}) string {
	// key = fmt.Sprintf("%s_title", key)
	return key
}
