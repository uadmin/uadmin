# uAdmin the Golang Web Framwork

Easy to use, balzing fast and secure.

For Documentation:

- [Read the Docs](https://readthedocs.org/projects/uadmin/)

## Installation

```bash
$ go get github.com/uadmin/uadmin/...
```

To test if your installation is fine, fun the `uadmin` command line:

```bash
$ uadmin
Usage: uadmin COMMAND [-e email] [-d domain]
This tools allows you to publish your project online

Commands:
  publish         This publishes your project online
  prepare         Generates folders and prepares static and templates

Arguments:
  -e, --email     Your email. This is required for you to be able to maintain your project.
  -d, --domain    You can choose your domain name which will customize your URL

Get full documentation online:
https://uadmin.io/docs/
```

## Your First App

Let's build your first app which is a Todo list. First, we will create a folder for your project and prepapre it.

```bash
$ mkdir -p ~/go/src/github.com/your_name/todo
$ cd ~/go/src/github.com/your_name/todo
$ uadmin prepare
[   OK   ]   Created: /home/abdullah/go/src/github.com/uadmin/todo1/models
[   OK   ]   Created: /home/abdullah/go/src/github.com/uadmin/todo1/api
[   OK   ]   Created: /home/abdullah/go/src/github.com/uadmin/todo1/views
[   OK   ]   Created: /home/abdullah/go/src/github.com/uadmin/todo1/media
[   OK   ]   Created: /home/abdullah/go/src/github.com/uadmin/todo1/static
[   OK   ]   Created: /home/abdullah/go/src/github.com/uadmin/todo1/templates
```

Now use your code editor to create `main.go` and put this code inside it.

```golang
package main

import (
	"github.com/uadmin/uadmin"
	"time"
)

type Todo struct {
	uadmin.Model
	Name        string
	Description string `uadmin:"html"`
	TargetDate  time.Time
	Progress    int `uadmin:"progress_bar"`
}

func main() {
	uadmin.Register(Todo{})
	uadmin.StartServer()
}
```

Now to run your code:

```bash
$ go build; ./todo
[   OK   ]   Initializing DB: [9/9]
[   OK   ]   Initializing Languages: [185/185]
[  INFO  ]   Auto generated admin user. Username: admin, Password: admin.
[   OK   ]   Server Started: http://0.0.0.0:8080
         ___       __          _
  __  __/   | ____/ /___ ___  (_)___
 / / / / /| |/ __  / __  __ \/ / __ \
/ /_/ / ___ / /_/ / / / / / / / / / /
\__,_/_/  |_\__,_/_/ /_/ /_/_/_/ /_/
```

# Quick Reference

## Overriding Save Function

```golang
func (m *Model)Save(){
	//business logic
	uadmin.Save(m)
}
```


## Validation

```
func (v Validate) Validate() (ret map[string]string) {
  ret = map[string]string{}
  if v.Name != "test" {
    ret["Name"] = "Error name not found"
  }
  return
}
```
