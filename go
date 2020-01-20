#!/usr/bin/env bash
#########################################################################
# ./go script for Tour - Docker 201 . This script is the entry point for all
# development needs and acts as a facade for gradle and other calls.
# Usage:
#   ./go <command>
#########################################################################
# Global variables
PS_DBNAME=Tours
PS_VERSION=11.5
PS_CONTAINER_NAME=tour-db
PS_USERNAME=touradmin
PS_PASSWORD=touradmin
PS_DATADIR="${PWD}/.postgres/"
#########################################################################
set -e

show_help() {
  echo "Usage: "
  echo "  ./go <command>"
  echo ""
  echo "Known commands are:"
  # Please keep these sorted alphabetically:
  echo "  build         run a gradle build"
  echo "  db-start      start the local PostgreSQL DB"
  echo "  db-stop       stop the PostgreSQL DB."
  echo "  dockerDown    bring down the Tour CRUD infrastructure via docker-compose"
  echo "  dockerUp      bring up the Tour CRUD infrastructure via docker-compose"
  echo "  help          show this help"
  echo "  install       list the dependencies you need to install"
  echo "  run-db        run the app locally with persistent PostgreSQL DB"
  echo "  style         run checkstyle for main and test classes"
  echo "  test          run test"
  echo ""
  echo "You can add commands by editing this script!"
}

runtest() {
  ./gradlew clean test intTest
}

build() {
  ./gradlew clean build
}

style() {
  ./gradlew clean :checkstyleMain :checkstyleTest
}


run() {
  ./gradlew bootRun
}

run-db() {
  db-start
  SPRING_PROFILES_ACTIVE=docker ./gradlew bootRun
}


db-start() {
  echo "------------------------------------------------------------"
  echo "Starting PostgreSQL docker container."
  echo "------------------------------------------------------------"
  docker-compose -f tour-crud-stack.yml up -d postgres
}

db-stop() {
  echo "------------------------------------------------------------"
  echo "Stopping PostgreSQL docker container."
  echo "------------------------------------------------------------"
  docker stop docker_postgres_1
}

install() {
  # FIXME this should be an automatic install section, not docs
  echo "Install AdoptJDK 11"
  echo "Install Docker"
  echo "Install HawkEye"
  echo "Install Talisman"
}

dockerUp() {
  docker-compose -f tour-crud-stack.yml up --detach
}

dockerDown() {
  docker-compose -f tour-crud-stack.yml down
}

command=$1
case ${command} in
build) build ;;
db-start) db-start ;;
db-stop) db-stop ;;
dockerDown) dockerDown ;;
dockerUp) dockerUp ;;
style) style ;;
help) show_help ;;
install) install ;;
run) run ;;
run-db) run-db ;;
test) runtest ;;
*) show_help ;;
esac

exit 0
