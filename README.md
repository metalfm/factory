# Factory gen

Factory gen or fcgen is a library that generates code for your Go structures.

It can easily generate complex objects by using this, and easily maintain those objects generators.

## Motivation

When writing code, we often use Go structures to store various data, such as various business entities. Structures can be large and small, passing between different parts of the business logic.

Almost always, we want to write tests and test our business process, make sure that everything works correctly (in fact, prove that it does not work incorrectly).

For example, a user leaves an order, we want to test that the order is valid, saved to the database or vice versa. The problem is that first you need to prepare a structure with an order, and it can contain quite a lot of fields or links to other structures.

Code generation and this library come to the rescue, which will help write factories preparing a structure with a pre-defined set of data with the ability to easily change the data.

## Install
```
go install github.com/metalfm/factory/fcgen@latest
```

## Okey, show me the code

This is just an example, for simplicity we will describe everything in one file.
Let's ask the factory to build us a `User` with the required values:

```go
package test

import (
    "fmt"
    "testing"

    "fc_test" // the generated factory code will be in this package
)

//go:generate go run metalfm/factory/fcgen -type=User

type User struct {
    ID       int
    Name     string
    Location string
}

func TestUser(t *testing.T) {
    factoryUser := fc_test.NewFactoryUser(User{
            Location: "Cyprus", // this is default value for Location
        }).
        SeqInt(func(e *User, n int) {
            e.ID = n
        }).
        OnName(func(e *User) {
            e.Name = fmt.Sprintf("user-%d", e.ID)
        })

    for range 3 {
        fmt.Printf("%+v", factoryUser.MustBuild())
    }

    // let's try to customize data
    factoryUser.Name("name").Location("Lissabon")

    fmt.Printf("%+v", factoryUser.MustBuild())
    // Output:
    // {ID:1 Name:user-1 Location:Cyprus}
    // {ID:2 Name:user-2 Location:Cyprus}
    // {ID:3 Name:user-3 Location:Cyprus}
    // {ID:4 Name:_name_ Location:Lissabon}
}

```

In terms of use, it looks simple, doesn't it? All the magic happens after launch `//go:generate`, we call a `fcgen` that analyzes the AST tree with the code, generates simple setters and setter functions for the `User` structure.

You can take a look at the generated code [here](https://github.com/metalfm/factory/blob/main/internal/tests/entity/fc/user_gen.go), it is very easy to understand what it does.

### Possibilities

1. Generation of factories for structures of any complexity
2. Support for all basic, composite and crazy data types, for example — `**func(a, b map[string]func(a, b time.Time, u **User, c map[any]any)) [5]**int`
3. Support generics
4. Support embed basic types or structs
5. Strong typing for all fields, no reflection and runtime type assertions
6. Sub factories work from the box 
7. Setter functions are executed when a new instance is created. Convenient if you need to generate a new uuid each time and use the current date and time


## Okey, show me more

### Random data
Feel free to use any library to generate random data if you need it:
```go
package test

import (
    "testing"
    
    "github.com/Pallinder/go-randomdata"
    "fc_test"
)

//go:generate go run metalfm/factory/fcgen -type=User

type User struct {
    ID       int
    Name     string
    Email    string
}

func TestUser(t *testing.T) {
    factoryUser := fc_test.NewFactoryUser(User{}).
        OnName(func(e *User) {
            e.Name = randomdata.FullName(randomdata.RandomGender)
        }).
        OnEmail(func(e *User) {
            e.Name = randomdata.Email()
        })
}
```

### UUID & CreatedAt
Sometimes we want to generate a structure with valid but different fields, such as UUID or CreatedAt:

```go
package test

import (
    "testing"
    "time"

    "github.com/google/uuid"

    "fc_test"
)

//go:generate go run metalfm/factory/fcgen -type=User

type User struct {
    ID        uuid.UUID
    CreatedAt time.Time
}

func TestUser(t *testing.T) {
    factoryUser := fc_test.NewFactoryUser(User{}).
        OnID(func(e *User) {
            e.ID = uuid.New()
        }).
        OnCreatedAt(func(e *User) {
            e.CreatedAt = time.Now()
        })
}
```

### Sub factory

```go
package test

import (
    "testing"

    "fc_test"
)

//go:generate go run metalfm/factory/fcgen -type=User,Order

type User struct {
    Name string
}

type Order struct {
    ID   int
    User User
}

func TestOrder(t *testing.T) {
    factoryUser := fc_test.NewFactoryUser(User{})
    factoryOrder := fc_test.NewFactoryOrder(Order{}).
        OnUser(func(e *User) {
            e.User = factoryUser.MusBuild()
        })

    // You can also change the default values
    factoryUser.Name("test")
    factoryOrder.User(factoryUser.MustBuild())
}
```

### Save to database
Sometimes we want not only to get a structure with the necessary data, but also to save it to a database. Or any other storage.

```go
package test

import (
    "context"
    "github.com/jmoiron/sqlx"
    "testing"

    "fc_test"
)

//go:generate go run metalfm/factory/fcgen -type=User

type User struct {
    ID   int
    Name string
}

type ctxKey int

func TestUser(t *testing.T) {
    var db *sqlx.DB // for example, here, we have a db instance

    ctx := context.Background()
    
    // let's put our db instance to context
    ctx = context.WithValue(ctx, ctxKey, db)

    factoryUser := fc_test.NewFactoryUser(User{}).
        OnCreate(func(ctx context.Context, e *User) error {
            db, ok := ctx.Value(ctxKey).(*sqlx.DB)
            if !ok {
                return nil
            }
            
            // here you have a database connection and structure, nothing complicated to save it
            // db — *sqlx.DB
            // e  — User struct
            // you can also extract the incremented ID from database and set it to User
        })

    // we created and got a user that is in the database
    user := factoryUser.MustBuildCtx(ctx)
}
```

### Generics
See real world example [here](https://github.com/metalfm/factory/blob/main/internal/tests/entity/document_test.go).

### Setter priority
You can always override the default values to get the desired data set for your test:

```go
package test

import (
    "testing"
    "fmt"
    "fc_test"
)

//go:generate go run metalfm/factory/fcgen -type=User

type User struct {
    Name string
}

func TestUser(t *testing.T) {
    factoryUser := fc_test.NewFactoryUser(User{
        Name: "name_0",   
    }).
    OnName(func(e *User) {
        e.Name = "name_1"
    }).
    Name("name_2")

    u := factoryUser.MusBuild()
    
    fmt.Println(u.Name)
    // Output:
    // name_2
}
```

### More examples

Try to find some example code in [tests](https://github.com/metalfm/factory/tree/main/internal/tests/entity).

## Known limitations

Dotted imports are not supported. Just don't use them.

```go
package test

import (
    "time"
    . "time"
)

type S struct {
    CreatedAt0 Time      // do not this
    CreatedAt1 time.Time // do this
}
```

Anonymous interfaces with methods are not supported, such fields will simply not be used in generation. 

```go
package test

type Handler interface { Handle() }

type S struct {
    Handler0 interface{ Handle() } // do not this
    Handler1 Handler               // do this
}
```

Generic names with lowercase letters are not supported, just name them with capital letters.

```go
package test

type S0[t any] struct { // do not this
    Field t
}

type S1[T any] struct { // do this
    Field T
}
```

## License

Fcgen is licensed under the MIT License. See [LICENSE](https://github.com/metalfm/factory/blob/main/LICENSE) for more information.