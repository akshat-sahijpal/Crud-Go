## Crud-Go
### Simple RESTful API to register, view, update and delete users.
#### Run the main.go file
## Endpoints 

### Get All Users
``` bash
GET /getAllUsers
```
### Get User For UserId
``` bash
GET /getUser/{id}
```

### Delete a User for a UserId 
``` bash
DELETE /deleteUser/{id}
```

### Register a User
``` bash
POST /register
```

``` bash
In PostMan Body Give Input 
  {
     "username": "@sahij.akshat",
     "fullName": "Akshat",
     "age": 21,
     "jobtitle": "Student"
  }
```

### Update a Particular User For UserID
``` bash
PUT /updateUser/{id}
``` 

``` bash
In PostMan Body Give Input 
  {
     "username": "@sahij.akshat",
     "fullName": "Akshat",
     "age": 21,
     "jobtitle": "Student"
  }
```


## Demo-Video
https://user-images.githubusercontent.com/80918746/160616718-88e7d14a-74a1-419e-b36e-083fdf337bd9.mp4



