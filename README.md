# restapi
This is a demo project, It implements the sample REST API (CRUD) via pure Go language. 


1. Design
user rquest -> handler 
                 -> service
                      ->common task, ConfigService(load properties file, and config backend server)
                          ->resource(app.properties)
                      ->general task,UserService(access database)
                      <-response
                <-response
            <-response

handler,  process user' reqeust, can call more then on services 
service,  service, process a kind of task, include common and general services currently. in the future can split common service to a new package. 

2. testing in development 
go run main.go

