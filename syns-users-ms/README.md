<p align="center">
<br />
<a href="https://github.com/syns-platform"><img src="https://github.com/syns-platform/materials/blob/master/main_logos/Syns_Official_Main_Logo_V3.svg?raw=true" width="150" alt=""/>
<h1 align="center">SYNS - Spark Your Noble Story - v2.0</h1>
<h4 align="center">User Microservice</h4>

## Folder structure

    .
    ├── controllers
    │   └── syns.user-controller.go
    ├── dao
    │   ├── syns.dao-impl.go
    │   └── syns.dao-interface.go
    ├── db
    │   └── syns.user-db.go
    ├── middleware
    │   └── syns.auth-middleware.go
    ├── models
    │   └── syns.user-model.go
    ├── routers
    │   └── syns.router.go
    ├── utils
    │   └── syns.dao-impl.go
    ├── .example.env
    ├── .gitignore
    ├── Makefile
    ├── README.md
    ├── go.mod
    ├── go.sum
    └── syns.user-main.go

## Getting Started

### Requirement

- [git](https://git-scm.com/)
- [golang](https://go.dev/)
<!-- - [docker](https://www.docker.com/) -->

### Clone the repo

```bash
git clone https://github.com/syns-platform/servers.git
cd syns-users-ms
```

### Set up environment variables

At the root of the directory, create a .env file using .env.example as the template and fill out the variables.

### Running the project

Build and run `syns-users-ms` locally using `Make` scripts

```bash
make dev-mode
```

<!-- 2. Build and run `agent` on Docker using `Make` scripts

```bash
make build-app
``` -->
