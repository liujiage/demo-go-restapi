# restapi
This is a demo project, It implements the sample REST API (CRUD) via pure Go language. 

<PRE>
1. Design  
user rquest -> handler   
                 -> common 
                 
                 -> service  
                      ->dao, access database task 
                          ->resource(app.properties)
                      <-response  
                <-response  
            <-response  

common, database migration, load properties. 
handler,  process user' reqeust, can call more then on services 
service,  service, process a kind of task, include common and general services currently. in the future can split common service to a new package. 
test, for unit test. the file name shuld be xxx_test.go, some demo for researching 

2. testing in development 
go run main.go

3. database migrate
Install scoop on window:
Set-ExecutionPolicy RemoteSigned -scope CurrentUser
iwr -useb get.scoop.sh | iex
scoop install migrate
migrate create -ext sql -dir database/migration/ -seq init_mg
https://github.com/golang-migrate/migrate/tree/master/database/sqlite3
</PRE>
