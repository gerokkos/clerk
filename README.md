<br />
<p align="center">
  <a href="https://github.com/gerokkos/clerk">
    <img src="https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2Fdb9d7081-fbdc-483e-af81-c09036f7ec36%2Ffavicon.jpg?table=block&id=64fe587d-3296-4e1c-b460-6b78169a49bf&width=250&userId=435c4628-e44f-43f5-964f-caecf3a3ed84&cache=v2" alt="Logo" width="80" height="80">
  </a>
  <h3 align="center">Clerk Randomuser API</h3>
</p>

# Table of Contents

* [Getting Started](#getting-started)
* [Project Structure](#project-structure)
* [Integration Guide ](#integration-guid)
* [Use the API](#use-the-api)
* [Clone the Application](#clone-the-application)
* [Run with Docker](#run-with-docker)
* [Run the Application Locally](#run-the-application-locally)
* [Clean tests cache if cached](#clean-tests-cache-if-cached)
* [Contributing](#contributing)





# Project Structure


    ├── Clerk                    
    │   ├── api             # Functionality/features
    │   ├── tests           # Testing endpoints  
    │   └── main            # Run the application
     



## Integration Guide 
The database used is PostgreSQL

### Use the API

| Endpoint         |                              |   
| -------------    | -----------------------------|
| /populate        |                              |
| /clerks          | ?limit=<>                    |
| /clerks          | ?starting_after=<>           |
| /clerks          | ?ending_before=<>            |
| /clerks          | ?starting_after=<>&limit=<>  |
| /clerks          | ?starting_after=<>&limit=<>  |
| /clerks          | ?email=""                    |




### Clone the Application

``
git clone git@github.com:gerokkos/clerk.git
``


### Run with Docker

Start up the application services by running:

``
docker-compose up
``

You call use the API from http://localhost:8080


To stop the services run: 

``
docker-compose down
``


Run the tests with Docker:

``
docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
``



To use the pgAdmin:

http://localhost:5050
i. Choose Create then Server
ii. Fill in any name that you want.
iii. Click on connection tab.

Run:
``
docker container ls
``

Copy the ID of the clerk-db-postgres and use it here:

``
docker inspect <container_id> | grep IPAddress
``

The IPAddress, is the host name in pgAdmin and username-password the ones in the .env



### Run the Application Locally

``
go run main.go
``

### Clean tests cache if cached

``
go clean -testcache
``


## Contributing

1. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request
