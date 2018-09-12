# TTN Intern 

This repository contains the answer to a case given as part of an application process.

The case is composed of the following three subtasks:
    - Delivering a server which counts the requests its handles and can return this number when requested
    - Implement a microservice that keep track of the temperature measurements published by a TTNode
    -Combine subtask one and two to allowing requests to the server which return the most current temperature measurement
    
##### Tree Structure     

The repository contains three folders one for each subtask.
RCServer relates to the first, OTemperature to the second and IntegrateRCServer to the third sub-task. Each folder holds its own makefile.

#### Build and execute
###### Building the project:
- This makefile expects that vgo installed to manage the dependencies
    - Download vgo: go get -u golang.org/x/vgo
    - Linke the vgo binary to your path
    - For more information see: https://github.com/golang/vgo
- Make sure that the directory is within GOPATH for vgo to be able to determine the module path
- Run "make all" to build

###### Execute binarys:
To execute RCServer run:  ./RCServer
To execute the OTemperature and the IntegrateRCServer tasks run: ./filname <devId> <appId> <appAccessKey>