package module

import (
	"log"

	"gopkg.in/yaml.v2"
)

type ResponseSetting struct {
	People  map[int64][]string
	KeyWord map[string][]string
}

var ResponseSettingMap map[string]*ResponseSetting = make(map[string]*ResponseSetting)

func GetResponseSetting(responseSettingFileName string) {
	if _, ok := ResponseSettingMap[responseSettingFileName]; !ok {
		responseSettingFile := CurrentWorkingDirectory.Join("configuration", responseSettingFileName)
		if responseSettingFile.IsExist() {
			fileData, err := responseSettingFile.ReadFile()
			if err != nil {
				log.Fatalln("Error appear when reading file " + responseSettingFile.String())
			}
			responseSetting := &ResponseSetting{make(map[int64][]string), make(map[string][]string)}
			err = yaml.Unmarshal(fileData, responseSetting)
			if err != nil {
				log.Fatalln("Error appear when unmarshalling file " + responseSettingFile.String())
			}
			ResponseSettingMap[responseSettingFileName] = responseSetting
		} else {
			log.Fatalln("Expected file " + responseSettingFile.String() + " does not exist")
		}
	}
}
