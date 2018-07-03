package main

import (
	"fmt"
	"os"
	"os/exec"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("GitBuddy", "A simple program simplifying interacting with Git.")
	// debug    = app.Flag("debug", "Enable debug mode.").Bool()
	// serverIP = app.Flag("server", "Server address.").Default("127.0.0.1").IP()

	stage = app.Command("stage", "Stage changes.")
	st    = app.Command("st", "Stage changes.")

	commit        = app.Command("commit", "Commit changes")
	commitMessage = commit.Arg("message", "The message for this commit.").Required().String()

	cm        = app.Command("cm", "Commit changes")
	cmMessage = cm.Arg("message", "The message for this commit.").Required().String()

	log = app.Command("log", "Display Git Logs")

	/*
		post        = app.Command("post", "Post a message to a channel.")
		postImage   = post.Flag("image", "Image to post.").File()
		postChannel = post.Arg("channel", "Channel to post to.").Required().String()
		postText    = post.Arg("text", "Text to post.").Strings()
	*/
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	// stage changes
	case stage.FullCommand():
		StageChanges()

	case st.FullCommand():
		StageChanges()

	// commit
	case commit.FullCommand():
		CommitWithMessage(*commitMessage)

	case cm.FullCommand():
		CommitWithMessage(*cmMessage)

	//display logs
	case log.FullCommand():
		PrintLogs()

	default:
		fmt.Println("No command submitted or unknown command.")

		/*case post.FullCommand():
		if *postImage != nil {
		}
		text := strings.Join(*postText, " ")
		println("Post:", text)
		*/
	}

}

func mainW() {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "git"
	cmdArgs := []string{"rev-parse", "--verify", "HEAD"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git rev-parse command: ", err)
		os.Exit(1)
	}
	sha := string(cmdOut)
	firstSix := sha[:6]
	fmt.Println("The first six chars of the SHA at HEAD in this repo are", firstSix)
}

// StageChanges stages all changes and returns wheter it what sucessful.
func StageChanges() bool {
	// var cmdOut []byte
	success := false
	var err error

	cmdArgs := []string{"add", "."}
	if _, err = exec.Command("git", cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error staging all changes", err)
		return success
	}
	success = true
	return success
}

// PushRepo pushes to remote
func PushRepo() bool {
	// var cmdOut []byte
	success := false
	var err error

	cmdArgs := []string{"push"}
	if _, err = exec.Command("git", cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error staging all changes", err)
		return success
	}
	success = true
	return success
}

// CommitWithMessage commits staged changes with a message.
func CommitWithMessage(message string) bool {
	success := false

	cmdArgs := []string{"commit", "-am", message}

	if outp, err := exec.Command("git", cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error commiting your changes", err)
		fmt.Println(outp)
		return success
	}
	success = true
	return success
}

// PrintLogs print Git logs
func PrintLogs() bool {
	success := false
	out, err := exec.Command("git", "log").Output()
	if err != nil {
		fmt.Println(err)
		return success
	}
	fmt.Println(string(out))
	success = true
	return success
}
