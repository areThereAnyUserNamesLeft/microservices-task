# My Progress - Richard Pape

I am choosing to use Go for the backend as it seems a good choice bearing in mind the JD - I can't say I have built a microservice aside from cloud functions on AWS Lambda and DynamoDB where alot of the work is done for you so I am enjoying this as a challenge.

## TL;DR

I did not finish 100%, in fact I finished the API and packaged it as a Microsevice. I had a plan to use a MySQL DB as the storage which sucked up a lot of time and really was not nesessary for the small scale of the task.

In the end I implemented a nice little protocol that writes to a CSV and altered it 5 different ways to accomedate the different API calls.

There is a Task breakdown below followed by a rough time line I put together to "show my workings" - I Really have tried to be honest as I can and not try to hide any bad decissions I made as I think it is only fair to show you my process warts and all.

I'd like to think i'd have finished the task had I not spent 11-1630 trying to implement the DB - Live and Learn, I guess.

### Task breakdown & initial thoughts

1. build MS

    Ideas
 - **Restful or gRPC?** Restful as I am only just getting to grips with gRPC and would not want to spoil an oppertunity just to be showing off that I think know something fancy better than I really do.

 Use suggested endpoints - I'll use these as they'll tick the boxes:

 - [GET] /users - Returns a list of user ids
 - [GET] /users/{id} - returns a users specific details
 - [POST] /users - Returns a list of user ids
 - [POST] /users/{id} - Returns a users specific details

2. Dockerize

    - Ubuntu?
    - MySQL
    - GO?
    - ...

3. Build UI

    - Not the crocodile closest to the boat right now


##### Where I can I'll cut corners for speed and not introduce complication of writing from scratch?

###### ~0930 Planning / Research

Borrowed / Repurposed API code from [http://himarsh.org/build-restful-api-microservice-with-go/](http://himarsh.org/build-restful-api-microservice-with-go/) as a boilerplate for a RestAPI and Tests. It is a lso kind of documented which is nice. This pretty much covers the backend of the project brief and on closer inspection has a Dockerfile ready and waiting in the accompanying repo. Taking the pressure off my writing code and allowing me to just refactor and repurpose what is here into a useable API for my ###### ~2350 Delete working now for updatenafarious purposes (Mwa ha ha ha ha...).

###### ~1100 Copy/Paste/Refactor - DB Choices

Code refactored and all seems to be running. The example I chose to use has a MySQL DB running locally and serving on port :3306 (as usual) - I guess I need to make that now.

- **~~Plan Ai)~~**

I've not used MySQL in a little while having moved to Arch Linux, who as a community use MariaDB over MySQL - see [https://wiki.archlinux.org/index.php/MySQL](https://wiki.archlinux.org/index.php/MySQL). This creates a blocker, so I've got some work to do here to make sure it is working locally.

- **Plan Aii)**

Look at dockerized versions of MySQL - probably should have used this as my first approch *there is a lesson to be learnt here*...

- **Plan B)**

If it seems like a non starter by around Lunchtime, I'll  look to remove the MySQL DB from the boilerplate and replace it with text file or Json file as storage to allow me to look at other features - this also has an advantage of not requiring Sipsynergy to configure a DB for looking at the project.

###### ~1200 Choice Made!

