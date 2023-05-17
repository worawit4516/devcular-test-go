# Devcular-test-go

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#notes">Notes</a></li>
  </ol>
</details>

## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.

- Golang : https://go.dev/dl/
- Postman : https://www.postman.com/downloads/
- Postman Collections : https://shorturl.at/fiVW0

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/worawit4516/devcular-test-go.git
   ```
2. Go to file
   ```sh
   cd devcular-test-go
   ```
3. Install GO package if it necessary
   ```sh
   go get github.com/gin-gonic/gin
   go get github.com/jinzhu/gorm
   go get gorm.io/driver/mysql
   go get github.com/joho/godotenv
   ```
4. Run go command to start server or use go built files
   ```sh
   go run main.go
   ```
   or
   <br>
   ```sh
   main.exe
   ```

   
## Usage

1. Open Postman Collections name TodoLists on this links : https://shorturl.at/fiVW0
2. Use API [GET] Check Health to see status of server. You should see this response                           
   ```sh
   "Backend is Good!!" 
   ```
4. Test all API in collections.
    
### Notes
Here's why:
* The ENV file is used to store connection strings for databases, but it doesn't encrypt them for easier testing.
* The database is an Azure Database for MySQL Flexible Server run on Azure Cloud with an open firewall for any IP address.
