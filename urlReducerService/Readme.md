# URL Reducer Service

## Description

This service provides two different methods for the same endpoint (`/url`):

1. **POST**: Use this method to retrieve data for a short URL provided in the body of the request.
   
    Example:
    ```bash
    curl -X POST -d 'shorturl' http://localhost:8081/url
    ```

2. **GET**: Use this method to fetch the short URL from the service if the corresponding long URL is available in the map.

    Example:
    ```bash
    curl -X GET 'http://localhost:8081/url?url=longurl'
    ```
