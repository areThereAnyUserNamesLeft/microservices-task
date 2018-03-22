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

    - Ubuntu
    - MySQL
    - GO
    - ...

3. Build UI

    - Not the crocodile closest to the boat right now


##### Where can I cut corners for speed and not introduce complication of writing from scratch?

Borrow API code from [http://himarsh.org/build-restful-api-microservice-with-go/](http://himarsh.org/build-restful-api-microservice-with-go/) as a boilerplate for a RestAPI and Tests. Also kind of documented which is nice. This pretty much covers the backend of the project brief and has a Dockerfile ready and waiting in the accompanying repo. Taking the pressure off my writing code and allowingme to just refactor and repurpose.M
