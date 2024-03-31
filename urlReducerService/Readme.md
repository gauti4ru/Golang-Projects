#URL Reducer Service 

##Description
1- There are two different methods available for one service 
    end point(/url)
2- Post is to get the data for a short url which is getting passed in body of the request.
   `curl -X POST -d 'shorturl' http://localhost:8081/url`
3- Get is to fetch the short url from the service if,that short url is available in the map.
    `curl -X  GET 'http://localhost:8081/url?url=longurl'`