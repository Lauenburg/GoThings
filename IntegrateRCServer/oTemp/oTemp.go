/*
Microservice that keeps track of the current temperature.
*/

package oTemp

import (
	ttnsdk "github.com/TheThingsNetwork/go-app-sdk"
	ttnlog "github.com/TheThingsNetwork/go-utils/log"
	"github.com/TheThingsNetwork/go-utils/log/apex"
	"github.com/TheThingsNetwork/ttn/core/types"
)

const (
	sdkClientName = "OfficeTemperature"
)

//this function subscribes to an MQTT broker and returns a channel linking to the subscription
func OTemp(devID, appID string, appAccessKey string) <-chan *types.UplinkMessage {

	//Setting up the logging to Stdout
	log := apex.Stdout() // We use a cli logger at Stdout
	ttnlog.Set(log)      // Set the logger as default for TTN

	//Setting up SDK configurations
	//NewCommunityConfig initializes the config
	config := ttnsdk.NewCommunityConfig(sdkClientName)
	config.ClientVersion = "1.0.1" // The version of the application

	//Client configuration, using the given Application ID and Application access key.
	client := config.NewClient(string(appID), string(appAccessKey))
	//closes the client when main returns
	//defer client.Close()

	// Subscribe to uplink reciving an instance of the ApplicationPubSub interface
	pubsub, err := client.PubSub()
	if err != nil {
		log.WithError(err).Fatalf("%s: could not get application pub/sub", sdkClientName)
	}

	//Retrieving an instance of the interface DevicePubSub which combines the DevicePub and DeviceSub interfaces for the device "devID"
	myNewDeviceSub := pubsub.Device(string(devID))

	//subscribe to uplink messages, logging them to the console as they arrive
	uplink, err := myNewDeviceSub.SubscribeUplink()
	if err != nil {
		log.WithError(err).Fatalf("%s: could not subscribe to uplink messages", sdkClientName)
	}

	// <- sends the UplinkMessage of chan *types.UplinkMessage to the varibale uplink
	return uplink

}
