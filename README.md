# Go Module & Run Commands

Here are some common Go commands for initializing modules and running code:

```bash
# Initialize a new Go module in the current folder
# Example: go mod init hello
# This creates a go.mod file and turns the folder into a Go module
go mod init <module-name>

# Update dependencies, remove unused ones, and update go.sum
go mod tidy

# Run the whole project (automatically finds main.go)
go run .

# Run a specific file
go run main.go
