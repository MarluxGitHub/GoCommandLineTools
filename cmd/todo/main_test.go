package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binName = "todo"
	fileName = "todo.json"
)

func TestMain(m *testing.M){
	fmt.Println("Building tool ...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName, ".")

	if err := build.Run(); err != nil {
		fmt.Println("Error building tool:", err)
		os.Exit(1)
	}

	fmt.Println("Running tests ...")

	result := m.Run()

	fmt.Println("Cleaning up ...")

	os.Remove(binName)
	os.Remove(fileName)



	os.Exit(result)
}

func TestTodoCli(t *testing.T) {
	task := "test task number 1"

	dir, err := os.Getwd()

	if err != nil {
		t.Error(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("add", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", task)

		if err := cmd.Run(); err != nil {
			t.Error(err)
		}
	})

	t.Run("List Tasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")

		out, err := cmd.CombinedOutput()

		if err != nil {
			t.Error(err)
		}

		expected := "0.:   " + task + "\n"

		if string(out) != expected {
			t.Errorf("expected %s, got %s", expected, out)
		}
	})

	t.Run("List open Tasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list", "-open")

		out, err := cmd.CombinedOutput()

		if err != nil {
			t.Error(err)
		}

		expected := "0.:   " + task + "\n"

		if string(out) != expected {
			t.Errorf("expected %s, got %s", expected, out)
		}
	})

	t.Run("Complete Task", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-complete", "0")

		if err := cmd.Run(); err != nil {
			t.Error(err)
		}
	})

	t.Run("List Completed Tasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")

		out, err := cmd.CombinedOutput()

		if err != nil {
			t.Error(err)
		}

		expected := "0.: X " + task + "\n"

		if string(out) != expected {
			t.Errorf("expected %s, got %s", expected, out)
		}
	})

	t.Run("List open Tasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list", "-open")

		out, err := cmd.CombinedOutput()

		if err != nil {
			t.Error(err)
		}

		expected := ""

		if string(out) != expected {
			t.Errorf("expected %s, got %s", expected, out)
		}
	})

	t.Run("Delete Task", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-delete", "0")

		if err := cmd.Run(); err != nil {
			t.Error(err)
		}
	})

	t.Run("List Tasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")

		out, err := cmd.CombinedOutput()

		if err != nil {
			t.Error(err)
		}

		expected := ""

		if string(out) != expected {
			t.Errorf("expected %s, got %s", expected, out)
		}
	})
}