# Framework (fwork)

![Main branch](https://github.com/ravelo-systematic-solutions/fwork/actions/workflows/go.yml/badge.svg)

## Motivation

This project is implemented to simplify and fasten the implementation of
testable/secure cloud APIs using a very specific set of conventions.
Feel free to use it, contribute and/or submit suggestions for design changes.

## Resource convention

A `resource` is a convention for referencing atomic entities which is represented by a url
and managed by HTTP verbs. Currently, the supported verbs are:

* GET: gets a list of resources with GET parameters for filtering capabilities
* POST: creates a single resource
* PUT: updates a single resource including all fields
* PATCH: updates a single resource updating only the values passed
* DELETE: deletes a single resource

## Resource formats

Currently, we only support JSON.

### Example

A common resource in applications are users. We could set a resource identify
by the "/users" (plural) url, managed by HTTP verbs (see above) and is expected
to only deal with one single instance at a time.

## Usage examples

### Simple Hello World

Create a new directory, create a `go module`, `myapp.go`, `controllers.go`
and `controllers_test.go`

```shell
mkdir myapp
cd myapp
touch myapp.go controllers.go controllers_test.go
go mod init github.com/org/myapp
```

Install this tool as a dependency

```shell
go get -u github.com/ravelo-systematic-solutions/fwork
```

Create the TLS subject

```go
	subject := api.CertificateSubject{
		SerialNumber:  int64(time.Now().Year()),
		Organization:  []string{"Company Inc."},
		Country:       []string{"CA"},
		Province:      []string{"BC"},
		Locality:      []string{"Vancouver"},
		StreetAddress: []string{"123 Some Street"},
		PostalCode:    []string{"A0A 0A0"},
		CertNotBefore: time.Now(),
		CertNotAfter:  time.Now().AddDate(10, 0, 0),
	}
```

Create Private key programmatically

```go
	privateKey, err := api.GeneratePrivateKey(4096)

    if err != nil {
        // handle error
    }
```

Configure the resource

```go
type UserDt struct {
	Id       string `json:"id"`
	FullName string `json:"full_name"`
}

func Get(scope api.Scope) {
	scope.JsonRes(http.StatusAccepted, UserDt{
		Id:       "1234",
		FullName: "Art Doe",
	})
}

var User = &user{
	api.NewResource("/users", api.Endpoints{
		Get: Get,
	}),
}

type user struct {
	api.Resource
}
```

Configure the application & run the API

```go
	conf := api.Config{
		Service: api.Service{
			Id:       "myapp",
			Name:     "myapp",
			Internal: ":50000",
			External: "https://localhost:50000",
		},
	}
    server, err := api.NewEngine(subject, privateKey, conf)
    server.Controller(controllers.User)
    server.Run()
```

Last but not least, test the controller

```go
//given
d := UserDt{
    Id:       "1234",
    FullName: "Art Doe",
}

//when
sut := api.NewTestScope(
    http.MethodGet,
    User.Url(),
    nil,
    User,
)

//then
if err := sut.IsStatus(http.StatusAccepted); err != nil {
    t.Errorf(
        "IsStatus(), %s",
        err.Error(),
    )
}

if err := sut.IsJsonRes(d); err != nil {
    t.Errorf(
        "IsStatus(), %s",
        err.Error(),
    )
}
```

Now you should be able to go to `https://localhost:50000/users`
and see the following response:

```json
{"id":"1234","full_name":"Art Doe"}
```

Here's what the files look like

#### server.go

```go
package main

import (
	"github.com/ravelo-systematic-solutions/fwork/api"
	"log"
	"time"
)

func main() {

	subject := api.CertificateSubject{
		SerialNumber:  int64(time.Now().Year()),
		Organization:  []string{"Company Inc."},
		Country:       []string{"CA"},
		Province:      []string{"BC"},
		Locality:      []string{"Vancouver"},
		StreetAddress: []string{"123 Some Street"},
		PostalCode:    []string{"A0A 0A0"},
		CertNotBefore: time.Now(),
		CertNotAfter:  time.Now().AddDate(10, 0, 0),
	}

	privateKey, err := api.GeneratePrivateKey(4096)
	if err != nil {
		log.Printf("failed to generate private key: %v", err)
		return
	}

	conf := api.Config{
		Service: api.Service{
			Id:       "myapp",
			Name:     "myapp",
			Internal: ":50000",
			External: "https://admin.ravelo.local:50000",
		},
	}

	server, err := api.NewEngine(subject, privateKey, conf)
	if err != nil {
		log.Printf("failed to instantiate service: %v", err)
		return
	}

	server.Controller(User)
	server.Run()

}

```

#### controllers.go

```go
package main

import (
	"github.com/ravelo-systematic-solutions/fwork/api"
	"net/http"
)

type UserDt struct {
	Id       string `json:"id"`
	FullName string `json:"full_name"`
}

func Get(scope api.Scope) {
	scope.JsonRes(http.StatusAccepted, UserDt{
		Id:       "1234",
		FullName: "Art Doe",
	})
}

var User = &user{
	api.NewResource("/users", api.Endpoints{
		Get: Get,
	}),
}

type user struct {
	api.Resource
}

```

#### controllers_test.go

```go
package main

import (
	"github.com/ravelo-systematic-solutions/fwork/api"
	"net/http"
	"testing"
)

func TestUser_List(t *testing.T) {
	//given
	d := UserDt{
		Id:       "1234",
		FullName: "Art Doe",
	}

	//when
	sut := api.NewTestScope(
		http.MethodGet,
		User.Url(),
		nil,
		User,
	)

	//then
	if err := sut.IsStatus(http.StatusAccepted); err != nil {
		t.Errorf(
			"IsStatus(), %s",
			err.Error(),
		)
	}

	if err := sut.IsJsonRes(d); err != nil {
		t.Errorf(
			"IsStatus(), %s",
			err.Error(),
		)
	}
}
```
