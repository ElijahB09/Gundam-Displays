package main

import (
	"fmt"
	"os"
	"strings"

	uc "github.com/ElijahB09/Gundam-Displays/uc"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

type Gundam struct {
	name  string
	topic string
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	switch msg.Topic() {
	case "gundam/uc/f91/f91gundam":
		uc.ToggleF91(msg.Payload())
		break
	case "gundam/uc/unicorn/unicorngundam":

		break
	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func init() {
	// Load .env variables into system
	if err := godotenv.Load(); err != nil {
		fmt.Print("No .env file found")
	}
}

func main() {
	var broker, exists = os.LookupEnv("PI_MQTT_BROKER")
	if !exists {
		fmt.Print("Something big gone wrong with .env")
	}
	var port = 1883

	var gundams []Gundam
	gundams = append(gundams, Gundam{name: "f91", topic: "gundam/uc/f91/f91gundam"})
	gundams = append(gundams, Gundam{name: "unicorn", topic: "gundam/uc/unicorn/unicorngundam"})

	for _, element := range gundams {
		opts := mqtt.NewClientOptions()
		opts.AddBroker(fmt.Sprintf("mqtt://%s:%d", broker, port))
		opts.SetClientID("mqtt_client_" + element.name)
		opts.SetUsername(element.name)
		password, isThere := os.LookupEnv("GUNDAM_" + strings.ToUpper(element.name) + "_PASSWORD")
		if !isThere {
			fmt.Errorf("Something big gone wrong with .env")
			break
		}
		opts.SetPassword(password)
		opts.SetDefaultPublishHandler(messagePubHandler)
		opts.OnConnect = connectHandler
		opts.OnConnectionLost = connectLostHandler
		client := mqtt.NewClient(opts)
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
		sub(client, element.topic)
	}
}

func sub(client mqtt.Client, mqtt_topic string) {
	token := client.Subscribe(mqtt_topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", mqtt_topic)
}
