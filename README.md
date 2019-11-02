# go-challenges

## Prerequisites

* Go Version 1.13
* Git

## Challenge 1

> Build a command-line application in Go that queries system utilization.

The application, when run, should return system utilization such as CPU utilization, RAM utilization, and Backing Store Utilization in a table format

* This code will be contained within a single main.go file

#### Key Concepts

* You will learn how to create a simple Go application and subsequently compile and run this application.
* You will gain a basic understanding of the os package and how to import packages from the standard library into your Go applications.

## Challenge 2

> You will extend the application developed in the first challenge so that these statistics can be queried via a http request. 

This application will run indefinitely on a machine and expose information on given port and an endpoint such that when you hit http://localhost:9000/stats it will fetch and display the stats for that machine in JSON format.

#### Key Concepts

* You will gain an understanding of how to write a simple HTTP server in Go
* You will learn about how to marshal and unmarshal information into structs

## Challenge 3

> TBD
