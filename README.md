# Project generator for Go

generate project [standard-like](https://github.com/golang-standards/project-layout) structure, sample Makefile and
Dockerfile.

## Install

````shell
go get -u https://github.com/alapierre/gogen
````

## Usage

Generate project with Dockerfile:

````shell
gogen -m test/test-service -d
````