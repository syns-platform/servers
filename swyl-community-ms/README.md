<p align="center">
<br />
<a href="https://github.com/SWYLy/servers"><img src="https://github.com/SWYLy/materials/blob/master/logo.svg?raw=true" width="150" alt=""/></a>
<h1 align="center">SWYL - Support Who You Love - v2.0</h1>
<h4 align="center">Community Microservice</h4>

## Folder structure

    .
    ├── controllers
    │   ├──community
    │   │  └── swyl.comm-controller.go
    │   └──post
    │      └── swyl.post-controller.go
    ├── dao
    │   ├──community
    │   │  ├── swyl.comm-dao-impl.go
    │   │  └── swyl.comm-dao-interface.go
    │   └──post
    │      ├── swyl.post-dao-impl.go
    │      └── swyl.post-dao-interface.go
    ├── db
    │   └── swyl.community-db.go
    ├── middleware
    │   └── swyl.auth-middleware.go
    ├── models
    │   └── swyl.community-models.go
    ├── routers
    │   ├──community
    │   │  └── swyl.comm-router.go
    │   └──post
    │      └── swyl.post-router.go
    ├── utils
    │   └── swyl.dao-impl.go
    ├── .example.env
    ├── .gitignore
    ├── Makefile
    ├── README.md
    ├── go.mod
    ├── go.sum
    └── swyl.community-main.go

## Getting Started

### Requirement

- [git](https://git-scm.com/)
- [golang](https://go.dev/)
<!-- - [docker](https://www.docker.com/) -->

### Clone the repo

```bash
git clone https://github.com/SWYLy/servers.git
cd swyl-community-ms
```

### Set up environment variables

At the root of the directory, create a .env file using .env.example as the template and fill out the variables.

### Running the project

Build and run `swyl-community-ms` locally using `Make` scripts

```bash
make dev-mode
```

<!-- 2. Build and run `agent` on Docker using `Make` scripts

```bash
make build-app
``` -->
