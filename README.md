# Golang Simple REST API 💃

💃 Golang Rest Api with bearer token JWT Authentication 💃

## Technology
- Language (golang)
- Database (mysql)
### Libraries
- Router (gorilla/mux)
- Server (net/http)
- JWT (dgrijalva/go-jwt)<!-- - Password Encryption (bcrypt) -->
- Database ORM (gorm) 
- Live Reload (cosmtrek/air)

### Run This Project : go run src/main.go

### Attention : you must signup for login because database fisrt is empty

## API Documentation in Postman = https://www.postman.com/dark-crescent-779247/workspace/team-workspace/request/6588994-187bcca6-4dc3-48d6-98ab-52514902f080

## API Documentation

> **POST** ``/auth/signup``

Create a new user in the database.

##### Body

```json
{
    "name": "admin baru",
    "email": "admin_new@ymail.com",
    "username": "adminnew",
    "password": "qwerty123",
    "role": "user-login" //role admin -> "administrator" || role user -> "user-login"
}
```

#### Output

```json
{
    "User": {
        "ID": 4,
        "CreatedAt": "2022-07-09T13:48:03.239+07:00",
        "UpdatedAt": "2022-07-09T13:48:03.239+07:00",
        "DeletedAt": null,
        "name": "admin baru",
        "email": "admin_new@ymail.com",
        "username": "adminnew",
        "password": "qwerty123",
        "role": "user-login"
    },
    "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWlubmV3IiwiZW1haWwiOiJrZWVtcGF0QHNzLmRkZGQiLCJpZCI6NCwiZXhwIjoxNjU3OTU0MDgzfQ.TMZ44E7O3gLliB14ZixAIVxVPf02pcnLqkTgB8asg_8"
}
```

### Authentication
> **POST** ``/auth/login``

Login with username and password.

##### Body

```json
{
     "username": "adminnew",
     "password": "qwerty123"
}
```

#### Output

```json
{
    "User": {
        "ID": 4,
        "CreatedAt": "2022-07-09T13:48:03.239+07:00",
        "UpdatedAt": "2022-07-09T13:48:03.239+07:00",
        "DeletedAt": null,
        "name": "admin baru",
        "email": "admin_new@ymail.com",
        "username": "adminnew",
        "password": "qwerty123",
        "role": "user-login"
    },
    "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWlubmV3IiwiZW1haWwiOiJrZWVtcGF0QHNzLmRkZGQiLCJpZCI6NCwiZXhwIjoxNjU3OTU0MDgzfQ.TMZ44E7O3gLliB14ZixAIVxVPf02pcnLqkTgB8asg_8"
}
```

### Data Manipulation

All endpoints are protected, must send valid **jwt** as ``Authorization`` header with each request.

> **GET** &nbsp; ``/getprofile``

Get User Profile (current user login)

#### Output

```json
{
    "success": true,
    "status": 200,
    "data": {
        "ID": 2,
        "CreatedAt": "2022-07-09T13:23:18.252+07:00",
        "UpdatedAt": "2022-07-09T13:23:18.252+07:00",
        "DeletedAt": null,
        "Name": "user 1",
        "Email": "user@mail.co",
        "Phone": "user456",
        "Password": "qwert123",
        "Role": "user-login"
    }
}
```


> **GET** &nbsp; ``/profile/all``

Get All User Profile

#### Output

```json
{
    "success": true,
    "status": 200,
    "data": [
        {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "Name": "testing",
            "Email": "testin@mail.co",
            "Phone": "test123",
            "Password": "********",
            "Role": "********"
        },
        {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "Name": "user 1",
            "Email": "user@mail.co",
            "Phone": "user456",
            "Password": "********",
            "Role": "********"
        },
        {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "Name": "admin 1",
            "Email": "admin@mail.co",
            "Phone": "admin890",
            "Password": "********",
            "Role": "********"
        }
    ]
}
```

> **GET** &nbsp; ``/profile/<id>``

Get single User by id

#### Output

```json
{
    "success": true,
    "status": 200,
    "data": {
        "ID": 6,
        "CreatedAt": "2022-07-02T22:31:08.395+07:00",
        "UpdatedAt": "2022-07-02T22:31:08.395+07:00",
        "DeletedAt": null,
        "Name": "testing 123",
        "Email": "yayayaya@gmail.org",
        "Phone": "same123",
        "Password": "09089876"
    }
}
```

> **POST** &nbsp; ``/profile``

Add a new User to the database.

##### Body

```json
{
    "name": "testing 123",
    "email": "yayayaya@gmail.org",
    "username": "same123",
    "password": "09089876"
}
```

#### Output

```json
{
    "success": true,
    "status": 200,
    "data": {
        "ID": 6,
        "CreatedAt": "2022-07-02T22:31:08.395+07:00",
        "UpdatedAt": "2022-07-02T22:31:08.395+07:00",
        "DeletedAt": null,
        "Name": "testing 123",
        "Email": "yayayaya@gmail.org",
        "Phone": "same123",
        "Password": "09089876"
    }
}
```
> **DELETE** &nbsp; ``/profile/<id>``

Delete one User from the database. (only input deleted_at not delete real data)

#### Output

```json
{
    "success": true,
    "status": 200,
    "messages": "success deleted data id = <id>"
}
```


> **PUT** &nbsp; ``/update/profile/<id>``

Add a new User to the database.

##### Body

```json
{
    "name": "sdasdas",
    "email": "ketdasdasdiga@gg.hh",
    "username": "asda",
    "password": "asdasd",
    "role": "administrator"
}
```

#### Output

```json
{
    "success": true,
    "status": 200,
    "data": {
        "CreatedAt": "2022-07-02T22:31:08.395+07:00",
        "UpdatedAt": "2022-07-02T22:31:08.395+07:00",
        "DeletedAt": null,
        "Name": "testing 123",
        "Email": "yayayaya@gmail.org",
        "Phone": "same123",
        "Password": "********"
    }
}
```

if you login as user-login but access api administrator

#### Output

```json
{
    "message": "accses denied admin only",
    "status": "500",
    "success": "false"
}
```"# poker-api" 
"# poker-api" 
"# poker-api" 
"# poker-api" 
"# poker-api" 
