package h3

import (
	"fmt"
	"os/exec"
	pathlib "path"
	"strings"

	log "github.com/sirupsen/logrus"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

// H3ArtifactDriver is a driver for H3
type H3ArtifactDriver struct {
	StorageUri string
}

// Load downloads artifacts from H3 compliant storage
func (h3Driver *H3ArtifactDriver) Load(inputArtifact *wfv1.Artifact, path string) error {
	log.Infof("H3 load path: %s, bucket: %s, key: %s", path, inputArtifact.H3.Bucket, inputArtifact.H3.Key)

	command := []string{"h3cli",
						"--storage",
						h3Driver.StorageUri,
						"cp",
						fmt.Sprintf("h3://%s", pathlib.Join(inputArtifact.H3.Bucket, inputArtifact.H3.Key)),
						path}
	log.Infof("H3 running: %s", strings.Join(command, " "))

	out, err := exec.Command(command[0], command[1:]...).Output()
	if err != nil {
		log.Errorf(err.Error())
		output := string(out[:])
		log.Infof("H3 output: %s", output)
		return err
	}
	output := string(out[:])
	log.Infof("H3 output: %s", output)

	return nil
}

// Save saves an artifact to H3 compliant storage
func (h3Driver *H3ArtifactDriver) Save(path string, outputArtifact *wfv1.Artifact) error {
	log.Infof("H3 save path: %s, bucket: %s, key: %s", path, outputArtifact.H3.Bucket, outputArtifact.H3.Key)

	command := []string{"h3cli",
						"--storage",
						h3Driver.StorageUri,
						"cp",
						path,
						fmt.Sprintf("h3://%s", pathlib.Join(outputArtifact.H3.Bucket, outputArtifact.H3.Key))}
	log.Infof("H3 running: %s", strings.Join(command, " "))

	out, err := exec.Command(command[0], command[1:]...).Output()
	if err != nil {
		log.Errorf(err.Error())
		output := string(out[:])
		log.Infof("H3 output: %s", output)
		return err
	}
	output := string(out[:])
	log.Infof("H3 output: %s", output)

	return nil
}
