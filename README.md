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

## 2. How to Create a Golang WebAPI CRUD Microservice

Run VSCode and create the project folder structure, we create three folder: **controllers**, **models** and **util**

![image](https://github.com/luiscoco/Golang-sample19-Azure-PostgreSQL-WebAPI-CRUD-Microservice/assets/32194879/187b57af-2c31-47b3-a99f-26f7e2642318)

### 2.1. main.go file

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

**Package Declaration**:

```go
package main
```

This line declares the **package name**, which in this case is **main** 

In Go, the main package is special as it defines a **standalone executable program**, not a library

**Imports**:

```go
import (
    "log"
    "net/http"
    "go_application/controllers"
    "go_application/util"
    "github.com/gorilla/mux"
)
```

This block **imports packages** needed for the application:

**log** for logging,

**net/http** for HTTP server and client,

**go_application/controllers** and **go_application/util** likely refer to custom packages within the application for controllers and utility functions,

**github.com/gorilla/mux** is an external package for handling HTTP request routing.

**Main Function:**

```go
func main() {
    ...
}
```

The main function is the entry point of the program. When the program runs, this function is automatically executed.

**Configuration Loading**:

```go
config, err := util.LoadConfig("config.json")
if err != nil {
    log.Fatalf("Failed to load configuration: %v", err)
}
```

Here, the program tries to load configuration settings from a file named **config.json** using a function **LoadConfig** from the util package. 

If an error occurs during loading, the program logs the error and terminates.

**Database Initialization**:

```go
util.InitDB(config.DatabaseURL)
```

This line initializes the database using a URL specified in the configuration. It uses the InitDB function from the util package.

**Router Setup with Gorilla Mux**:

```go
r := mux.NewRouter()
r.HandleFunc("/items", controllers.GetItems).Methods("GET")
```

**mux.NewRouter()** creates a new router using the Gorilla Mux package, which is more powerful and flexible than the default mux provided by net/http.

**r.HandleFunc("/items", controllers.GetItems).Methods("GET")** sets up a route. 

It means that for the HTTP path /items with method GET, the function GetItems from the controllers package will handle the request.

**Starting the HTTP Server**:

```go
log.Println("Starting server on :8081")
if err := http.ListenAndServe(":8081", r); err != nil {
    log.Fatalf("Error starting server: %s\n", err)
}
```

The server is started on port 8081 using http.ListenAndServe.

It logs a message indicating that the server is starting.

If there's an error in starting the server, it logs the error and terminates.

In summary, this Go application loads its configuration, initializes a database, sets up HTTP routing using Gorilla Mux, and starts an HTTP server to handle requests for the /items endpoint with a GET method.

### 2.2. config.go and config.json files

```go
package util

import (
    "encoding/json"
    "os"
)

type Config struct {
    DatabaseURL string `json:"DatabaseURL"`
}

func LoadConfig(path string) (*Config, error) {
    configFile, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer configFile.Close()

    var config Config
    jsonParser := json.NewDecoder(configFile)
    err = jsonParser.Decode(&config)
    if err != nil {
        return nil, err
    }

    return &config, nil
}
```

This code snippet is a Go **package** named **util** and it primarily deals with loading configuration settings from a JSON file. Let's break it down:

**Package Declaration**:

```go
package util
```

This line declares the package name as util. This package is meant to provide utility functions, which can be used across the application.

**Imports**:

```go
Copy code
import (
    "encoding/json"
    "os"
)
```

This block **imports two packages**:

**encoding/json** for JSON encoding and decoding,

**os** for interacting with the operating system, like opening files.

**Config Struct**:

```go
type Config struct {
    DatabaseURL string `json:"DatabaseURL"`
}
```

This defines a struct named Config with a single field DatabaseURL. 

The json:"DatabaseURL" part is a field tag used by the json package to map this field to the DatabaseURL key in a JSON object.

**LoadConfig Function**:

```go
func LoadConfig(path string) (*Config, error) {
    ...
}
```

**LoadConfig** is a function that takes a file path as a string and returns a pointer to a **Config** struct and an error.

It is designed to load configuration data from a JSON file whose path is provided as an argument.

**Opening the Configuration File**:

```go
configFile, err := os.Open(path)
if err != nil {
    return nil, err
}
defer configFile.Close()
```

The function attempts to open the file at the given path using **os.Open**

If there is an error in opening the file (e.g., file does not exist), it returns nil and the error

**defer configFile.Close()** schedules the closing of the file when the function exits 

This is important for resource management and to prevent file descriptor leaks

**Decoding the JSON File**:

```go
var config Config
jsonParser := json.NewDecoder(configFile)
err = jsonParser.Decode(&config)
if err != nil {
    return nil, err
}
```

It creates a variable config of type Config.

A JSON decoder is created for the opened file.

The decoder then tries to decode the JSON content of the file into the config variable.

If there's an error during decoding (e.g., if the JSON structure doesn't match the Config struct), it returns nil and the error.

