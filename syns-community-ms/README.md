<p align="center">
<br />
<a href="https://github.com/syns-platform/servers"><img src="https://github.com/syns-platform/materials/blob/master/logo.svg?raw=true" width="150" alt=""/></a>
<h1 align="center">SYNS - Spark Your Noble Story - v2.0</h1>
<h4 align="center">Community Microservice</h4>

## Folder structure

    .
    ├── controllers
    │   ├──community
    │   │  └── syns.comm-controller.go
    │   └──post
    │      └── syns.post-controller.go
    ├── dao
    │   ├──community
    │   │  ├── syns.comm-dao-impl.go
    │   │  └── syns.comm-dao-interface.go
    │   └──post
    │      ├── syns.post-dao-impl.go
    │      └── syns.post-dao-interface.go
    ├── db
    │   └── syns.community-db.go
    ├── middleware
    │   └── syns.auth-middleware.go
    ├── models
    │   └── syns.community-models.go
    ├── routers
    │   ├──community
    │   │  └── syns.comm-router.go
    │   └──post
    │      └── syns.post-router.go
    ├── utils
    │   └── syns.dao-impl.go
    ├── .example.env
    ├── .gitignore
    ├── Makefile
    ├── README.md
    ├── go.mod
    ├── go.sum
    └── syns.community-main.go

## Getting Started

### Requirement

- [git](https://git-scm.com/)
- [golang](https://go.dev/)
<!-- - [docker](https://www.docker.com/) -->

### Clone the repo

```bash
git clone https://github.com/syns-platform/servers.git
cd syns-community-ms
```

### Set up environment variables

At the root of the directory, create a .env file using .env.example as the template and fill out the variables.

### Running the project

Build and run `syns-community-ms` locally using `Make` scripts

```bash
make dev-mode
```

<!-- 2. Build and run `agent` on Docker using `Make` scripts

```bash
make build-app
``` -->
