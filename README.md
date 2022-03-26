# Golang Blog Api

This is simple blog api with golang and gin web framework. The main resources of this API is posts and categories. For authenticating i'm using jwt-go and gorm for database ORM. Installation guide, api endpoints and sample request/response are given below.

## For installing this Application please do this following steps

```
1. At first create a database by psql or pgAdmin
2. Then clone repository: git clone https://github.com/abdurraufraihan/golang-blog-api.git
3. Change directory to golang-blog-api: cd golang-blog-api
4. Add database configuration to config/database-config.go
5. Now install dependency: go mod download
6. Run the app: go run server.go
7. The project will now open on 8080 port of your localhost
8. Test with your favorite api client (e.g. postman)
```

## Endpoints

- POST api/v1/auth/signup
- POST api/v1/auth/login
- POST api/v1/auth/token/verify
- POST api/v1/auth/token/refresh
- GET api/v1/posts
- GET api/v1/posts/:id
- POST api/v1/posts
- PUT api/v1/posts/:id
- DELETE api/v1/posts/:id
- GET api/v1/categories
- POST api/v1/categories
- PUT api/v1/categories/:id
- DELETE api/v1/categories/:id

## Sample API Request and Response

##### POST api/v1/auth/signup

request body:

```json
{
  "name": "Raihan",
  "email": "abdurraufraihan@gmail.com",
  "password": "123465"
}
```

response body:

```json
{
  "id": 1,
  "name": "Raihan",
  "email": "abdurraufraihan@gmail.com"
}
```

##### POST api/v1/auth/login

request body:

```json
{
  "email": "abdurraufraihan@gmail.com",
  "password": "123465"
}
```

response body:

```json
{
  "access_token": "<access_token>",
  "refresh_token": "<refresh_token>"
}
```

##### POST api/v1/auth/token/verify

request body:

```json
{
  "token": "<access_token>"
}
```

response body:

```json
{
  "is_valid": true
}
```

##### POST api/v1/auth/token/refresh

request body:

```json
{
  "token": "<refresh_token>"
}
```

response body:

```json
{
  "access_token": "<access_token>",
  "refresh_token": "<refresh_token>"
}
```

##### GET api/v1/posts?limit=5&offset=0

response body:

```json
{
  "totalPost": 74,
  "posts": [
    {
      "id": 74,
      "title": "this is title",
      "description": "this is description",
      "image": "media/images/13-300x300.jpg",
      "created_at": "2022-03-05T23:50:17.526683+06:00",
      "updated_at": "2022-03-05T23:50:17.526683+06:00",
      "category": {
        "id": 1,
        "name": "category one"
      }
    } ......
  ]
}
```

##### GET api/v1/posts/74

response body:

```json
{
  "id": 74,
  "title": "this is title",
  "description": "this is description",
  "image": "media/images/13-300x300.jpg",
  "created_at": "2022-03-05T23:50:17.526683+06:00",
  "updated_at": "2022-03-05T23:50:17.526683+06:00",
  "category": {
    "id": 1,
    "name": "category one"
  }
}
```

##### POST api/v1/posts

request body:

```json
{
  "title": "this is title",
  "description": "this is description",
  "image": "<binary file>",
  "category": 1
}
```

response body:

```json
{
  "id": 74,
  "title": "this is title",
  "description": "this is description",
  "image": "media/images/13-300x300.jpg",
  "created_at": "2022-03-05T23:50:17.526683+06:00",
  "updated_at": "2022-03-05T23:50:17.526683+06:00",
  "category": {
    "id": 1,
    "name": "category one"
  }
}
```

##### PUT api/v1/posts/74

request body:

```json
{
  "title": "this is title (modified)",
  "description": "this is description (modified)",
  "image": "<binary file>",
  "category": 2
}
```

response body:

```json
{
  "id": 74,
  "title": "this is title (modified)",
  "description": "this is description (modified)",
  "image": "media/images/13-300x300.jpg",
  "created_at": "2022-03-05T23:50:17.526683+06:00",
  "updated_at": "2022-03-05T23:50:17.526683+06:00",
  "category": {
    "id": 2,
    "name": "category two"
  }
}
```

##### DELETE api/v1/posts/74

this will delete post with id 74 and response back a status 204 no content

##### GET api/v1/categories

response body:

```json
[
  {
    "id": 1,
    "name": "category one",
  },
  {
    "id": 2,
    "name": "category two",
  }, ......
]
```

##### POST api/v1/categories

request body:

```json
{
  "name": "category one"
}
```

response body:

```json
{
  "id": 1,
  "name": "category one"
}
```

##### PUT api/v1/categories/1

request body:

```json
{
  "name": "category one (modified)"
}
```

response body:

```json
{
  "id": 1,
  "name": "category one (modified)"
}
```

##### DELETE api/v1/categories/1

this will delete category with id 1 and response back a status 204 no content

```

```
