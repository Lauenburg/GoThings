/*
The task was to create a microservice that keeps track of the current temperature.
This script represents a stand-alone program which returns the most recent temperature measurement.
For the third task this file is changed from "package main" to "package oTemper" to be accessible as service, this naturally also includes changing the function from main to OTemp.
The arguments device ID, application ID and application access key are passed to the programm as command line arguments:

./OTemperature <devId> <appId> <appAccessKey> 

*/

package main

import (
"os"
ttnsdk "github.com/TheThingsNetwork/go-app-sdk"
ttnlog "github.com/TheThingsNetwork/go-utils/log"
"github.com/TheThingsNetwork/go-utils/log/apex"
)



const (
    sdkClientName = "OfficeTemperature"
)

func main() {

    //Setting up the logging to Stdout
    log := apex.Stdout() // We use a cli logger at Stdout
    ttnlog.Set(log)      // Set the logger as default for TTN

    //Reading the arguments device ID, application ID and application access key as command line arguments 
    //Do to the securrerty issues the idea of an config file has been omited
    devID := os.Args[1]
    appID := os.Args[2]
    appAccessKey := os.Args[3]



    //Setting up SDK configurations
    //NewCommunityConfig initializes the config
    config := ttnsdk.NewCommunityConfig(sdkClientName)
    config.ClientVersion = "1.0.0" // The version of the application

    //Client configuration, using the given Application ID and Application access key.
    client := config.NewClient(string(appID), string(appAccessKey))
    //closes the client when main returns
    defer client.Close()

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

    // <- sends the UplinkMessage of <-chan *types.UplinkMessage to the varibale uplink
    uplinkMessage := <- uplink
    log.WithField("data", uplinkMessage.PayloadFields["temperature"]).Infof("%s: received uplink", sdkClientName)

}