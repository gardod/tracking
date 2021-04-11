# Go HTTP Tracking Service

We are running an HTTP server written in go that accepts HTTP and H2C requests at a specified port. We are using [gorilla/mux](github.com/gorilla/mux) for routing, [sirupsen/logrus](github.com/sirupsen/logrus) for logging, and [spf13/viper](github.com/spf13/viper) for config.

We divided the responsibility of serving a request to different modules for clarity and modularity. The modules follow a domain-driven design where package Model is at the center holding our entity definitions. The flow is as follows:
```
Request hits server -> Middlewares -> Handler -> Service -> Repository
```

## Internal

Everything service specific finds it's place in here.

### Middleware

All our requests go through them which helps with generating tracking id, logging, metrics etc. Currently only `recoverer` is included to narrow down the scope of the project.

### Handler

Handler package's main responsibility is to receive incoming requests and generate a proper response after running the business logic.

### Service

Is the hearth of our server and is used to execute business logic. It can call repository package to retrieve or store data in database, cache, or message broker.

### Repository

Repository package is a wrapper for outside services. It isolates the underlying techologies so they can be easily replaced.

### API

Responsible for initializing the app and starting the server. It is the glue that holds all the handlers together.

## PKG

This packages are meant to be shared between projects. Included in this project are helpers for establishing connections.

