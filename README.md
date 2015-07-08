Git-Monitoring 
==============

Running Locally 
---------------
First install golang. 
On ubuntu, run: 
```
sudo apt-get install golang
```
or find the download for your OS [here](https://golang.org/doc/install#tarball).
If you are new to go, you will have to setup your workspace first. See [How to Write Go Code](https://golang.org/doc/code.html) for official documentation.
Your `$GOPATH` environment variable should be set to your desired workspace.
If you have not already setup up a directory in your go workspace for jazz-hub, you can do so by running the following (unix only):
```
$ mkdir -p $GOPATH/src/hub.jazz.net/project
```
Then, create a directory for my user
```
$ mkdir $GOPATH/src/hub.jazz.net/project/schurman93
```
Clone the Git-Monitor project into `schurman93` (or just move it there if you have already cloned it)
```
$ git clone https://hub.jazz.net/project/schurman93/Git-Monitor $GOPATH/src/hub.jazz.net/project/schurman93/
```
Now change to the project directory
```
$ cd $GOPATH/src/hub.jazz.net/project/schurman93/Git-Monitor
```
You are ready to build the project. Execute the following:
```
$ go get
$ go build 
```
Then run with 
```
$ ./Git-Monitor
```
Navigate to http://localhost:8080 in your browser
