package sink

import "context"

type FileConnection string

const MockFileConnection FileConnection = "mock-file-connexion"

func Init(_ context.Context) (FileConnection, error) {
	return MockFileConnection, nil
}

