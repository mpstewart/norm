`norm` is a naïve object relational mapper for SQL databases. It is a toy.

# How to use it

## Declare a model using `norm.Table`

```
type account struct {
	norm.Table `norm:"accounts"`
	Username   string `norm:"username"`
	Email      string `norm:"email"`
}
```

## Register it with `norm.Register[X]()`

```
norm.Register[account]()
```

## Generate queries that can find your model

```
norm.Search[account]{
    "username": "Greg",
    "email":    "greg@example.com",
}.Query()
```

# Should you use this?

Probably not. It's just a toy.

# Will there be a future?

Probably not. Maybe. I might integrate some things with github.com/jmoiron/sqlx
