package module

import (
	"log"

	"gopkg.in/yaml.v2"
)

type BotSettingStruct struct {
	Token string
}

var BotSetting = getBotSetting()

func getBotSetting() *BotSettingStruct {
	botSettingFile := CurrentWorkingDirectory.Join("configuration", "bot-setting.yml")
	if botSettingFile.IsExist() {
		fileData, err := botSettingFile.ReadFile()
		if err != nil {
			log.Fatalln("Error appear when reading file " + botSettingFile.String())
		}
		botSetting := &BotSettingStruct{}
		err = yaml.Unmarshal(fileData, botSetting)
		if err != nil {
			log.Fatalln("Error appear when unmarshalling file " + botSettingFile.String())
		}
		return botSetting
	} else {
		log.Fatalln("Expected file " + botSettingFile.String() + " does not exist")
		return nil
	}
}
