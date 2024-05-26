run:
    @gomod2nix generate
    @templ generate
    @go run main.go
