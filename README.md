Weather Service
================

This weather service accepts a latitude and longitude and returns the short version of the current weather at that location.


Table of Contents
-----------------

* [Getting Started](#getting-started)
* [Prerequisites](#prerequisites)
* [Building the Project](#building-the-project)
* [Running the Project](#running-the-project)
* [Testing the Project](#testing-the-project)


Getting Started
---------------

These instructions will get you a copy of the project up and running on your local machine.


Prerequisites
------------

* Go (version specified in `go.mod`)

Building the Project
---------------------

1. Clone the repository using the following command:
```bash
git clone https://github.com/mnmeyers/weather-service.git
```
2. Navigate to the project directory:
```bash
cd weather-service
```
3. Build the project using the following command:
```bash
go build main.go
```
This will create a binary file named `weather-service` in the project directory.


Running the Project
---------------------

Run the project using the following command:
```bash
./weather-service
```
This will start the project and make it available at port 8080.

Navigate to `http://localhost:8080/weather/37.7749/-122.4194` in your browser to see the current weather in San Francisco, CA.

Testing the Project
---------------------
Run the tests using the following command:
```bash
go test ./...
```
This will run all the tests in the project and report any failures.