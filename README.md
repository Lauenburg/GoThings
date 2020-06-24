# The Things Network - Project

The repository is composed of the three subtasks:\
    - Delivering a server which counts the requests its handles and can return this number when requested\
    - Implement a microservice that keep track of the temperature measurements published by a TTNode\
    - Combine subtask one and two to allowing requests to the server which return the most current temperature measurement\
    
##### Tree Structure     

Three folders one for each subtask.
RCServer relates to the first, OTemperature to the second and IntegrateRCServer to the third sub-task. 
Each folder holds its own makefile.

#### Build and execute
###### Building the project:
- The makefiles expect that vgo is installed to manage the dependencies
    - Download vgo: go get -u golang.org/x/vgo
    - Link the vgo binary to your path
    - For more information see: https://github.com/golang/vgo
- Make sure that the directory is within GOPATH for vgo to be able to determine the module path
- Run "make all" to build

###### Execute binaries:
To execute RCServer run:  ./RCServer\
To execute the OTemperature and the IntegrateRCServer tasks run: ./filname &lt;devId&gt; &lt;appId&gt; &lt;appAccessKey&gt;

#### Port and handles

The RCServer and IntegrateRCServer run on port 8080\
Receive the counter with "/count" in case of RCServer and IntegrateRCServer.\
Receive the office temperature via "/temperature/appID/devID" in case of the IntegrateRCServer.

