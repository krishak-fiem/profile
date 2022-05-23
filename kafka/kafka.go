package kafka

import (
	"encoding/json"
	"fmt"
	constants2 "github.com/krishak-fiem/constants/go"
	kafka "github.com/krishak-fiem/kafka/go"
	kafkamodels "github.com/krishak-fiem/models/go/kafka"
	profilemodels "github.com/krishak-fiem/models/go/profile"
	"github.com/krishak-fiem/profile/constants"
	kafka2 "github.com/segmentio/kafka-go"
)

func UserCreatedReader() {
	go kafka.MessageReader([]string{constants.KAFKA_BROKER_URL}, string(constants2.USER_CREATED), UserCreatedHandler)
}

func UserCreatedHandler(message kafka2.Message) error {
	m := new(kafkamodels.UserCreatedMessage)
	err := json.Unmarshal(message.Value, m)
	if err != nil {
		fmt.Println(err.Error())
	}

	profile := new(profilemodels.Profile)
	profile.Email = m.Email
	profile.Name = m.Name
	err = profile.CreateProfile()
	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
