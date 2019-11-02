[![Build Status](https://travis-ci.org/emilhein/go-aws-webserver.svg?branch=master)](https://travis-ci.org/emilhein/go-aws-webserver)


Example [here](https://github.com/emilhein/go-heroku-server)
# go-aws-webserver

A small webserver written in Go, to perform simple tasks on your AWS ressources

### How to use

NOTE: Remeber to set aws credentials when you deploy this server

Make sure you have AWS credentials available on your computer:

Windows: C:\Users\{USERNAME}\.aws\credentials

```
package main

import (
	"github.com/emilhein/go-aws-webserver/webserver"
)

func main() {
	webserver.Start()
}

```

Now you have a webserver with some already defined endpoints for your AWS ressources.

For now the most usefull endpoint is:

POST: localhost:3000/getS3files

A POST body request in postman could look like this:

```
{
	"bucket" : "YOUR_BUCKET",
	"filepaths" : ["json_file_1", "json_file_2"]
}

```
