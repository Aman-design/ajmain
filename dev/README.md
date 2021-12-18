# Docker suite for development

**NOTE**: This exists only for local development. If you're interested in using Docker for a production setup, visit the [docs](https://listmonk.app/docs/installation/#docker) instead.

### Objective

The purpose of this docker suite for local development is to isolate all the dev dependencies in a docker environment. The containers have a host volume mounted inside for the entire app directory. This helps us to not do a full `docker build` for every single local change, only restarting the docker environment is enough.

## Setting up a dev suite

To spin up a local suite of

- PostgreSQL
- Mailhog
- Node.js frontend app
- Golang backend app

### Setup DB

```bash
make init-dev-docker
```

### Start frontend and backend apps

```bash
make dev-docker
```

Visit `http://localhost:8080` on your browser.

### Tear down

This will tear down all the data, including DB.

```bash
make rm-dev-docker
```

### See local changes in action

- Backend: Anytime you do a change to the Go app, it needs to be compiled. Just run `make dev-docker` again and that should automatically handle it for you.
- Frontend: Anytime you change the frontend code, you don't need to do anything. Since `yarn` is watching for all the changes and we have mounted the code inside the docker container, `yarn` server automatically restarts.
