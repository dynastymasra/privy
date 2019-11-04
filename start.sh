#!/bin/bash -e

case $1 in
  "run")
    docker-compose build privy
    docker-compose run --rm privy migrate:run
    docker-compose run --rm privy migrate:seed
    docker-compose up privy
    ;;
  "up")
    docker-compose up privy
    ;;
  *)
    echo "usage: $0 [run|up]"
    exit 1
    ;;
esac