**Returning the Config Data**:

```go
return &config, nil
```

Finally, if everything goes well, the function returns a pointer to the config struct filled with data from the JSON file, and a nil error

In summary, this code provides a utility function LoadConfig to load and parse configuration data from a JSON file into a Config struct

It handles file operations and JSON parsing, returning the parsed configuration and any errors encountered in the process

We also input the database connection string in the **config.json** file

```json
{
    "DatabaseURL": "postgres://adminpostgresql:Luiscoco123456@postgresqlserver1974.postgres.database.azure.com:5432/postgresqldb?sslmode=require"
}
```

### 2.3. database.go file

```go
package util

import (
    "database/sql"
    _ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(dataSourceName string) {
    var err error
    db, err = sql.Open("postgres", dataSourceName)
    if err != nil {
        panic(err)
    }
    if err = db.Ping(); err != nil {
        panic(err)
    }
}

func GetDB() *sql.DB {
    return db
}
```

This code snippet is part of a Go package named **util** that is designed to **initialize** and provide access to a **PostgreSQL** database. Let's break down its components:

**Package Declaration**:

```go
package util
```

This line declares the package name as **util**. 

This package is intended to provide utility functions, including those for database interactions

**Imports**:

```go
Copy code
import (
    "database/sql"
    _ "github.com/lib/pq"
)
```

**database/sql** is a Go standard library package that provides a generic interface around SQL (or SQL-like) databases

**_ "github.com/lib/pq"** is an import for the PostgreSQL driver. 

The underscore _ before the import path is used for importing a package solely for its side-effects (in this case, registering the driver with the database/sql package), without directly using any functions or variables from the package in the code.

**Global Variable**:

```go
var db *sql.DB
```

This line declares a global variable db of type ***sql.DB**.

This pointer to an sql.DB will represent the database connection pool.

**InitDB Function**:

```go
func InitDB(dataSourceName string) {
    var err error
    db, err = sql.Open("postgres", dataSourceName)
    if err != nil {
        panic(err)
    }
    if err = db.Ping(); err != nil {
        panic(err)
    }
}
```

**InitDB** is a function that initializes the database connection.

It takes a dataSourceName string which contains the connection information for the PostgreSQL database.

**sql.Open("postgres", dataSourceName)** attempts to open a database connection with the given data source name. 

It doesn’t immediately establish a connection but prepares the database connection for use.

If there is an error during Open, the function panics (abruptly terminates the program with the error)

**db.Ping()** checks if the database is accessible and also establishes a connection to the database. If it fails, the function again panics

This function sets the global db variable to be used across the application

**GetDB Function**:

```go
func GetDB() *sql.DB {
    return db
}
```

**GetDB** is a simple function that returns the global db variable

This allows other parts of the application to access the initialized database connection

In summary, this code provides functionality to initialize and access a PostgreSQL database within a Go application

It uses the **database/sql** package and the PostgreSQL driver (**lib/pq**) to manage database connections

The **InitDB** function sets up the connection, and GetDB provides access to this connection

The use of panic for error handling suggests that any error in connecting to the database is considered fatal for the application

### 2.4. items_controller.go file

```go
package controllers

import (
    "encoding/json"
    "net/http"
    "go_application/models"
    "go_application/util"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
    db := util.GetDB()
    items, err := models.GetAllItems(db)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(items)
}
```

This code snippet is part of a Go package named controllers, which appears to be responsible for handling HTTP requests in a web application

The specific function in this snippet is GetItems, which is designed to handle requests for retrieving items. Let's break it down:

**Package Declaration**:

```go
package controllers
```

This line declares the package name as controllers

This package likely contains functions that act as controllers in the MVC (Model-View-Controller) design pattern.

**Imports**:

```go
import (
    "encoding/json"
    "net/http"
    "go_application/models"
    "go_application/util"
)
```

