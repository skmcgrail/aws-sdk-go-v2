package manifest

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type Manifest struct {
	Module       string            `json:"module"`
	Go           string            `json:"go"`
	Dependencies map[string]string `json:"dependencies"`
	Files        []string          `json:"files"`
}

func ValidateManifest(manifest Manifest) error {
	if len(manifest.Go) == 0 {
		return fmt.Errorf("missing Go minimum version")
	}
	if len(manifest.Module) == 0 {
		return fmt.Errorf("missing module path")
	}
	return nil
}

func LoadManifest(path string) (Manifest, error) {
	mf, err := os.Open(path)
	if err != nil {
		return Manifest{}, fmt.Errorf("failed to open manifest: %w", err)
	}
	defer mf.Close()
	return ReadManifest(mf)
}

func ReadManifest(reader io.Reader) (m Manifest, err error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return Manifest{}, err
	}
	if err = json.Unmarshal(data, &m); err != nil {
		return Manifest{}, err
	}
	if err = ValidateManifest(m); err != nil {
		return Manifest{}, err
	}
	return m, nil
}

type SmithyArtifactFinder []string

func (a *SmithyArtifactFinder) Walk(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if !info.IsDir() {
		return nil
	}

	pluginOutput := filepath.Join(path, "go-codegen")
	stat, err := os.Stat(pluginOutput)
	if err != nil {
		return nil
	}

	if !stat.IsDir() {
		return nil
	}

	*a = append(*a, pluginOutput)

	return filepath.SkipDir
}
