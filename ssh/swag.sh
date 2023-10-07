#!/bin/bash

swag i -g init_router.go -dir app/admin/router --instanceName admin --parseDependency -o docs/admin
swag i -g init_router.go -dir app/sequence/router --instanceName sequence --parseDependency -o docs/sequence