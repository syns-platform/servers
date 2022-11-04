<p align="center">
<br />
<h1 align="center">SWYL - Support Who You Love - v1.0</h1>
<h4 align="center">Users Microservice</h4>

## Folder structure 

    .
    ├── controllers             
    │   ├──club  
    │   │  └── swyl.club-controller.go
    │   ├──subscription  
    │   │  └── swyl.sub-controller.go
    │   └──tier  
    │      └── swyl.tier-controller.go
    ├── dao
    │   ├──club     
    │   │  ├── swyl.club-dao-impl.go
    │   │  └── swyl.club-dao-interface.go
    │   ├──subscription     
    │   │  ├── swyl.sub-dao-impl.go
    │   │  └── swyl.sub-dao-interface.go
    │   └──tier     
    │      ├── swyl.tier-dao-impl.go
    │      └── swyl.tier-dao-interface.go
    ├── db
    │   └── swyl.db.go
    ├── middleware
    │   └── swyl.auth-middleware.go
    ├── models
    │   └── swyl.models.go
    ├── routers             
    │   ├──club  
    │   │  └── swyl.club-router.go
    │   ├──subscription  
    │   │  └── swyl.sub-router.go
    │   └──tier  
    │      └── swyl.tier-router.go
    ├── utils
    │   └── swyl.dao-impl.go
    ├── .example.env         
    ├── .gitignore      
    ├── Makefile
    ├── README.md
    ├── go.mod
    ├── go.sum
    └── swyl.club-main.go
    

## Getting Started

### Requirement

- [git](https://git-scm.com/)
- [golang](https://go.dev/)
<!-- - [docker](https://www.docker.com/) -->

### Clone the repo

```bash
git clone https://github.com/SWYLy/servers.git
cd swyl-users-ms
```

### Set up environment variables

At the root of the directory, create a .env file using .env.example as the template and fill out the variables.

### Running the project

Build and run `swyl-users-ms` locally using `Make` scripts

```bash
make dev-mode
```

<!-- 2. Build and run `agent` on Docker using `Make` scripts

```bash
make build-app
``` -->
