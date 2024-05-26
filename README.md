# go_fintech
A simple project building backend CRUD app using Golang

Using example from [Duomly](https://dev.to/duomly/learn-golang-by-building-a-fintech-banking-app-lesson2-login-and-rest-api-1hh2)

## Development

Before running the app, please make sure that you have the following installed:
1. Docker
2. Make utility

### Starting the database
1. Create two secret files for the database. Refer to the example files in the secrets folder. The files are:
    - `db_user.txt`
    - `db_password.txt`
2. Run`make docker-compose` which will download the postgres image and start the database

### Running the app
1. Compile the app by typing `make build`
2. Run `DB_USER={YOUR_DB_USER} DB_PASSWORD={YOUR_DB_PASSWORD} make run`. Replace the database username and password as what is set in the previous section.
