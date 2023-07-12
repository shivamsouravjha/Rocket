## How to get started?

* docker-compose up (to start rabbitmq and mysql image)

* In a new terminal set up user table
    *  docker exec -it mysql bash (to connec to mysql database)
    * mysql -uroot -p (to login in as root user password is newpassword)
    * use productio (our database for the project)
    * Create table for user and product (code give to generate table)
    * run go run main.go to run the backend server
