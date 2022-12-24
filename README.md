# BookStore-User-Apis

## Project Requirements
- <b>Go>=1.19</b>.
- <b>MySQL</b> as Database Machine

### Create env file named .env in the root directory to get neccessary informations
````
mysql_db_username="test"
mysql_db_password="test"
mysql_db_host="127.0.0.1:3306"
mysql_schema_name="users_db"
````

## Necessary Commands to run the project
### Add missing and/or remove unused modules 
````
go mod tidy
go run main.go
````

## Module Definition.
- <b>App</b> is for building our app, basically for building our webserver.
- <b>Controllers</b> is like Views for our application like Django. This will handle request and sending request accross different services in order to handle and process.
- <b>Domain</b> is the Core for our entire microservice.
- <b>Service</b> will hold entire bussiness logics for our application.
- <b>Data Sources </b> All database tables are here.