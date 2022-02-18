package module

import (
	"log"

	"gopkg.in/yaml.v2"
)

var ChannelsSetting = getChannelsSetting()

func getChannelsSetting() map[int64]string {
	channelsSettingFile := CurrentWorkingDirectory.Join("configuration", "channels-setting.yml")
	if channelsSettingFile.IsExist() {
		fileData, err := channelsSettingFile.ReadFile()
		if err != nil {
			log.Fatalln("Error appear when reading file " + channelsSettingFile.String())
		}
		channelsSetting := make(map[int64]string)
		err = yaml.Unmarshal(fileData, channelsSetting)
		if err != nil {
			log.Fatalln("Error appear when unmarshalling file " + channelsSettingFile.String())
		}
		return channelsSetting
	} else {
		log.Fatalln("Expected file " + channelsSettingFile.String() + " does not exist")
		return nil
	}
}
