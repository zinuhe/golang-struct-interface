Interfaces

```go
//GOOD                                                                | //BAD
                                                                      | 
//db.go                                                               | //postgres.go
package db                                                            | package db
                                                                      | type Store interface {
type Store struct {                                                   |   Insert(item interface{}) error
   db *sql.DB                                                         |   Get(id int) error
}                                                                     | }
                                                                      |
// returning concrete type here!                                      | type MyStore struct {
func NewDB() *Store { ... } //func to initialise DB                   |   db *sql.DB
                                                                      | }
func (s *Store) Insert(item interface{}) error { ... } //insert item  |
func (s *Store) Get(id int) error { ... } //get item by id            | func InitDB() Store { ... } //func to initialise DB
                                                                      |
                                                                      | func (s *Store) Insert(item interface{}) error { ... } //insert item
                                                                      | func (s *Store) Get(id int) error { ... } //get item by id
                                                                      |
                                                                      |
//user.go                                                             | //user.go
package user                                                          | package user

type UserStore interface {
   Insert(item interface{}) error
   Get(id int) error
}

type UserService struct {
   store UserStore
}

// Accepting interface here!
func NewUserService(s UserStore) *UserService {
   return &UserService{
      store: s,
   }
}

func (u *UserService) CreateUser() { ... }
func (u *UserService) RetrieveUser(id int) User { ... }
```
