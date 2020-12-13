# Timezone application

## How to use
This a simple application written in plain go which allows you to get current time on a specific timezone:

```
   docker run -p 8080:8080 -d mr14mr/timezone-app:latest
   curl http://localhost:8080/Europe/Moscow
```

You can specify any valid timezone listed [here](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). Input is **case sensitive**

## Docker

Docker image uses two stages in order to get resulting image with the smaller size.