**encoding/json** is used for encoding and decoding JSON

**net/http** is used to handle HTTP requests and responses.

**go_application/models** is likely a custom package in the application, probably containing data models and business logic.

**go_application/util** is another custom package, possibly containing utility functions like database connection handling (as suggested by its usage in the code).

**GetItems Function**:

```go
func GetItems(w http.ResponseWriter, r *http.Request) {
    ...
}
```

**GetItem** is an HTTP handler function. It takes two parameters: w (an http.ResponseWriter) for writing the HTTP response, and r (an *http.Request) representing the incoming HTTP request.

This function is responsible for processing HTTP requests that are intended to retrieve items.

**Database Connection and Retrieval of Items**:

```go
db := util.GetDB()
items, err := models.GetAllItems(db)
if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}
```

**db := util.GetDB()** calls the GetDB function from the util package to get a database connection

**items, err := models.GetAllItems(db)** calls a function from the models package, which likely queries the database to retrieve all items. It returns the items and any error that occurred

If there's an error (**err != nil**), the function sends an HTTP error response with status code 500 (Internal Server Error) and returns early from the function

**Encoding and Sending the Response**:

```go
json.NewEncoder(w).Encode(items)
```

If no error occurs, the items are encoded into JSON format using **json.NewEncoder(w).Encode(items)**

The JSON-encoded items are written to the **http.ResponseWriter** (w), which sends the data back to the client making the request

In summary, this **GetItems** function is an HTTP handler that interacts with the database to retrieve items and returns them to the client in JSON format

If an error occurs during database interaction, it responds with an HTTP error

This function is part of the application's controller layer, handling the logic for HTTP requests related to items

### 2.5. item.go

```go
package models

import (
    "database/sql"
)

type Item struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func GetAllItems(db *sql.DB) ([]Item, error) {
    items := []Item{}
    rows, err := db.Query("SELECT id, name FROM items")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var i Item
        if err := rows.Scan(&i.ID, &i.Name); err != nil {
            return nil, err
        }
        items = append(items, i)
    }
    return items, nil
}
```

This code is a part of the models package in a Go application and defines the functionality for interacting with a database to retrieve a list of Item objects. Let's break down the code:

**Package Declaration**:

```go
package models
```

This line declares that the code belongs to the models package

In the context of an MVC (Model-View-Controller) framework, models typically handle data and business logic

**Imports**:

```go
import (
    "database/sql"
)
```

This imports the database/sql package, which is a Go standard library package for interacting with SQL databases

**Item Struct**:

```go
type Item struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
```

This defines a struct named Item, which represents an item in the database

It has two fields: ID (an integer) and Name (a string)

The struct tags (e.g., json:"id") indicate how these fields should be encoded and decoded when converting to and from JSON

This is useful when sending/receiving data in JSON format over HTTP.

**GetAllItems Function**:

```go
func GetAllItems(db *sql.DB) ([]Item, error) {
    ...
}
```

**GetAllItems** is a function that takes a pointer to an sql.DB (representing the database connection) and returns a slice of Item structs and an error

This function is responsible for querying the database and returning all items

**Querying the Database**:

```go
Copy code
rows, err := db.Query("SELECT id, name FROM items")
if err != nil {
    return nil, err
}
defer rows.Close()
```

The database is queried for all rows in the items table, selecting the id and name columns

If an error occurs during the query, it returns nil and the error

**defer rows.Close()** ensures that the result set (rows) is closed when the function exits

This is important for resource management and avoiding memory leaks

**Iterating Over the Result Set**:

```go
for rows.Next() {
    var i Item
    if err := rows.Scan(&i.ID, &i.Name); err != nil {
        return nil, err
    }
    items = append(items, i)
}
```

The function iterates over each row in the result set

For each row, it creates an Item struct (i) and uses **rows.Scan** to copy the columns from the current row into the struct's fields (i.ID and i.Name)

If there's an error in scanning, it returns nil and the error

The Item struct is then appended to the items slice

**Returning the Items**:

```go
return items, nil
```
After processing all rows, the function returns the slice of Item structs

In summary, the **GetAllItems** function in the models package is designed to query a SQL database for all entries in the items table, and it returns these as a slice of Item structs

This function encapsulates the logic for data retrieval, abstracting the database interaction away from other parts of the application

### 2.6. How to load the dependencies










