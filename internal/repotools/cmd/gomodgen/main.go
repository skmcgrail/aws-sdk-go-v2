package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/internal/repotools"
	"github.com/aws/aws-sdk-go-v2/internal/repotools/git"
	"github.com/aws/aws-sdk-go-v2/internal/repotools/gomod"
	"github.com/aws/aws-sdk-go-v2/internal/repotools/manifest"
	"golang.org/x/mod/modfile"
	"golang.org/x/mod/semver"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const manifestFileName = "generated.json"

var config = struct {
	BuildArtifactPath string
}{}

func init() {
	flag.StringVar(&config.BuildArtifactPath, "build", "", "build artifact path")
}

func main() {
	flag.Parse()

	if len(config.BuildArtifactPath) == 0 {
		log.Fatalf("expect build artifact path to be provided")
	}

	if stat, err := os.Stat(filepath.Join(config.BuildArtifactPath)); err != nil {
		log.Fatalf("failed to stat build artifact path: %v", err)
	} else if !stat.IsDir() {
		log.Fatalf("build artifact path must be a directory")
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current working directory: %v", err)
	}

	repoRoot, err := repotools.FindRepoRoot(cwd)
	if err != nil {
		log.Fatalf("failed to determine repository root: %v", err)
	}

	rootMod, err := gomod.LoadModuleFile(filepath.Join(repoRoot, "go.mod"), nil)
	if err != nil {
		log.Fatalf("failed to read repo root go module, %v", err)
	}

	rootModulePath, err := gomod.GetModulePath(rootMod)
	if err != nil {
		log.Fatalf("unable to determine repo root module path, %v", err)
	}

	av := manifest.SmithyArtifactFinder{}
	if err = filepath.Walk(config.BuildArtifactPath, av.Walk); err != nil {
		log.Fatalf("failed to walk build artifacts: %v", err)
	}

	if len(av) == 0 {
		log.Fatalf("no build artifacts found: %v", err)
	}

	taggedModules, err := getModuleTags(repoRoot)
	if err != nil {
		log.Fatalf("failed to query module tags: %v", err)
	}

	for _, artifactPath := range av {
		buildManifest, err := manifest.LoadManifest(filepath.Join(artifactPath, manifestFileName))
		if err != nil {
			log.Fatalf("failed to load manifest: %v", err)
		}
		if !strings.HasPrefix(buildManifest.Module, rootModulePath) {
			log.Fatalf("%v is not a sub-module of %v", buildManifest.Module, rootModulePath)
		}

		moduleRelativePath := strings.TrimPrefix(buildManifest.Module+"/", rootModulePath)

		targetPath := filepath.Join(repoRoot, moduleRelativePath)
		if err := prepareTargetDirectory(targetPath, buildManifest.Module); err != nil {
			log.Fatalf("failed to prepare target directory: %v", err)
		}

		if err := copyArtifactsToPath(artifactPath, targetPath, buildManifest); err != nil {
			log.Fatalf("failed to copy build artifact to target: %v", err)
		}

		generated, err := generateModuleDefinition(rootModulePath, buildManifest, taggedModules)
		if err != nil {
			log.Fatalf("failed to generate go module file: %v", err)
		}

		err = gomod.WriteModuleFile(targetPath, generated)
		if err != nil {
			log.Fatalf("failed to write go module file: %v", err)
		}
	}
}

func getModuleTags(repoRoot string) (git.ModuleTags, error) {
	if err := git.Fetch(repoRoot); err != nil {
		return nil, err
	}

	tags, err := git.Tags(repoRoot)
	if err != nil {
		return nil, err
	}

	return git.ParseTags(tags), nil
}

func generateModuleDefinition(rootModule string, m manifest.Manifest, tags git.ModuleTags) (*modfile.File, error) {
	var mod modfile.File

	if err := mod.AddGoStmt(m.Go); err != nil {
		return nil, fmt.Errorf("failed to set Go version: %v", err)
	}

	for depPath, depVersion := range m.Dependencies {
		latest := depVersion

		if strings.HasPrefix(depPath, rootModule) {
			tagVersion, ok := tags.Latest(strings.TrimPrefix(depPath+"/", rootModule))
			if ok && semver.Compare(tagVersion, latest) > 0 {
				latest = tagVersion
			}
		}

		if err := mod.AddRequire(depPath, latest); err != nil {
			return nil, fmt.Errorf("failed to add dependency %v@%v", depPath, depVersion)
		}
	}

	return &mod, nil
}

func prepareTargetDirectory(path string, module string) error {
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	} else if err != nil {
		return err
	}

	targetManifest, err := manifest.LoadManifest(filepath.Join(path, manifestFileName))
	if err != nil {
		return err
	}
	if targetManifest.Module != module {
		return fmt.Errorf("target module %v does not match build artifact %v", targetManifest.Module, module)
	}

	for _, fileName := range targetManifest.Files {
		filePath := filepath.Join(path, fileName)
		log.Printf("removing %v", filePath)
		if err := os.Remove(filePath); err != nil {
			return fmt.Errorf("failed to remove %v: %w", filePath, err)
		}
	}

	return nil
}

func copyArtifactsToPath(source, target string, m manifest.Manifest) error {
	for _, fp := range m.Files {
		sfp := filepath.Join(source, fp)
		tfp := filepath.Join(target, fp)

		if err := copyArtifact(sfp, tfp); err != nil {
			return err
		}
	}
	return nil
}

func copyArtifact(sourcePath, targetPath string) (err error) {
	dirs, _ := filepath.Split(targetPath)
	if len(dirs) != 0 {
		err = os.MkdirAll(dirs, 0755)
		if err != nil {
			return err
		}
	}

	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	targetFile, err := os.OpenFile(targetPath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() {
		fErr := targetFile.Close()
		if fErr != nil && err == nil {
			err = fErr
		}
	}()

	_, err = io.Copy(targetFile, sourceFile)
	return err
}
