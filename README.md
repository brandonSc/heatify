# Heatify

Heatify tracks and measures the commits made to git repositories hosted on Github, GHE, GitLab, and IBM DevOps Services. Heatmaps showing daily activity are rendered on individual projects, users, as well as grouped into Squads to conveniently analyze team performance. 

## Running Instructions

You will need [Go](https://golang.org/) installed as well as an [IBM Cloudant](https://cloudant.com/) account. 
IBM Cloudant can be used for free under light usage, and can be obtained as a service on [IBM Bluemix](https://console.ng.bluemix.net). 
See [here](https://golang.org/doc/install#tarball) for instructions on downloading and installing Go.
First obtain a copy of the Heatify source by running `go get`:

```
$ go get github.ibm.com/kdaihee/heatify
```

Navigate to the source directory

```
$ cd $GOPATH/src/github.ibm.com/kdaihee/heatify
```

Then build the executable:

```
$ go build 
```

### Setting Up Git

Heatify uses a local git account and credential to clone and update repositories. 
Ensure that you're able to clone or update repositories from public Github, GHE, IDS, and GitLab without having to manaully enter your username and password. 
For use with HTTPS, you can configure git to persist your credentials using the following command:

```
$ git config --global credential.helper store --file ~/.git-credentials
```

for more info on git credentials, see [here](https://git-scm.com/docs/git-credential-store). You must ensure that the credentials for all git accounts are stored.

Heatify also uses SSH keys to clone and update repositories over SSH.
Generate an ssh-key and copy and add it into each git account online. 
For example, see Github's documentation on using SSH keys [here](https://help.github.com/articles/generating-an-ssh-key/).
Before running the app, ensure the ssh-key is added to the ssh agent. It can be added with the following commands on Ubuntu:

```
$ eval `ssh-agent -s`
$ ssh-add ~/.ssh/id_rsa
```

It is important that the ssh key is permanently stored when running in a production VM. Issues can arise if the ssh-agent loses access to the key.


### Setting Up Cloudant

In your Cloudant account, you should have a database for RepoCommits and a database for UserCommits.
Heatify has a data-collector process which constantly keeps collections in sync of commits-per-day on each repository, and commits-per-day made by users. 
By default, Heatify expects these databases to be named `gitmonitor-repos-dev` and `gitmonitor-users-dev`. 
Create these databases on Cloudant using the 'Create Database' button on the web UI within your account. 
The names for each database can be changed, but the change must also be reflected in the Heatify source code by editing the `COMMITS_DB` and `USERS_DB` variables within `model/RepoCommits.go` and `model/UserCommits.go`. 

Before running Heatify, the Cloudant credentials must be exported to the process environment in a variable named `VCAP_SERVICES`. 
If using Bluemix, a JSON object with the credentials can be found on the "Service Credentials" page of your Cloudant service. 
For example, adding these credentials can be accomplished as follows:

```
$ export VCAP_SERVICES='{
  "credentials": {
    "username": "e469e71e-caa1-****-****-********-bluemix",
    "password": "41ea1d3************************************5dd3c1d2be83971",
    "host": "e469e71e-caa1-****-****-*******-bluemix.cloudant.com",
    "port": 443,
    "url": "https://e469e71e-caa1-****-****-********-bluemix:41ea1d3******************************5dd3c1d2be83971@e469e71e-caa1-****-****-**********-bluemix.cloudant.com"
  }
}'
```

### Running Locally

Run the executable

```
$ ./heatify
```

and navigate to `http://localhost:5050` in your browser. 


### Deployment in a Production Environment

Heatify should be deployed in a VM, or a Container with access to a Volume or block storage device, so that the `clones`  directory can be accessed. 
The clones directory is by default located at `.clones/` within the same directoy as the source code. You can change it's location by editing the `CLONES_DIR` variable within `gitutil/updateLoop.go`.

You can now simply run the executable. For example, run in background and redirect all logs to a file with the following command:
```
$ ./heatify > ~/heatify.log 2>&1 & 
```

A better and safer approach would be to use supervisord to monitor and auto-restart the app. See [here](https://serversforhackers.com/monitoring-processes-with-supervisord) for instructions on configuring and running supervisord. 
In summary, use `supervisorctl start` and `supervisorctl stop` to start and stop the app using supervisosrd.

Notes to IBM: the supervisord config I used for Heatify for the VM located at `heatify.rtp.raleigh.ibm.com` as follows:
```
[program:heatify]
command=/home/ibmadmin/gocode/src/github.ibm.com/kdaihee/heatify.git/heatify.git
directory=/home/ibmadmin/gocode/src/github.ibm.com/kdaihee/heatify.git
autostart=true
autorestart=true
startretries=3
stderr_logfile=/var/log/heatify/heatify.err.log
stdout_logfile=/var/log/heatify/heatify.out.log
user=ibmadmin
environment=VCAP_SERVICES='{ "credentials": { "username": "e469e71e-caa1-****-****-********-bluemix", "password": "41ea1d3************************************5dd3c1d2be83971", "host": "e469e71e-caa1-****-****-*******-bluemix.cloudant.com", "port": 443, "url": "https://e469e71e-caa1-****-****-********-bluemix:41ea1d3******************************5dd3c1d2be83971@e469e71e-caa1-****-****-**********-bluemix.cloudant.com" } }'

```
Note the location of the log files are in `/var/log/heatify`.

## Architecture 

Heatify is written in Go is broken down into several packages, the key packages are `gitutil`, `model`, `route` and `cadb`. 

### The `gitutil` package

## Future Improvements and Current Issues

- Current versions of Heatify have been known to have stability issues. After the app has been running for hours or days, the app might crash while synchronizing git repositories with Cloudant. Heatify should be run with a process monitor such as supervisord to auto-restart the app if it crashes.
- Issues 
