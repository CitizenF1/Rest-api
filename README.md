# Rest-api

```
docker-compose up -d --build
```

* `GET` : /rest/email/check

<img src="./img/screen1.jpg" height="300px"/>

##### ИИН

<img src="./img/screen2.jpg" height="300px"/>

##### PostgreSQL Table
```
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name varchar(255) NOT NULL,
    last_name varchar(255) NOT NULL
);
```

###### Postgre routes

###### create user

* `POST` : /rest/user

<img src="./img/screen3.jpg" height="300px"/>

###### Get user
* `GET` : /rest/user/{id}

###### Update user
* `PUT` : /rest/user/{id}

<img src="./img/screen.jpg" height="300px"/>

###### Delete user
* `DELETE` : /rest/user/{id}

<img src="./img/screen4.jpg" height="300px"/>
