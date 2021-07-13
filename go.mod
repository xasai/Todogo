module github.com/xasai/todogo

go 1.16

replace github.com/xasai/todogo/internal/cli => ./internal/cli

require (
	github.com/golang/protobuf v1.5.2
	github.com/inancgumus/screen v0.0.0-20190314163918-06e984b86ed3
	github.com/xasai/todogo/internal/cli v0.0.0-00010101000000-000000000000
	golang.org/x/sys v0.0.0-20210331175145-43e1dd70ce54 // indirect
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
)
