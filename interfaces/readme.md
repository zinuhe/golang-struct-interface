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
                                                                      |
type UserStore interface {                                            | type UserService struct {
   Insert(item interface{}) error                                     |   store db.Store
   Get(id int) error                                                  | }
}                                                                     |
                                                                      | func NewUserService(s db.Store) *UserService {
type UserService struct {                                             |   return &UserService{
   store UserStore                                                    |       store: s,
}                                                                     |   }
                                                                      | }
// Accepting interface here!                                          |
func NewUserService(s UserStore) *UserService {                       |
   return &UserService{                                               | func (u *UserService) CreateUser() { ... }
      store: s,                                                       | func (u *UserService) RetrieveUser(id int) User { ... }
   }                                                                  |
}                                                                     |

func (u *UserService) CreateUser() { ... }
func (u *UserService) RetrieveUser(id int) User { ... }
```

More Information
[accept-interfaces-return-structs-in-go](https://bryanftan.medium.com/accept-interfaces-return-structs-in-go-d4cab29a301b)