**Choice made** - I downloaded a copy of MySQL from the AUR (Arch user repo) [https://aur.archlinux.org/packages/mysql/](https://aur.archlinux.org/packages/mysql/) but it seems like it will need an unklnown amount of troubleshooting around chroot environments and permissions - in short a potential time suck.

Considering dockerized versions of MySQL [https://hub.docker.com/_/mysql/](https://hub.docker.com/_/mysql/) - Hitting some issues with Docker now so time for a system reboot...

###### ~1230 Docker nightmare (turned out to be MySql versioning)

**Dockers working now!**
For my reference-

Launching with:

`sudo docker run --name userAPI-mysql -e MYSQL_ROOT_PASSWORD=oranges -d mysql:8`

Container number: `659d98a83fb2703d461db02be49a0ea782ff146a1349e9d13fe93fe9d67bbb8c`

Name:`userAPI-mysql`

MySQL version: `8`

password: `oranges`

CLI connect command:
`docker run -it --link userAPI-mysql:mysql --rm mysql sh -c 'exec mysql -h"$MYSQL_PORT_3306_TCP_ADDR" -P"$MYSQL_PORT_3306_TCP_PORT" -uroot -p"$MYSQL_ENV_MYSQL_ROOT_PASSWORD"'``

###### ~1310 LunchTime - One last try!

Before admitting defeat and turning to plan B as MySQL is throwing an error. I am going to create a fresh image with an older version of MySQL as 8 seems to have a new plugin(!?! - latest is not always best) and throws this error...

`ERROR 2059 (HY000): Authentication plugin 'caching_sha2_password' cannot be loaded: /usr/lib/mysql/plugin/caching_sha2_password.so: cannot open shared object file: No such file or directory`

Launched command:

`sudo docker run --name microservice-mysql -e MYSQL_ROOT_PASSWORD=oranges -d mysql:5.7`

Containeer Number:`dcdfac0fe16a7b2608ba7c7ffd8d5bc4165e0faf1047cb17edc9ed3ddae36bf2`

Name:`microservice-mysql`

MySQL version: `5.7`

password: `oranges`

CLI connect command:
`docker run -it --link microservice-mysql:mysql --rm mysql sh -c 'exec mysql -h"$MYSQL_PORT_3306_TCP_ADDR" -P"$MYSQL_PORT_3306_TCP_PORT" -uroot -p"$MYSQL_ENV_MYSQL_ROOT_PASSWORD"'``

###### ~1320 Success

**I'm in to mySQL CLI** - *Lesson here - MySQL 8 needs a plugin to avoid SHA2 password encryption error and docker hub documentation out of date currently around this!*

###### ~1330 Database sculpting

### DB build:

I am going to make only one table with three columns as below

### Table: USERS

|columns name|id|user_name|user_status|date|
|-|---|----------|------------|-|
|**Data type**|*int64=>BIGINT*|*string=>varchar(100) |string=>varchar(40) (I considered a bool but may want more options so will make it a str)|date=>date|

Manual commands - I might get to putting these isn a dockerfile but right now I'll put them here and add them to a config at the end...

`CREATE DATABASE USERS`

`create table users(id BIGINT NOT NULL AUTO_INCREMENT, user_name VARCHAR(100) NOT NULL, user_status VARCHAR(40), date DATE, PRIMARY KEY (id));
`

`insert into USERS (user_name, user_status,date) VALUES ("Richard Pape", "Present", NOW());
`

` select * from users;`

<pre>
+----+--------------+-------------+------------+
| id | user_name    | user_status | date       |
+----+--------------+-------------+------------+
|  1 | Richard Pape | Present     | 2018-03-22 |
+----+--------------+-------------+------------+
</pre>

###### ~1400 Connecting stuff and E2E API testing

Run: `docker ps` and not container id

Then: `docker inspect container_id | grep IPAddress`

###### ~1500 Troubleshooting - So close

Tests are failing to get to the DB with connection refused

I have commented out the lines that try to create the DB in the db.go file so I get more meaning ful messages from the tests.

 - GET  tests = 404s

 - POST test  = 200
 - PUT  test  = 404

Curl comes back with the same...

 `docker `

<code>HTTP/1.1 404 Not Found
Content-Type: application/json; charset=utf-8
Date: Thu, 22 Mar 2018 15:24:53 GMT
Content-Length: 37
{"error":"no user(s) into the table"}</code>

2 ideas -

1. Move on to dockerizing the Go app to link the application within containers. I believe that the connection could be refused due to the docker MySQL container not being given express permission.
2. Keep trying to figure it out via Google searches

I am going to choose option 1. as it is aligned with where I am heading anyway and if it seems more likely to be the issue.

building dockerfile went well

Now first a bit of a brain break.

###### ~1630 Off to the Kids Martial Arts Class

![http://toonbarn.com/wordpress/wp-content/uploads/2011/08/Nickelodeon-and-IDW-unveil-new-Teenage-Mutant-Ninja-Turtles-comic-book.jpg](http://toonbarn.com/wordpress/wp-content/uploads/2011/08/Nickelodeon-and-IDW-unveil-new-Teenage-Mutant-Ninja-Turtles-comic-book.jpg)

###### ~1830 Back from the Kids Martial Arts Class

Give permissions for app to speak to DB

`sudo docker run --name microservices-task --link microservice-mysql:mysql -d microservices-task`

Then `ran docker inspect...`

and got <code>            "SecondaryIPAddresses": null,
"IPAddress": "",
"IPAddress": "",
</code> So it looks like I need to rebuild with a valid IP exposed rather than just a port!

Looked at `docker logs xxx`

and see my old friend

`Create table failed dial tcp 172.17.0.4:3306: connect: connection refused`

The program is running but does not reach the DB

###### ~1930 Is it time for **plan B)**? - *kick MySQL to the kerb*

I'd love to run MySQL as part of this but the brief does not require it and I have already spent 1/2 a day trying to get a good connection - It is not worth it bearing in mind the small scall of the DB needed - I am sad to say -I am going to build a flat file store and replace the DB connection, ***Really, Really big lesson I have learnt here*** but conversely I have built the docker container and it is working which is something worth while.

###### ~2015 Back on target - kicking myself for not admitting failiure earlier
Running =>
`curl -i http://localhost:8080/api/v1/users`
Returns =>
<pre>
`[{"id":"1","user_name":" Richard Pape","user_status":" Present","date":" 22-03-2018"}]`
</pre>


###### ~2040 two down...

Now both get requests are behaving

###### ~2150 Post request working - now UPDATE

Made a delete function while working on update

###### ~2230 Delete working now for update

OK finally managed to do some coding and the API is working flawlessly

Although I think the update function will allow a user to update with blank values.

###### ~2350 Delete working now for update

Spent some time building the docker image to realise I have not compiled the main.go - and test was failing where it was still looking for the DB - *Rookie mistake*

###### ~0010 Done!

Docker image working instructions at the bottom

#### Have a play --- Local host examples

If you `go run main.go` you can have a play using the following curls

- `curl -i http://localhost:8080/api/v1/users`
    - Should give a list of users in Json...
- `curl -i http://localhost:8080/api/v1/users/1`
    - Should just give you one user...
- `curl -i -X POST -H "Content-Type: application/json" -d "{ \"user_status\": \"No longer potus\", \"user_name\": \"Donald Trump\", "date": "04-11-2020" }" http://localhost:8080/api/v1/users`
    - Will create a new user
- `curl -i -X PUT -H "Content-Type: application/json" -d "{ \"user_status\": \"Potus\", \"user_name\": \"Donald Trump\", \"date\": \"01-04-2018\" }" http://localhost:8080/api/v1/users/3
`   
    - Will update all of the fields

- `curl -i -X DELETE http://localhost:8080/api/v1/users/3`
    - Will remove the person at id number 3

You'll see the affects in the users.csv file which took the place of the DB as each of these are called.


### Docker instructions

Run

=>`docker build -t microservice-test .`

then

=> `docker run -P microservice-test`

then `docker ps`

Grab the `<Container-ID\>`

then

`docker inspect <Container-ID> | grep IPAddress`

grab the `<\IP-Address\>`

Then

`curl -i <\IP-Address\>:8080/api/v1/users`

just replace Localhost in the 'go run ...' examples with the `<\IP-Address\>`


<pre>
curl -i <\IP-Address\>:8080/api/v1/users
</pre>

- Should give a list of users in Json...

<pre>
curl -i <\IP-Address\>:8080/api/v1/users/1
</pre>

- Should just give you one user...

<pre>
curl -i -X POST -H "Content-Type: application/json" -d "{ \"user_status\": \"No longer potus\", \"user_name\": \"Donald Trump\", "date": "04-11-2020" }" <\IP-Address\>:8080/api/v1/users
</pre>

- Will create a new user

<pre>

curl -i -X PUT -H "Content-Type: application/json" -d "{ \"user_status\": \"Potus\", \"user_name\": \"Donald Trump\", \"date\": \"01-04-2018\" }" <\IP-Address\>:8080/api/v1/users/3
</pre>

- Will update all of the fields

<pre>
curl -i -X DELETE <\IP-Address\>:8080/api/v1/users/3
</pre>

- Will remove the person at id number 3


## I think that about covers it...

Thanks for the oppertunity to do this - I have enjoyed it and learnt some lessons.

Sorry for the typos and I hope to speak in the future
