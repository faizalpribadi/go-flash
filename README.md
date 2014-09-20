go-flash
========

short tutorial flash message with go-lang


run command
===========
`go run main.go flash.go`


open the browser 
================
* `http://localhost:2222`
* `http://localhost:2222/setFlashMessage`
* `http://localhost:2222/getFlashMessage`


request cookie with curl command 
================================

* `curl -i --cookie-jar sfm localhost:2222/setFlashMessage`
* `curl -i --cookie-jar cs --cookie sfm localhost:2222/getFlashMessage`
