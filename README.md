<p align="center">
<br />
<a href="https://github.com/SWYLy/servers"><img src="https://github.com/SWYLy/materials/blob/master/logo.svg?raw=true" width="150" alt=""/></a>
<h1 align="center">SWYL - Support Who You Love - v2.0</h1>
</p>

## Overview

The SWYL platform's backend and engine is powered by a collection of Golang MVC microservices known as **_SWYL/Servers_**. This restful server hosts plenty endpoints that enable SWYL users to easily interact with the platform and store their platform information in a MongoDB Atlas client.

## Highlighted Features:

- [Swyl User Microservice](https://github.com/SWYLy/servers/tree/master/swyl-users-ms) allows users to connect their cryptocurrency wallet and sign up for a record in the MongoDB database on their first connection. It also allows users to configure their account or deactivate (delete) it if desired.

- [Swyl Club Microservice](https://github.com/SWYLy/servers/tree/master/swyl-club-ms) mimics the logic of the [SwylClub smart contract](https://github.com/SWYLy/contracts/blob/main/contracts/v1/SwylClub.sol) and maintains off-chain records about `SWYL Clubs`.

- [Swyl Community Microservice](https://github.com/SWYLy/servers/tree/master/swyl-community-ms) is the \***\*`most honored`\*\*** out of the three microservices. Holding many complex endpoints, \***\*Swyl Community Microservice\*\*** was developed with the vision of being a full-fledged social media platform which is able to:

  - Allow potential community owners to create an online community where followers can follow for free to receive updates from the owner.

  - Enable community owners to create posts or blogs that followers can interact with by reacting with one of the four `SWYL emotions` (SUPPORT, BRAVO, LAUGH, FIRE), leaving comments, and editing or deleting them as desired.

  - Allow other users to reply to comments and react to them with the `SWYL emotions`.

## Resources

- [Golang](https://go.dev/)
- [Gin-Gonic/Gin](https://github.com/gin-gonic/gin)
- [MongoDB-Atlas](https://www.mongodb.com/atlas)
- [Mongo Go Driver](https://www.mongodb.com/docs/drivers/go/current/)
- [Godotenv](https://github.com/joho/godotenv)
- [Go Validator](https://pkg.go.dev/github.com/go-playground/validator/v10#section-readme)
- [Docker](https://www.docker.com/) - comming soon...
