# restapi
This is a demo project, It implements the sample REST API (CRUD) via pure Go language. 

<PRE>
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
test, for unit test. the file name shuld be xxx_test.go
demo, for studing and researching. 
database, database migration, access database.

2. testing in development 
go run main.go

3. database migrate
Install scoop on window:
Set-ExecutionPolicy RemoteSigned -scope CurrentUser
iwr -useb get.scoop.sh | iex
scoop install migrate
migrate create -ext sql -dir database/migration/ -seq init_mg
</PRE>
