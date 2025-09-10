package utils

import (
	"encoding/json"
	"log"
	"os"
)

type Data interface {
}

type LoadingData struct {
	MainHeader string `json:"mainHeader"`
	SubHeader  string `json:"subHeader"`
	StartTime  string `json:"startTime"`
	Timestamp  int64  `json:"timestamp"`
}

type PauseData struct {
	HeaderText   string `json:"headerText"`
	DescribeText string `json:"describeText"`
}

func WriteJSONToFile[T LoadingData | PauseData](instance T, filePath string) error {
	//序列化
	jsonBytes, marshalErr := json.MarshalIndent(instance, "", "  ")
	if marshalErr != nil {
		log.Println(marshalErr)
		return marshalErr
	}

	writeErr := os.WriteFile(filePath, jsonBytes, 0666)
	if writeErr != nil {
		log.Println(writeErr)
		return writeErr
	}
	return nil
}

func ReadJSONFromFile[T LoadingData | PauseData](filePath string) (*T, error) {
	//读取之前确保数据文件夹以及文件存在
	EnsureDataDir("ssp-data")
	EnsureDataFile[T](filePath)

	var instance T
	jsonBytes, readErr := os.ReadFile(filePath)
	if readErr != nil {
		log.Println(readErr)
		return nil, readErr
	}

	//反序列化
	unmarshalErr := json.Unmarshal(jsonBytes, &instance)
	if unmarshalErr != nil {
		log.Println(unmarshalErr)
		return nil, unmarshalErr
	}

	return &instance, nil
}

func EnsureDataFile[T LoadingData | PauseData](filePath string) {
	if _, statErr := os.Stat(filePath); statErr != nil {

		file, createErr := os.Create(filePath)
		if createErr != nil {
			log.Println(createErr)
		}
		defer file.Close()

		var instance T
		writeErr := WriteJSONToFile(instance, filePath)
		if writeErr != nil {
			log.Println(writeErr)
		}
	}
	return
}

func EnsureDataDir(dirPath string) {
	if _, statErr := os.Stat(dirPath); statErr != nil {
		mkErr := os.Mkdir(dirPath, 0766)
		if mkErr != nil {
			log.Println(mkErr)
		}
	}
}
