# gRPC Go

## Notes

### Windows

- It is recommended to use PowerShell 7 (try to update: [see](https://github.com/PowerShell/PowerShell/releases)). You might have unexpected behavior if you use Git Bash or some other shell (especially with OpenSSL).
- I recommend you use [Chocolatey](https://chocolatey.org/) as package installer (see [Install](https://chocolatey.org/install)).


### Build Code

#### Linux/MacOS

```shell
make users
```
***`users` is a Makefile rule** - check the other rules [here](#makefile).

#### `Windows - Chocolatey`
```shell
choco install make
make users
```
***`users` is a Makefile rule** - check the other rules [here](#makefile).

#### Windows - Without Chocolatey

```shell
protoc -Iproto --go_opt=module=github.com/mrehanabbasi/user-data-grpc --go_out=. --go-grpc_opt=module=github.com/mrehanabbasi/user-data-grpc --go-grpc_out=. proto/*.proto

go build -o bin/server.exe ./server
go build -o bin/client.exe ./client
```

## Running the Code

### Linux/MacOS
```shell
./bin/server
./bin/client
```

### Windows
```powershell
.\bin\server.exe
.\bin\client.exe
```

## Generating SSL Certificates

```shell
make certs
```
<a name="makefile"></a>
## Makefile

For more information about what are the rules defined in the Makefile, please type:

```shell
make help
```

## Reporting a bug

When filing an issue, please provide the output of:

```shell
make about
```
