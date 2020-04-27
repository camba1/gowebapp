/*
	Main: Entry point package to a small application that allows users to list, create and modify files in a fixed, predetermined directory.
	The application is composed of three parts:
	- The main package that handles all the http requests
	- The fileManager package is used to manage all the file operations (save, edit, etc)
	- The front end is handled using html template files and rendering them using the html/template package

	The application can be compiled if you have go installed or using docker (or docker-compose)
*/
package main
