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

**IMPORTANT** By default, Heatify expects these databases to be named `gitmonitor-repos-dev` and `gitmonitor-users-dev`. 
Create these databases on Cloudant using the 'Create Database' button on the web UI within your account. 
The names for each database can be changed, but the change must also be reflected in the Heatify source code by editing the `COMMITS_DB` and `USERS_DB` variables within `model/RepoCommits.go` and `model/UserCommits.go`. 

It is a good idea to keep databases used by a production copy of Heatify separate from a local/development version, since the dates listed on commits may differ from machine to machine depending on the OS and timezone.

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

Notes to IBM: the supervisord configuration used by Heatify on the VM located at `heatify.rtp.raleigh.ibm.com` is as follows:
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

### The `model` Package
Here you will find some Go structs which model the entities used in Heatify. Entities are Users, Squads, RepoCommits, and UserCommits. The sturcts also contain procedures to populate themselves from local JSON files or from data in Cloudant

Each User has a username, and a list of git aliases they go by (for example 'Brandon Schurman <schurman@ca.ibm.com>').

Each Squad has a squad name, a profile picture, a list of squad members (referenced by a User's username) and a list of git repositories that the Squad owns. 

_Users and Squads data are loaded from local JSON files found in the `config/` directory of Heatify._

RepoCommits are keyed by a YYYY/MM/DD date (note there is no time, just the day of the month). Each RepoCommit stores the number of commits made on each day to each repository. This is basically a compressed representation of the git commit log on a repository.

UserCommits are similar to RepoCommits, however they store the number of commits made on each day to a given repository by a given user. 

_RepoCommits and UserCommits are the structures which are kept in sync with Cloudant._

### The `gitutil` Package
In this package you will find code that handles git repository processing. Code in this package is capable of cloning new git repositories, updating git repos via `git pull`, and keeping new commits in sync with Cloudant. 
An important file to note is `gitutil/updater.go`. Here you will find the Code which syncs commits with Cloudant. The update algorithm works roughly as follows: 
```
For each repo in the CLONES_DIR
    Download RepoCommits from Cloudant
    Run 'git log' on the local copy of the repo 
    If there are RepoCommits in the local copy that are not in Cloudant
        Send the new commits to Cloudant
    Download UserCommits from Cloudant
    Run 'git log' on the local copy of the repo
    If there are UserCommits in the local copy that are not in Cloudant
        Send the new commits to Cloudant
```
This update algorithm is run every `UPDATE_INTERVAL` which is set to 6 hours by default. 

### The `cadb` Package

This package contains a simple 'driver' for Cloudant. It abstracts methods for reading and writing data to Cloudant. These methods are usually called by code in the `model` package. 

### Other Components

The UI code is found in the `views/` directory. CSS, JavaScript and libraries like [Materialize](http://materializecss.com/) which are used for the UI are found in the `static/` directory.



## Future Improvements and Current Issues

- Current versions of Heatify have been known to have stability issues. After the app has been running for hours or days, the app might crash while synchronizing git repositories with Cloudant. Heatify should be run with a process monitor such as supervisord to auto-restart the app if it crashes.
- When using Heatify on a VM with supervisord, there can be issues with cloning new repositories. The issue is due to `git clone` not running and exiting with status 128. 
- The Users feature is not fully implemented. User entities are used on the Squad page, however they should also be able to be viewed individually in separate User pages in the Heatify web UI.
- The git update loop should be run on a different process, so that the UI and data-collector components of Heatify are decoupled. This allows for a much more scalable and distributed architecture.

## Contributing
The Heatify project started at IBM and was originally sourced on IBM DevOps Services. The code has now been forked to a public open source version, found at https://github.com/brandonSc/heatify. Pull requests and issues are welcome at this location. Feel free to get in touch with me (Brandon) directly at schurman93@gmail.com with questions or suggestions. 
