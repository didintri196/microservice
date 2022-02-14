
# Microservice

Gabungan service dari bahasa pemograman go, nodejs

## Demo API

ms-auth  : [https://lab.integrasi.in/user](https://lab.integrasi.in/user) 

ms-fetch : [https://lab.integrasi.in/fetch](https://lab.integrasi.in/fetch)



## Installation

Clone the project

```bash
  git clone https://github.com/didintri196/microservice
```

Go to the project directory

```bash
  cd microservice
```

Start the server

```bash
  docker-compose up -d
```


## Documentation

[Markdown Documentation](https://github.com/didintri196/microservice/blob/master/API.md)

[Postman Documentation](https://documenter.getpostman.com/view/3419442/UVeMKQFS)


## Context & Deployment

### Context
![Logo](https://raw.githubusercontent.com/didintri196/microservice/master/dokumentasi/structurizr-MicroserviceSystem-SystemContext.png)


### Deployment
![Logo](https://raw.githubusercontent.com/didintri196/microservice/master/dokumentasi/structurizr-LiveDeployment-001.png)


## Tech Stack

**Server:** Node, Go, Docker


## Authors

- [@didintri196](https://www.github.com/didintri196)


## Goals

- [x]  Servers bisa dinyalakan di port berbeda
- [x]  Semua endpoint berfungsi dengan semestinya (3 endpoint auth, 3 endpoint fetch)
- [x]  Dokumentasi endpoint dengan format OpenAPI (API.md)
- [x]  Dokumentasi system diagram-nya dalam format C4 Model (Context dan Deployment)
- [x]  Pergunakan satu repo git untuk semua apps (mono repo)
- [x]  Dockerfile untuk masing-masing app
- [x]  Petunjuk penggunaan dan instalasi di README.md yang memudahkan

## Additional Goals

- [x]  Deployed ke Host/Penyedia Layanan (semacam surge, heroku, vercel, firebase, glitch,
host anda pribadi)
- [x]  Docker Compose
- [ ]  Unit Testing
