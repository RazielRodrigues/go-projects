package persistence

import "kv-store/internal/domain/repository"

type AOFProviderOption struct {
	EnableAOF bool
	FilePath  string
}

func NewAOFProvider(opt AOFProviderOption) (repository.PersistencxeRepository, error) {
	if !opt.EnableAOF {
		return nil, nil
	}

	return NewAOF(opt.FilePath)
}
