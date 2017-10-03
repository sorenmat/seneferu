package main

import (
	"encoding/base64"
	"fmt"
	"github.com/cncd/pipeline/pipeline/frontend/yaml"
	"strings"
	"testing"
)

func TestSomePath(t *testing.T) {
	commands := []string{"ls -la"}
	b64 := generateScript(commands)
	cmd, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(string(cmd), "ls -la") {
		t.Error("Seems like the command wasn't generated correctly: ", err)
	}
}

func TestDoneCmd(t *testing.T) {
	if doneCmd(0) == "build.done" {
		t.Error()
	}
	if doneCmd(1) == "build0.done" {
		t.Error()
	}

}

func TestCreateBuild(t *testing.T) {
	build := &Build{Number: 1, Owner: "sorenmat", Repo: "ci-server"}
	container := &yaml.Container{Environment: map[string]string{"Name": "sorenmat"}}
	cfg := &yaml.Config{}
	cfg.Pipeline.Containers = append(cfg.Pipeline.Containers, container)

	steps, err := createBuildSteps(build, cfg)
	if err != nil {
		t.Error("createBuildStep should not have failed")
	}
	if len(steps) != 1 {
		t.Error("expected 1 build")
	}
	fmt.Println(steps)
	found := false
	for _, value := range steps[0].Env {
		if value.Name == "Name" && value.Value == "sorenmat" {
			found = true
		}
	}
	if !found {
		t.Error("should have found environment variable in container")
	}
}