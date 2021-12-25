package schema

import (
	"GustavoDelfim/graphql-go-example/resolver"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/graph-gophers/graphql-go"
)

func SchemasString() (string, error) {
	var buf bytes.Buffer
	var files []string

	readSchemas := func(path string, info os.FileInfo, _ error) error {
		if !strings.HasSuffix(path, ".graphql") {
			return nil
		}
		files = append(files, path)

		b, err := ioutil.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading file %q: %w", path, err)
		}

		b = append(b, []byte("\n")...)

		if _, err := buf.Write(b); err != nil {
			return fmt.Errorf("writing %q bytes to buffer: %w", path, err)
		}

		return nil
	}

	err := filepath.Walk("./schema", readSchemas)
	if err != nil {
		return buf.String(), fmt.Errorf("walking content directory: %w", err)
	}

	return buf.String(), nil
}

func GetSchema(resolver *resolver.RootResolver) *graphql.Schema {
	opts := []graphql.SchemaOpt{
		graphql.UseFieldResolvers(),
		graphql.MaxParallelism(int(50)),
	}

	schemaString, err := SchemasString()

	if err != nil {
		panic(err)
	}

	parsedSchema, err := graphql.ParseSchema(
		schemaString,
		resolver,
		opts...,
	)

	if err != nil {
		panic(err)
	}

	return parsedSchema
}
