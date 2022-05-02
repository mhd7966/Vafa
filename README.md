# Vafa

For running migration, run this command ):

 Download swag by using:
```sh
$ go get -u github.com/swaggo/swag/cmd/swag
```
 Swagger init changes, run in main project directory : 
```
$ swag init
```

  To run program : 
```
$ docker-compose up
```
 Download sql-migrate by using:
```
$ go get -v github.com/rubenv/sql-migrate/...
```
For running migration, run this command (Note: The database must be up)
```
$ sql-migrate up
```

