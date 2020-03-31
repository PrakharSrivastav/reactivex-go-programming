package source

import "context"

type DBConnection string

const MockDBConnection DBConnection = "mock-connexion"

func Init(_ context.Context) (DBConnection, error) {
	return MockDBConnection, nil
}
