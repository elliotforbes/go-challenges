# go-challenges

These challenges have been designed and created to gradually build up your knowledge of Go. Each challenge will ultimately cover more and more advanced concepts within the Go programming language and help you to master the basics of the language so that you can go forth and write your own beautiful Go code.

Each challenge will feature a couple of links to tutorials which should contain everything you need in order to complete the challenge. 

## Prerequisites

* A Text Editor such as Visual Studio Code to write your Go applications in
* Go Version 1.13
* Git installed and a github account

## Challenge 1

> Build a command-line application in Go that queries system utilization.

The application, when run, should return system utilization such as CPU utilization, RAM utilization, and Backing Store Utilization in a table format

* This code will be contained within a single main.go file

#### Key Concepts

* You will learn how to create a simple Go application and subsequently compile and run this application.
* You will gain a basic understanding of the os package and how to import packages from the standard library into your Go applications.

<details>
  <summary>Tutorials</summary>
  
  [Getting Started with Go](https://tutorialedge.net/golang/getting-started-with-go/)
</details>


## Challenge 2

> You will extend the application developed in the first challenge so that these statistics can be queried via a http request. 

This application will run indefinitely on a machine and expose information on given port and an endpoint such that when you hit http://localhost:9000/stats it will fetch and display the stats for that machine in JSON format.

#### Key Concepts

* You will gain an understanding of how to write a simple HTTP server in Go
* You will learn about how to marshal and unmarshal information into structs

<details>
  <summary>Tutorials</summary>
  
  [Creating a RESTful API in Go](https://tutorialedge.net/golang/creating-restful-api-with-golang/)
</details>

## Challenge 3

> TBD

## Challenge 4

> 

## Challenge 5

>

## Challenge 6

>

## Challenge 7

>
