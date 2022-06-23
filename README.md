# STRIPE PAY SERVICE

Server side implementation of Stripe payment gateway.

### RUN THE PROJECT

Make sure you have docker installed.

#### ``LOCALLY``
#### Clone project

```bash
git clone https://github.com/swagftw/stripe_pay_service
```

##### Spin up server

```bash
docker-compose up
```

#### ``DEPLOYED SERVER``

Go to the postman environments and add variable `url` as `https://stripepayservice-production.up.railway.app/api/v1`


### REQUESTING
##### Download postman collections

Setup an environment in postman add the following variable:

If you are running locally

`url`=`http://localhost:8080/api/v1`

If you are running on the server

`url`=`https://stripepayservice-production.up.railway.app/api/v1`

[Download postman collection](https://www.getpostman.com/collections/e63b9ec893946fcc55df)

[Download postman test collection](https://www.getpostman.com/collections/5afac639f2c0f2d4b593)

Choose the environment just created and check requests in above collections.

### RUN THE TESTS

Make sure you have go installed


##### Get dependencies
cd to cloned project

```bash
go mod download && go mod vendor
```

##### Get Stripe mock server

```bash
docker run -p 12111:12111 stripe/stripe-mock
```

##### Run tests

```bash
go test -v ./...
```

### MORE ABOUT PROJECT
`Project architecture`

```
Project is based on hexagonal architecture. It consists of three primary layers:
 - business logic layer
 - data layer
 - transport layer
```
More about [architecture](https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749)

![image](https://miro.medium.com/max/1400/1*NfFzI7Z-E3ypn8ahESbDzw.png "hey")

`Project structure`

```
- /cmd              // contains all the executables
  |- /api           // contains the API server startpoint
  
- /pkg              // contains all the packages, which is core business logic
  |- /api           // initial API server, acts as dependency injection container for starting api server.
  |- /payments      // payments core logic package
     |- /repository // payments database repository package
     |- payments.go // holds implementation of payments service interface
     |- service.go  // contains repository interface and db models
       
- /transaction      // contains the global transaction interface, that can be implemented by multiple dbs
  |- /postgres      // contains postgres implementation of transaction interface
  
- /transport        // contains the all sorts of transports, currently over http may contain rpc, grpc, etc
  |- /payments      // contains the http handlers for payments service
    
- /types            // contains all the service interfaces & types, required for service and it sits on top of the project heirarchy
 
- /utl              // contains all the utility functions
  |- /config        // config utility functions
  |- /constant      // contains constants used over project
  |- /fault         // fault is custom error type used over project to throw errors 
  |- /logger        // custom logger implementation over Uber's zap logger
  |- /migration     // database migration utility functions
  |- /mock          // mocks for different services
  |- /server        // server utility functions and custom error handler, validators, middlewares     
  |- /storage       // database utitilies
  |- /stripeclient  // custom implementation over stripe go sdk for abstracting stripe api
```

`API Testing`

```
To test API run postman test collection provided above.
```


`Future improvements`

```
- Writing more test cases
- Writing API test cases in project itself.
- Remove env file from project, kept file for now in remote repo just for dev purposes.
```

