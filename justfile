run:
    @templ generate
    @gomod2nix generate
    @go run main.go
    @echo " "
