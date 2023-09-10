package BuilderHttpClient

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
)

func encodeDataFormValue(dataFormValue interface{}) string {
	switch dataForm := dataFormValue.(type) {
	case map[string]string:
		formData := make([]string, len(dataForm))
		i := 0
		for k, v := range dataForm {
			formData[i] = fmt.Sprintf("%s=%s", k, v)
			i++
		}
		return strings.Join(formData, "&")

	case map[string]interface{}:
		formData := make([]string, len(dataForm))
		i := 0
		for k, v := range dataForm {
			formData[i] = fmt.Sprintf("%s=%v", k, v)
			i++
		}
		return strings.Join(formData, "&")

	case string:
		formDataString, err := url.QueryUnescape(dataForm)
		if err != nil {
			return ""
		}
		return formDataString

	case url.Values:
		return dataForm.Encode()
	}
	return ""
}
func encodeJsonDataValue(jsonData any) string {
	switch v := jsonData.(type) {
	case nil:
		return ""
	case string:
		if isValidJSON(v) {
			return v
		}
	case []byte:
		if isValidJSON(string(v)) {
			return string(v)
		}

	default:
		jsonBytes, err := json.Marshal(jsonData)
		if err != nil {
			log.Printf("序列化失败: %s", err)
			log.Printf("序列化失败的对象: %v", jsonData)
			return ""
		}
		return string(jsonBytes)
	}
	return ""
}

func isValidJSON(s string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(s), &js) == nil
}
