## How to get started?

* ```docker-compose up``` (to start rabbitmq and mysql image)

* In a new terminal set up the user table
    *  ```docker exec -it mysql bash``` (to connect to mysql database)
    * ```mysql -uroot -p``` (to login in as root user password is newpassword)
    * ```use productio``` (our database for the project)
    * Create table for [user](https://github.com/shivamsouravjha/Rocket/blob/main/helper/sql/usertable.sql) and [product](https://github.com/shivamsouravjha/Rocket/blob/main/helper/sql/productTable.sql) (code give to generate table)
* run ```go run main.go``` to run the backend server
