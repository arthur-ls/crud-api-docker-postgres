# CRUD API with Docker Postgresql

The objective of this project is to get data from an external api and insert it in a postgresql table created in a Docker image besides doing CRUD actions.

## Used API to get Data

The used API is [Sport Data API](https://sportdataapi.com/) that is a Free API that provides data of all kinds of sports and the URL used get all soccer championships in the world. To get data of this API it is necessary to get the api-key provided by them after you sign up in the Sport Data API website.
If you want to get Data from a different API it is possible to change the URL and the necessary keys to make requests.

## Using Docker

Before get Docker started [Colima](https://smallsharpsoftwaretools.com/tutorials/use-colima-to-run-docker-containers-on-macos/) was used to prepare the environment to get the [Docker](https://www.docker.com/) up.

So after installing Colima and Docker and creating docker-compose.yml the terminal was opened:

'colima start'

After it successfully ran:

'docker-compose up'

Then we will have our Postgresql ready to accept connections. Furthermore it is possible to connect the database in a database tool as DBeaver to make queries and see data added in the table.

## Using DBeaver

DBeaver is the database tool used to make queries in the created table and see the data in Postgresql but it is possible to use different kind of tools as Datagrip, etc.

To configure it is necessary to create a new database connection selecting Postgresql as the database to be used and after that add the used host, port, database, username and password. Lastly test the connection and finish the database configuration. 

## Running Project

After Docker is up and the dbParams.env is correctly configured with host, port, user, password and database name we can run the project by:

'go run main.go'

Done that we will see the lines being added to the table by the terminal output. Then we can check if all the data is in the Postgresql table by making a query in a database tool such as DBeaver:

'SELECT * FROM competition'

## CRUD actions in Postman

Postman was used to make CRUD actions. For this we have to use the following URLS to each action:

'ACTION: POST
URL: http://localhost:3000/create

ACTION: GET
URL: http://localhost:3000/get
URL: http://localhost:3000/get/{id}

ACTION: PUT
URL: http://localhost:3000/update/{id}

ACTION: DELETE
URL: http://localhost:3000/delete/{id}'

Then we get the response with a successful message and the data requested.