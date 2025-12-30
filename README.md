# Go Backend Docker
A simple Go Backend that uses a MySQL Database
You can run it using `docker compose up --build`

You can create users with `curl -X POST http://localhost:8080/user/ -d "username=johndoe" -d "password=secret"`

You can get all users with http://localhost:8080/user

You can get a specific user with http://localhost:8080/user/{userid}

You need to create a .env file with following variables
```
DB_USER=gotutorial
DB_PASSWORD=mydbpassword
DB_NAME=gotutorial
DB_HOST=mysql
DB_PORT=3306
```

Reference for http server and mysql: `https://gowebexamples.com/`
