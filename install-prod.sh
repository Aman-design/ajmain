#!/usr/bin/env sh
set -eu

# Listmonk production setup using `docker-compose`.
# See https://listmonk.app/docs/installation/ for detailed installation steps.

printf '\n'

RED="$(tput setaf 1 2>/dev/null || printf '')"
BLUE="$(tput setaf 4 2>/dev/null || printf '')"
GREEN="$(tput setaf 2 2>/dev/null || printf '')"
NO_COLOR="$(tput sgr0 2>/dev/null || printf '')"

info() {
  printf '%s\n' "${BLUE}> ${NO_COLOR} $*"
}

error() {
  printf '%s\n' "${RED}x $*${NO_COLOR}" >&2
}

completed() {
  printf '%s\n' "${GREEN}✓ ${NO_COLOR} $*"
}

exists() {
  command -v "$1" 1>/dev/null 2>&1
}

check_dependencies() {
	if ! exists curl; then
		error "curl is not installed."
		exit 1
	fi

	if ! exists docker; then
		error "docker is not installed."
		exit 1
	fi

	if ! exists docker-compose; then
		error "docker-compose is not installed."
		exit 1
	fi
}

download() {
	curl --fail --silent --location --output "$2" "$1"
}

is_healthy() {
	info "waiting for db container to be up. retrying in 3s"
	health_status="$(docker inspect -f "{{.State.Health.Status}}" "$1")"
	if [ "$health_status" = "healthy" ]; then
		return 0
	else
		return 1
	fi
}

is_running() {
	info "checking if "$1" is running"
	status="$(docker inspect -f "{{.State.Status}}" "$1")"
	if [ "$status" = "running" ]; then
		return 0
	else
		return 1
	fi
}

generate_password(){
	echo $(tr -dc A-Za-z0-9 </dev/urandom | head -c 13 ; echo '')
}

get_config() {
	info "fetching config.toml from listmonk repo"
	download https://raw.githubusercontent.com/knadh/listmonk/master/config.toml.sample config.toml
}

get_containers() {
	info "fetching docker-compose.yml from listmonk repo"
	download https://raw.githubusercontent.com/knadh/listmonk/master/docker-compose.yml docker-compose.yml
}

modify_config(){
	info "generating a random password"
	db_password=$(generate_password)

	info "modifying config.toml"
	# Replace `db.host=localhost` with `db.host=db` in config file.
	sed -i "s/host = \"localhost\"/host = \"listmonk_db\"/g" config.toml
	# Replace `db.password=listmonk` with `db.password={{db_password}}` in config file.
	# Note that `password` is wrapped with `\b`. This ensures that `admin_password` doesn't match this pattern instead.
	sed -i "s/\bpassword\b = \"listmonk\"/password = \"$db_password\"/g" config.toml
	# Replace `app.address=localhost:9000` with `app.address=0.0.0.0:9000` in config file.
	sed -i "s/address = \"localhost:9000\"/address = \"0.0.0.0:9000\"/g" config.toml

	info "modifying docker-compose.yml"
	sed -i "s/POSTGRES_PASSWORD=listmonk/POSTGRES_PASSWORD=$db_password/g" docker-compose.yml
}

run_migrations(){
	info "running migrations"
	docker-compose up -d db
	while ! is_healthy listmonk_db; do sleep 3; done
	docker-compose run --rm app ./listmonk --install
}

start_services(){
	info "starting app"
	docker-compose up -d app db
}

show_output(){
	info "finishing setup"
	sleep 3

	if is_running listmonk_db && is_running listmonk_app
	then completed "Listmonk is now up and running. Visit http://localhost:9000 in your browser."
	else
		error "error running containers. something went wrong."
	fi
}


check_dependencies
get_config
get_containers
modify_config
run_migrations
start_services
show_output
