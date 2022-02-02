# Orchestrator Service in golang
Orchestrator service which would read any request it receives and forwards it to other orchestrator services and data services.

# Setup
### Cloning the Repository
    git clone https://github.com/AnshVM/orchestrator-service.git
    cd orchestrator-service

#### Run two orchestrator instances on ports `9000` and `9001`

    go run server.go --port=9000
    
In a **new terminal** cd into the orchestrator-service repository and run the following command

    go run server.go --port=9001
#### Starting the dummy data service
In a **new terminal** cd into the orchestrator-service repository and run the following commands

    cd datamock
    go run dummy_service.go
 
#### Running the client
In a **new terminal** cd into the orchestrator-service repository and run the following commands

    cd client
    go run main.go
    
 ## Development environment 
 
 **OS system used for development**  - `Ubuntu - 20.04 (wsl)` 
 
 **GOPATH** - `/home/ansh/Go` 
 
 **Repository path on local** - `~/dev/orchestrator-service`
  
