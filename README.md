# How to create Golang CRUD WebAPI Azure PostgreSQL Microservice and deploy to Docker Desktop and Kubernetes (in your local laptop)

The source code is available in this github: 

https://github.com/luiscoco/Golang-sample19-Azure-PostgreSQL-WebAPI-CRUD-Microservice

This example was develop for protocol **HTTP** if you would like to use also protocol **HTTPS** please refer to the repo:

https://github.com/luiscoco/MicroServices_dotNET8_CRUD_WebAPI-Azure-MySQL

## 1. Prerequisite

### 1.1. Create Azure PostgreSQL instance

We navigate to **Create a resource** and select **Databases**

![image](https://github.com/luiscoco/MicroServices_dotNET8_CRUD_WebAPI-Azure-PostgreSQL/assets/32194879/34b76cbb-5e03-4959-9eee-3e5d822b09e5)

We select Azure Database for **PostgreSQL**

![image](https://github.com/luiscoco/MicroServices_dotNET8_CRUD_WebAPI-Azure-PostgreSQL/assets/32194879/dd6b48c7-0195-4cbf-90d8-0bd5863a3932)

We input the required data to create a new **Flexible Server**

**Server Name**: postgresqlserver1974

![image](https://github.com/luiscoco/MicroServices_dotNET8_CRUD_WebAPI-Azure-PostgreSQL/assets/32194879/76d5eef3-e234-4b40-9236-bdbb6e6e29b9)

We **configure the server** (compute and storage) abd press the **Save** button 

![image](https://github.com/luiscoco/MicroServices_dotNET8_CRUD_WebAPI-Azure-PostgreSQL/assets/32194879/6e2434a0-70d0-4821-844b-9ceddf0011e4)

![image](https://github.com/luiscoco/MicroServices_dotNET8_CRUD_WebAPI-Azure-PostgreSQL/assets/32194879/c9677b96-fee6-4020-8938-4cf18ddd8730)

![image](https://github.com/luiscoco/MicroServices_dotNET8_CRUD_WebAPI-Azure-PostgreSQL/assets/32194879/299ae028-9da3-4d99-8466-33fd9b485c62)

We continue configuring the **Authentication** data

![image](https://github.com/luiscoco/MicroServices_dotNET8_CRUD_WebAPI-Azure-PostgreSQL/assets/32194879/e9655f2e-04c5-4a68-9475-f0fc7477da28)

We navigate to the **Networking** tab and we add our local laptop IP address as a FireWall rule

![image](https://github.com/luiscoco/MicroServices_dotNET8_CRUD_WebAPI-Azure-PostgreSQL/assets/32194879/08ef7fc4-adbc-4a78-925d-c6363cc0b2b1)

We press the **Review and create** button 

![image](https://github.com/luiscoco/MicroServices_dotNET8_CRUD_WebAPI-Azure-PostgreSQL/assets/32194879/7923bad1-364e-4932-8507-b1acc3c1335d)

### 1.2. Install and Run PostgreSQL and pgAdmin

We first **download and install PostgreSQL** from this URL: https://www.postgresql.org/download/windows/

![image](https://github.com/luiscoco/MicroServices_dotNET8_CRUD_WebAPI-Azure-PostgreSQL/assets/32194879/3a50ac64-83ab-4aa3-912f-a06faea2ec69)

https://www.enterprisedb.com/downloads/postgres-postgresql-downloads

![image](https://github.com/luiscoco/MicroServices_dotNET8_CRUD_WebAPI-Azure-PostgreSQL/assets/32194879/2e3af074-77b9-48da-8401-2156176cd409)

We open a command prompt window as Administrator and **connect to the Azure PostgreSQL** instance with this command

```
psql -h postgresqlserver1974.postgres.database.azure.com -d postgres -U adminpostgresql
```

We **create a database** with this command

```
CREATE DATABASE postgresqldb
    WITH ENCODING 'UTF8'
    LC_COLLATE='en_US.utf8'
    LC_CTYPE='en_US.utf8'
    TEMPLATE=template0;
```
![image](https://github.com/luiscoco/MicroServices_dotNET8_CRUD_WebAPI-Azure-PostgreSQL/assets/32194879/6a68ffee-6182-468b-bdce-190f23830501)

**IMPORTANT NOTE**: 

The **LC_COLLATE** parameter determines the sorting order of strings in the database, such as how names and titles are sorted in queries. This setting is crucial for ensuring that data is sorted correctly according to the local conventions of the database's intended audience. Ensure that you're using the correct syntax when specifying the LC_COLLATE setting. 

Use **Template0**: When creating a new database with specific **LC_COLLATE** and **LC_CTYPE** settings, it's recommended to use **TEMPLATE=template0**, as shown in the example above, because template0 is guaranteed to have the default settings, ensuring that the new database will inherit the specified LC_COLLATE and LC_CTYPE settings without issues.

We can now access to **Azure PostgreSQL** from **pgAdmin 4** setting the hostname, username and password

We download and install **pgAdmin 4**

https://www.pgadmin.org/download/pgadmin-4-windows/

![image](https://github.com/luiscoco/MicroServices_dotNET8_CRUD_WebAPI-Azure-PostgreSQL/assets/32194879/3c0d33f4-bfa8-4565-a78a-d0a9faebeee5)

We run **pgAdmin 4** and we create a new connection

![image](https://github.com/luiscoco/MicroServices_dotNET8_CRUD_WebAPI-Azure-PostgreSQL/assets/32194879/870af1a5-c770-4862-9902-9b78ba55eadb)

We input the new connection data and we press the **Save** button to connect to the Azure PostgreSQL instance

![image](https://github.com/luiscoco/MicroServices_dotNET8_CRUD_WebAPI-Azure-PostgreSQL/assets/32194879/47ff6864-3aa8-4c92-86b8-2a620b09bb1a)

![image](https://github.com/luiscoco/MicroServices_dotNET8_CRUD_WebAPI-Azure-PostgreSQL/assets/32194879/a1c4914b-32ff-42b4-a034-d415bab3f3a1)

We verify the new Server in the list

![image](https://github.com/luiscoco/MicroServices_dotNET8_CRUD_WebAPI-Azure-PostgreSQL/assets/32194879/fc2ca09c-8390-43ae-81fe-3cf5515e865d)

If we click on the **databsename**, after we click on the **Schemas** and the in **Tables**

![image](https://github.com/luiscoco/MicroServices_dotNET8_CRUD_WebAPI-Azure-PostgreSQL/assets/32194879/5e308d0e-7159-47a8-9771-24796900accc)

For running a SQL query we click on the **Query Tool** button

![image](https://github.com/luiscoco/MicroServices_dotNET8_CRUD_WebAPI-Azure-PostgreSQL/assets/32194879/0451248e-fea7-4c05-873c-7c512a1087a4)

### 2. How to Create a Golang WebAPI CRUD Microservice

Run VSCode and create the project folder structure, we create three folder: **controllers**, **models** and **util**

![image](https://github.com/luiscoco/Golang-sample19-Azure-PostgreSQL-WebAPI-CRUD-Microservice/assets/32194879/187b57af-2c31-47b3-a99f-26f7e2642318)

We create then **main.go** file

```go
package main

import (
    "log"
    "net/http"
    "go_application/controllers"
    "go_application/util"
    "github.com/gorilla/mux"
)

func main() {
    config, err := util.LoadConfig("config.json")
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }

    util.InitDB(config.DatabaseURL)

    r := mux.NewRouter()
    r.HandleFunc("/items", controllers.GetItems).Methods("GET")

    log.Println("Starting server on :8081")
    if err := http.ListenAndServe(":8081", r); err != nil {
        log.Fatalf("Error starting server: %s\n", err)
    }
}
```

This code snippet is a Go application that sets up a web server. Let's break it down into its main components:

Package Declaration:

go
Copy code
package main
This line declares the package name, which in this case is main. In Go, the main package is special as it defines a standalone executable program, not a library.

Imports:

go
Copy code
import (
    "log"
    "net/http"
    "go_application/controllers"
    "go_application/util"
    "github.com/gorilla/mux"
)
This block imports various packages needed for the application:

log for logging,
net/http for HTTP server and client,
go_application/controllers and go_application/util likely refer to custom packages within the application for controllers and utility functions,
github.com/gorilla/mux is an external package for handling HTTP request routing.
Main Function:

go
Copy code
func main() {
    ...
}
The main function is the entry point of the program. When the program runs, this function is automatically executed.

Configuration Loading:

go
Copy code
config, err := util.LoadConfig("config.json")
if err != nil {
    log.Fatalf("Failed to load configuration: %v", err)
}
Here, the program tries to load configuration settings from a file named config.json using a function LoadConfig from the util package. If an error occurs during loading, the program logs the error and terminates.

Database Initialization:

go
Copy code
util.InitDB(config.DatabaseURL)
This line initializes the database using a URL specified in the configuration. It uses the InitDB function from the util package.

Router Setup with Gorilla Mux:

go
Copy code
r := mux.NewRouter()
r.HandleFunc("/items", controllers.GetItems).Methods("GET")
mux.NewRouter() creates a new router using the Gorilla Mux package, which is more powerful and flexible than the default mux provided by net/http.
r.HandleFunc("/items", controllers.GetItems).Methods("GET") sets up a route. It means that for the HTTP path /items with method GET, the function GetItems from the controllers package will handle the request.
Starting the HTTP Server:

go
Copy code
log.Println("Starting server on :8081")
if err := http.ListenAndServe(":8081", r); err != nil {
    log.Fatalf("Error starting server: %s\n", err)
}
The server is started on port 8081 using http.ListenAndServe.
It logs a message indicating that the server is starting.
If there's an error in starting the server, it logs the error and terminates.
In summary, this Go application loads its configuration, initializes a database, sets up HTTP routing using Gorilla Mux, and starts an HTTP server to handle requests for the /items endpoint with a GET method.



