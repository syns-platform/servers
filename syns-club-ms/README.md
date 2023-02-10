<p align="center">
<br />
<a href="https://github.com/syns-platform/servers"><img src="https://github.com/syns-platform/materials/blob/master/logo.svg?raw=true" width="150" alt=""/></a>
<h1 align="center">SYNS - Spark Your Noble Story - v2.0</h1>
<h4 align="center">Club Microservice</h4>

## Folder structure

    .
    ├── controllers
    │   ├──club
    │   │  └── syns.club-controller.go
    │   ├──subscription
    │   │  └── syns.sub-controller.go
    │   └──tier
    │      └── syns.tier-controller.go
    ├── dao
    │   ├──club
    │   │  ├── syns.club-dao-impl.go
    │   │  └── syns.club-dao-interface.go
    │   ├──subscription
    │   │  ├── syns.sub-dao-impl.go
    │   │  └── syns.sub-dao-interface.go
    │   └──tier
    │      ├── syns.tier-dao-impl.go
    │      └── syns.tier-dao-interface.go
    ├── db
    │   └── syns.db.go
    ├── middleware
    │   └── syns.auth-middleware.go
    ├── models
    │   └── syns.models.go
    ├── routers
    │   ├──club
    │   │  └── syns.club-router.go
    │   ├──subscription
    │   │  └── syns.sub-router.go
    │   └──tier
    │      └── syns.tier-router.go
    ├── utils
    │   └── syns.dao-impl.go
    ├── .example.env
    ├── .gitignore
    ├── Makefile
    ├── README.md
    ├── go.mod
    ├── go.sum
    └── syns.club-main.go

## Getting Started

### Requirement

- [git](https://git-scm.com/)
- [golang](https://go.dev/)
<!-- - [docker](https://www.docker.com/) -->

### Clone the repo

```bash
git clone https://github.com/syns-platform/servers.git
cd syns-club-ms
```

### Set up environment variables

At the root of the directory, create a .env file using .env.example as the template and fill out the variables.

### Running the project

Build and run `syns-club-ms` locally using `Make` scripts

```bash
make dev-mode
```

<!-- 2. Build and run `agent` on Docker using `Make` scripts

```bash
make build-app
``` -->
