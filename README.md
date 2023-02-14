# Database Migrations

## What is database migration:
Database migration is a way to manage database state. By migrating database up or down, we can control database state and tables.
We can see current migration version in table `schema_migration` in each database.

## Install golang migration:

### For Ubuntu:
1. Open Ubuntu terminal
2. Run the following command in terminal:  
`curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash`
3. Run `sudo apt-get update`
4. You're now all-set to setup golang migrate. Run the following command to install migrate:
`sudo apt-get install migrate`

## Create database migration file:
When creating a new migration version, 2 files will be created having postfix `.up.sql` and `.down.sql`. The first file is used to migrate database to an upper version, while the second is to move the database to a previous state.

To create a new migration version, navigate to service directory and run the following command in terminal:

`make migrate name={given_name}`

where `given_name` is the name to describe the version. Two up and down files will be created with corresponding sequence number and name.

## Migrate database up:
To migrate database up, navigate to service directory and run the following command in terminal:

`make migrate-up`

Database state will be updated to the latest version.

## Migrate database down:
To migrate database down, navigate to service directory and run the following command in terminal:

`make migrate-down`

Scripts in all down files will be run, and database state will be moved to the first one.

# Run application:

## Start all services:
To start all service at once, at project directory, run command `make all_start` to start all services.

## Start specific services:
To start some specific services, navigate to that service directory and run command `make start` to start that service.

# References:

1. https://viblo.asia/p/dung-golang-microservice-boilerplate-theo-clean-architecture-1Je5EzG0KnL
2. https://dev.to/koddr/let-s-write-config-for-your-golang-web-app-on-right-way-yaml-5ggp
3. https://gorm.io/docs/index.html
4. https://dev.to/techschoolguru/load-config-from-file-environment-variables-in-golang-with-viper-2j2d