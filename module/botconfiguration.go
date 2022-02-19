package module

import (
	"log"

	"gopkg.in/yaml.v2"
)

var (
	BotConfiguration = getBotConfiguration()
	ChannelsMapping = getChannelsMapping()
	ResponseHandlingMap map[string]*ResponseHandling = make(map[string]*ResponseHandling)
)

type BotConfigurationStruct struct {
	Token string `yaml:"Token"`
}

type ResponseHandling struct {
	ByPeople  map[int64][]string `yaml:"ByPeople"`
	ByKeyword map[string][]string `yaml:"ByKeyword"`
}

func readConfiguration(configurationFilePath Path, ptr interface{}) {
	if configurationFilePath.IsExist() {
		fileData, err := configurationFilePath.ReadFile()
		if err != nil {
			log.Fatalln("Error appear when reading file " + string(configurationFilePath))
		}
		err = yaml.Unmarshal(fileData, ptr)
		if err != nil {
			log.Fatalln("Error appear when unmarshalling file " + string(configurationFilePath))
		}
	} else {
		log.Fatalln("Expected file " + string(configurationFilePath) + " does not exist")
	}
}

func getBotConfiguration() *BotConfigurationStruct {
	botConfigurationFile := CurrentWorkingDirectory.Join("configuration", "bot-configuration.yml")
	botConfiguration := &BotConfigurationStruct{}
	readConfiguration(botConfigurationFile, botConfiguration)
	return botConfiguration
}

func getChannelsMapping() map[int64]string {
	channelsMappingFile := CurrentWorkingDirectory.Join("configuration", "channels-mapping.yml")
	channelsMapping := make(map[int64]string)
	readConfiguration(channelsMappingFile, channelsMapping)
	return channelsMapping
}

func getResponseHandling(responseHandlingFileName string) *ResponseHandling {
	responseHandlingFile := CurrentWorkingDirectory.Join("configuration", responseHandlingFileName)
	responseHandling := &ResponseHandling{make(map[int64][]string), make(map[string][]string)}
	readConfiguration(responseHandlingFile, responseHandling)
	return responseHandling
}

func GetResponseHandling(responseHandlingFileName string) {
	if _, ok := ResponseHandlingMap[responseHandlingFileName]; !ok {
		ResponseHandlingMap[responseHandlingFileName] = getResponseHandling(responseHandlingFileName)
	}
}
