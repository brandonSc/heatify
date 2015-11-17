Git-Monitoring 
==============

Running Locally 
---------------
You will need Cloudant DBaaS configured within your `VCAP_SERVICES` environment. This means you will need to either deploy the application as Cloud Foundry droplet, deploy it within a container bound to a Cloud Foundry app to access it's services, or copy-and-paste the contents of `VCAP_SERVICES` into an environment variable if testing on your local machine. 

youniteyouniteFirst download and install the latest Golang distribution [here](https://golang.org/doc/install#tarball).
Note: if you are using ubuntu, do not install Go through apt-get. The version in the apt repository is out of date and will not work JazzHub URLs.
If you are new to go, you will have to setup your workspace first. See [How to Write Go Code](https://golang.org/doc/code.html) for official documentation.
Ensure your `$GOPATH` environment variable should be set to your desired workspace.
Now, install the Git-Monitor project by running the following:
```
go get hub.jazz.net/git/schurman93/Git-Monitor
```
If the command fails, ensure you are using a version of Go that supports JazzHub URLs. At this time, the latest version 1.4.2 worked with Jazz.
Navigate to the Git-Monitor source code
```
$ cd $GOPATH/src/hub.jazz.net/git/schurman93/Git-Monitor
```
You are ready to build the project. Execute the following:
```
$ go build 
```
Then run with 
```
$ ./Git-Monitor
```
(or if you have the $GOPATH/bin configured globally, you can run in any directory by invoking `$ Git-Monitor`).

Navigate to http://localhost:8080 in your browser. 
