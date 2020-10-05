# Using Wire

## Preliminary
Before getting into wiring up applications let's get into the right mindset.

### Dependency injection
Wire assumes that the code is structured with dependency injection in mind. Generally this means types come with 
constructor functions that lets users provide dependencies at initialization time instead of having logic that 
constructing them internally. 

Usually such dependencies are defined in terms of an [interface instead of a value type](https://medium.com/@cep21/what-accept-interfaces-return-structs-means-in-go-2fe879e25ee8)

```go
package example

type (
    Service struct {
        repo Repository
    }

    Repository interface{
        // <snip>
    }
)

func NewService(repo Repository) *Service { 
    return &Service{
    	repo: repo,
    }
}
```
There are many benefits to doing this but to mention a few; it will ease the job of testing as it makes injecting mocks 
or fakes trivial, it helps decouple responsibilities and generally lends its self well to [SOLID](https://dave.cheney.net/2016/08/20/solid-go-design) code.  

### The lifecycle of values 
The lifecycle of values as they are created in an application is important to understand to keep applications lean 
performance wise, and maintainable (aka sustainable). A particular easy to understand model is based on a concept where
all types are categorised as **Services** or **data transfer objects**. 
Services (sub-categorised by concepts such as gateways, repositories, handlers etc) are rich with domain-logic and 
follow the lifecycle of the application, whereas DTOs typically are flat structures without domain logic that exist 
only in the context of a given request or command.

A consequence of this philosophy is that we need only worry about the construction of **service** types.

### Wiring up an application
When wiring up our application, we need to construct (by calling the constructors) all the _services_ that make up the 
application. If we subscribe to the above philosophy we will only need to do it once as our services will only ever be
constructed once. If we do this carefully there is no reason why we can't do it by hand. **No reason for fancy frameworks**.

## Wire concepts
Wire provides a few helper types together with a pattern that will help structure the application bootstrapping logic.

### Factories
A factory is an argument free function / method that provides a new instance of a value

```go
factory := func() interface{} {
    return foobar.NewClient(/* TODO */)
}
``` 

For each of our _services_ we will create a factory and register it with a repository that can invoke it on demand. 

### Repositories
A repository holds a collection of factories and manages the invocation as needed. In particular we may want a given
service to be a **singleton** or create a new instance whenever we request a reference to a service. 

```go
type Repository interface {
    Register(name string, factory FactoryMethod) 
    Get(name string) interface{}
}
```

Usually we will only need a single repository in our application.

### Structuring application bootstrapping

Bootstrapping will be based around a basic `Application` type that embeds the repository and possibly hold other data 
such as configuration.

```go
package main

type Application struct {
    // embed a repository that will provide access to our service type
    wire.AbstractServiceRepository

    // a context that spans the lifecycle of our application
    ctx context.Context
}
```

The actual "work" lies in creating a constructor that initializes the `application` by registering factories and using them.

```go
package main

func NewApplication() (*Application, context.CancelFunc) {
    ctx, cancelFunc := context.Background()
    
    app := &Application{
        AbstractServiceRepository: wire.SingletonProvider(),
        ctx:                       ctx,
    }
	
    // ... <snip>

    app.Register("foobarclient", func() interface{} {
        return foobar.NewClient(app.Connection())
    })

    app.Register("service", func() interface{} {
        return user.NewService(app.FoobarClient())
    })
    
    return app, cancelFunc
}
```
As a general pattern, instead of accessing the repositories `Get` method directly, factory methods may access the 
registered instances via typed helper methods.

```go
package main

func (a *Application) Context() context.Context {
	return a.ctx
}

func (a *Application) Connection() *grpc.ClientConn {
	return a.Get("connect").(*grpc.ClientConn)
}

func (a *Application) FoobarClient() *foobar.Client {
	return a.Get("omniclient").(omnipb.PercipientClient)
}

func (a *Application) UserService() *user.Service {
	return a.Get("service").(*user.Service)
}

```
The `main` function of the application is now reduced to calling the constructur of the _Application_ and instructing 
any rpc handlers to bind the network interface or UIs to launch etc.  

```go
package main 
func main() {
	app, closer := NewApplication()
	defer closer()
	// ... <snip>
}

```
