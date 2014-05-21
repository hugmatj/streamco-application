# StreamCo Application

TODO: write something nice here.

## Developer setup

vagrant up
vagrant ssh
sw (start work: change directory to /go/src/garethstokes/streamco-application)
./run_webserver.sh (starts the webserver on port 3000)

inside of vagrant there is a nginx process running on port 80 that is a proxy to port 3000. 
outside of vagrant, we can hit port 8000 which is being forwarded to port 80 inside. 

ie: once the go webserver is running on port 3000 inside vagrant, 
you can run "curl localhost:8000" on your host to get a hello world

## Deploy to production aka Heroku

heroku create -b https://github.com/kr/heroku-buildpack-go.git
heroku push heroku master
