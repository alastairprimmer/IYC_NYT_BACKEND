.PHONY: build

clean:  $(EDGEBUILDERCNTLR)
	rm -rf graphql/graph/generated

types:
	go run github.com/99designs/gqlgen generate