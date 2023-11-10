# go-subscription-service
Go concurrency practical project. Implemented for 'Working with Concurrency in Go (Golang)' course.

Application provides possibility of basic authorization, new user registration, dummy subscription representation and functional for change subscription for user. 
Main feature of application is send email notification (in separate go routine) for user notification about incorrect login, account activation after registration and about subscription change.
Also application send dummy attachment with notification in case of user subscription change

## To run application:
Install Go on your computer https://go.dev/doc/install

Install Make https://www.gnu.org/software/make/

Install Docker Desktop https://www.docker.com/

Clone git repository:
```
git clone https://github.com/viktor-mazepa/go-subscription-service.git go-subscription-service
```
Navigate to repository folder:
```
cd go-subscription-service folder
```
Run docker compose to create all needful virtual environments:
```
docker-compose up -d
```
Connect to Postgresql database and execute SQL queries from ```./scripts/init.sql``` to create all needful tables, add admin user and test data
```
{
  "label": "go-subscription-service",
  "host": "localhost",
  "user": "postgres",
  "port": 5432,
  "ssl": false,
  "database": "concurrency",
  "password": "password"
}
```
Run application:
```
make build
```
```
make start
```

Open browser and navigate to ```http://localhost:8080/```

Execute test login ```admin@example.com/verysecret```

Navigate to ```Plans``` tab and subscribe any plan

Open email server ```http://localhost:8025/``` and check mailbox - two email should be present (one with simple invoice and the second with plan manual attachment in PDF format)
