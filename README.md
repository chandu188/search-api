# Search API
##### Assumptions:
    • Authentication: Implemented authentication using a middleware, which is BaiscAuth middleware and currently, the user accounts are hardcoded with a single user (username:password=chandu:password). It can be easily extended to write JWT middleware to support other types of authentication/authorization mechanisms
    • The elastic search used in the docker-compose is a single node cluster.
    • 

### Build
``` 
To build the application: make build 
To run tests: make test
To run the application using docker: make docker-compose
```

### Usage
    1) To search products filtered by brand name Nike and 
            http://localhost:8080/api/v1/products?filter=brand:Nike
    2) To query for `black shoes` use
            http://localhost:8080/api/v1/products?q=black shoes
    3) To sort the results by a key (default is desc)
            http://localhost:8080/api/v1/products?filter=brand:Nike&sort_by=price:asc
            http://localhost:8080/api/v1/products?filter=brand:Nike&sort_by=price (desc)

    4) To limit the results by 5 items 
            http://localhost:8080/api/v1/products?filter=brand:Nike&limit=5
        and to retreive the next batch of results use `offset` key
            http://localhost:8080/api/v1/products?filter=brand:Nike&limit=5&offset=6

    5) http://localhost:8080/index gives a simple UI to add and retrieve the products (it doesn't give any feedback if the add to product fails).
            