<p align="center">
<br />
<a href="https://github.com/syns-platform"><img src="https://github.com/syns-platform/materials/blob/master/logo.svg?raw=true" width="150" alt=""/></a>
<h1 align="center">SYNS - Spart Your Noble Story - v2.0</h1>
</p>

<h5 align="center"> 🏵️ The platform is currently open for beta testing at https://syns.vercel.app </h5>

<div align="center" >

![](https://img.shields.io/badge/Golang-1.9.0-blue?style=flat-square&logo=go)
![](https://img.shields.io/badge/MongoDB-6.0.0-blue?style=flat-square&logo=mongodb)
</div>

## Overview

**_SYNS/Servers_**, an architecture comprising of a collection of Golang-based RESTful microservices, provide a plethora of endpoints that facilitate seamless user interactions with the platform and enable the storage of platform-related information in a MongoDB Atlas client, thus ensuring a robust and efficient data management infrastructure.


## Highlighted Features:

- [SYNS User Microservice](https://github.com/syns-platform/servers/tree/master/syns-users-ms) allows users to connect their cryptocurrency wallet and sign up for a record in the MongoDB database on their first connection. It also allows users to configure their account or deactivate (delete) it if desired.

- [SYNS Club Microservice](https://github.com/syns-platform/servers/tree/master/syns-club-ms) mimics the logic of the [SynsClub smart contract](https://github.com/syns-platform/contracts/blob/main/contracts/v1/SynsClub.sol) and maintains off-chain records about `SYNS Clubs`.

- [Syns Community Microservice](https://github.com/syns-platform/servers/tree/master/syns-community-ms) is the \***\*`most honored`\*\*** out of the three microservices. Holding many complex endpoints, \***\*Syns Community Microservice\*\*** was developed with the vision of being a full-fledged social media platform which is able to:

  - Allow potential community owners to create an online community where followers can follow for free to receive updates from the owner.

  - Enable community owners to create posts or blogs that followers can interact with by reacting with one of the four `SYNS emotions` (SUPPORT, BRAVO, LAUGH, FIRE), leaving comments, and editing or deleting them as desired.

  - Allow other users to reply to comments and react to them with the `SYNS emotions`.

## Resources

- [Golang](https://go.dev/)
- [Gin-Gonic/Gin](https://github.com/gin-gonic/gin)
- [MongoDB-Atlas](https://www.mongodb.com/atlas)
- [Mongo Go Driver](https://www.mongodb.com/docs/drivers/go/current/)
- [Godotenv](https://github.com/joho/godotenv)
- [Go Validator](https://pkg.go.dev/github.com/go-playground/validator/v10#section-readme)
- [Docker](https://www.docker.com/) - comming soon...
