package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/rentziass/gh-runmd/blocks"
)

func main() {
	fmt.Println("hi world, this is the gh-runmd extension!")
	client, err := api.DefaultRESTClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	response := struct{ Login string }{}
	err = client.Get("user", &response)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("running as %s\n", response.Login)

	blocks := blocks.Blocks{
		&blocks.Block{Content: "echo hello\n"},
		&blocks.Block{Content: "echo world\n"},
	}

	cmd := exec.Command("bash", "-c", blocks.ToScript())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
