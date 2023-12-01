// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"path/filepath"

// 	uc "github.com/ElijahB09/Gundam-Displays/uc"
// 	mqtt "github.com/eclipse/paho.mqtt.golang"
// 	"github.com/joho/godotenv"
// )

// type Gundam struct {
// 	name  string
// 	topic string
// }

// var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
// 	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
// 	switch msg.Topic() {
// 	case "gundam/uc/f91/f91gundam":
// 		uc.ToggleF91(msg.Payload())
// 		break
// 	case "gundam/uc/unicorn/unicorngundam":

// 		break
// 	}
// }

// var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
// 	fmt.Println("Connected")
// }

// var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
// 	fmt.Printf("Connect lost: %v", err)
// }

// func init() {
// 	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	environmentPath := filepath.Join(dir, ".env")
// 	err = godotenv.Load(environmentPath)
// }

// func main() {
// 	var broker = os.Getenv("PI_MQTT_BROKER")
// 	if broker == "" {
// 		fmt.Errorf("Secret not recieved")
// 	}
// 	var port = 1883

// 	var gundams []Gundam
// 	gundams = append(gundams, Gundam{name: "f91", topic: "gundam/uc/f91/f91gundam"})
// 	gundams = append(gundams, Gundam{name: "unicorn", topic: "gundam/uc/unicorn/unicorngundam"})

// 	opts := mqtt.NewClientOptions()
// 	opts.AddBroker(fmt.Sprintf("mqtt://%s:%d", broker, port))
// 	opts.SetClientID("mqtt_client_f91")
// 	opts.SetUsername("f91")
// 	password, isThere := os.LookupEnv("GUNDAM_F91_PASSWORD")
// 	if !isThere {
// 		fmt.Errorf("Something big gone wrong with .env")
// 	}
// 	opts.SetPassword(password)
// 	opts.SetDefaultPublishHandler(messagePubHandler)
// 	opts.OnConnect = connectHandler
// 	opts.OnConnectionLost = connectLostHandler
// 	client := mqtt.NewClient(opts)
// 	if token := client.Connect(); token.Wait() && token.Error() != nil {
// 		panic(token.Error())
// 	}
// 	sub(client, "gundam/uc/f91/f91gundam")
// }

// func sub(client mqtt.Client, mqtt_topic string) {
// 	token := client.Subscribe(mqtt_topic, 1, nil)
// 	token.Wait()
// 	fmt.Printf("Subscribed to topic: %s", mqtt_topic)
// }
