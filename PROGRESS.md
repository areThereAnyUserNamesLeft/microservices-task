# My Progress - Richard Pape

#### Task breakdown

1. build MS

    Ideas
 - **Restful or gRPC?** Restful asa I am only just getting to grips with gRPC and would not want to spoil an oppertunity just to be showing off that I know something fancy.

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


##### Where can I cut corners for speed and not introduce complication of writing from scratch?

###### ~0930

Borrow API code from [http://himarsh.org/build-restful-api-microservice-with-go/](http://himarsh.org/build-restful-api-microservice-with-go/) as a boilerplate for a RestAPI and Tests. Also kind of documented which is nice. This pretty much covers the backend of the project brief and on closer inspection has a Dockerfile ready and waiting in the accompanying repo. Taking the pressure off my writing code and allowing me to just refactor and repurpose what is here into a useable API for my nafarious purposes (Mwa ha ha ha ha...).

###### ~1100

Code refactored and all seems to be running. The example I chose to use has a MySQL DB running locally and serving on port :3306 (as usual) - I guess I need to make that now.

- **~~Plan Ai)~~**

I've not used MySQL in a little while having moved to Arch Linux, who as a community use MariaDB over MySQL - see [https://wiki.archlinux.org/index.php/MySQL](https://wiki.archlinux.org/index.php/MySQL). This creates a blocker, so I've got some work to do here to make sure it is working locally.

- **Plan Aii)**

Look at dockerized versions of MySQL - probably should have used this as my first approch **there is a lesson to be learnt here**...

- **Plan B)**

If it seems like a non starter by around Lunchtime, I'll  look to remove the MySQL DB from the boilerplate and replace it with text file or Json file as storage to allow me to look at other features - this also has an advantage of not requiring Sipsynergy to configure a DB for looking at the project.

###### ~1200

**Choice made** - I downloaded a copy of MySQL from the AUR (Arch user repo) [https://aur.archlinux.org/packages/mysql/](https://aur.archlinux.org/packages/mysql/) but it seems like it will need an unklnown amount of troubleshooting around chroot environments and permissions - in short a potential time suck.

Considering dockerized versions of MySQL [https://hub.docker.com/_/mysql/](https://hub.docker.com/_/mysql/) - Hitting some issues with Docker now so time for a system reboot...
###### ~1230

**Dockers working now!**
For my reference-

Launching with:

`sudo docker run --name userAPI-mysql -e MYSQL_ROOT_PASSWORD=oranges -d mysql:8`

Container number: `659d98a83fb2703d461db02be49a0ea782ff146a1349e9d13fe93fe9d67bbb8c`

Name:`userAPI-mysql`

MySQL version: `8`

password: `oranges`
