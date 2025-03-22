DEFAULT_GOAL := dev

dev:
	air

gqlgenc:
	go generate tools.go