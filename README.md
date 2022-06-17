# RATE-LIMITER


command to run:
go run app.go -threshold <num >= 0> -ttl <num >= 0> -port < num >= 1023 > 49151
example: 
go run app.go -threshold 20 -ttl 20000 -port 3000

 if argument will not be supplied default values will kick: 
 threshold = 10
 ttl = 60000 (1 minute)
 port = 8081 

 explanations: 
 we have 
 services package: logic and model layer
 1. routine that runs every ttl * time in milliseconds and clear the Url map.
    when reset time arrives, i just create a new map, garbage collector will clean the old one as you know..
 2. url-service is where the insertion, synchronization and cache managment happens
    the chache is a struct with
    - map<string, int>: to cache the urls and count,
    - RWMutex to insure synchronization
    - ttl: for reset
    - threshold for maximun attempts per url.
handler package: validation and controller level 
 3. handlers package is the "controller" upcoming request will trigger this function 
    this layer is also validating the request's body contentType and dto's.
    if there's were more validations and dto, i would separate them but for this assignment i flet like putting them together is ok... 

