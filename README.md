<p align="center">
<br />
<a href="https://github.com/SWYLy/servers"><img src="https://github.com/SWYLy/materials/blob/master/logo.svg?raw=true" width="150" alt=""/></a>
<h1 align="center">SWYL - Support Who You Love - v2.0</h1>
</p>

## Overview

The SWYL platform's backend and engine is powered by a collection of Golang MVC microservices known as **_SWYL/Servers_**. This restful server hosts plenty endpoints that enable SWYL users to easily interact with the platform and store their platform information in a MongoDB Atlas client.

## Highlighted Features:

- [Swyl User Microservice](https://github.com/SWYLy/servers/tree/master/swyl-users-ms) allows users to connect their crypto wallet to the platform and sign up a record in `mongoDB` on the first connect. It also allows users to configure and/or deactivate (i.e. delete) own accounts if desired.

- [Swyl Club Microservice](https://github.com/SWYLy/servers/tree/master/swyl-club-ms) mimics the logic of the [SwylClub smart contract](https://github.com/SWYLy/contracts/blob/main/contracts/v1/SwylClub.sol) and keeps track of the off-chain data about `SwylClub`.

- [Swyl Community Microservice](https://github.com/SWYLy/servers/tree/master/swyl-community-ms) is the \***\*`most honored`\*\*** out of the three microservices. Holding many complex endpoints, \***\*Swyl Community Microservice\*\*** was developed with the vision of being a full-fledged social media platform which is able to:
  - allows a potential Swyl community owner to start an online community where potential followers can follow the community for free to keep getting updated by the community owner
  - allows Swyl community owner to create posts or blogs which then can be interacted by the followers. With that said, followers can react to the posts with the four `Swyl emotions` (SUPPORT, BRAVO, LAUGH, FIRE), leave comment on the posts which is then can be editted and/or deleted if desired.
  - allows other users to reply to comments and also react to comments with the `Swyl emotions`

## Resources

- [Golang](https://go.dev/)
- [Gin-Gonic/Gin](https://github.com/gin-gonic/gin)
- [MongoDB-Atlas](https://www.mongodb.com/atlas)
- [Mongo Go Driver](https://www.mongodb.com/docs/drivers/go/current/)
- [Godotenv](https://github.com/joho/godotenv)
- [Go Validator](https://pkg.go.dev/github.com/go-playground/validator/v10#section-readme)
- [Docker](https://www.docker.com/) - comming soon...
