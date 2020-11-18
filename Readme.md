# HOW TO INSTALL

##### 1. Clone the repository

```bash
git clone https://github.com/feedlyy/Koala.id.git
```

##### 2. Install Prerequisites
Install go if you didn't have one
https://golang.org/dl/

- Validator (form request validator)
```bash
go get gopkg.in/go-playground/validator.v10
```
- Postgres
```
go get github.com/lib/pq
```
or you can using docker if you had already install docker, run this command
```
make postgres
```
this command will automatically run docker container
with this credentials and port
```
port: 5432
POSTGRES_USER=root
POSTGRES_PASSWORD=secret

docker image postgres = postgres:latest
```
- Random string generator
```
go get github.com/dchest/uniuri
```

# HOW TO USE

##### 1. Logic Test
go to logic-test folder and run the file, Example :
```
go run Logic-Test/prima.go
```
##### 2. Database Test
On Database-Test folder, there are file.sql which each of
it contains query and the example of result (.png)
##### 3. Rest Test
- run database migration
```bash
make migrate
```
this command will automatically create column for the table.
on Makefile, please change the table name as you need 
(my current table name are koala)
```
migrate -source file://Rest/db/migrations \
 			-database postgres://root:secret@localhost/koala?sslmode=disable up
```
- run the server
```
go run main.go && Koala (this project folder)
```
##### 4. All of route list in main.go