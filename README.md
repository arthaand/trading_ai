OttoPoint Pos Service is microservices for integration ottopoint x hellobill

# Prerequisites
1. Install latest golang version. Visit https://golang.org/doc/install for detail instructions. Please use always use latest stable golang.
2. Add bin folder from GOPATH to your environment. 
   For Mac User, add "export PATH=$HOME/go/bin:$PATH" to $HOME/.zshrc
   For Windows User, please visit https://www.architectryan.com/2018/08/31/how-to-change-environment-variables-on-windows-10
   For Linux User, add "export PATH=$HOME/go/bin:$PATH" to $HOME/.bash_profile.
3. Install swag command line. We use Swag CLI to generate swagger docs on Api Gateway. Execute "go get -u github.com/swaggo/swag/cmd/swag" in your terminal. Execute this command "swag --help" to check swag is working.
4. For migration tool, we use goose cli. For installation cli, please visit and follow the instruction at https://github.com/pressly/goose.
4. Install docker. 
5. Execute this command "go env -w GOPRIVATE=andromeda.ottopay.id/*" in your terminal. Because we use our private repository on go mod.
6. Execute in your terminal: GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgres://{postgres_user}:{postgres_password}@{postgres_host}:{postgres_port}/{postgres_dbname}?sslmode=disable" goose status
7. Set your git global configuration (type this : git config --global url.git@andromeda.ottopay.id:/.insteadOf https://andromeda.ottopay.id/)

# Caution!
- Because Api Gateway used as API Documentation, you HAVE to make sure you're already execute "swag init -g main.go --parseDependency" before commit your changes.
- If you want to changes tables DML or DDL operations or all things related to database, you have to add exec goose create {filename} sql