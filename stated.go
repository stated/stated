package main

import (
	"flag"
	"github.com/stated/stated/parser"
	"github.com/stated/stated/state"
	// "log"
	// "os/exec"
	"path/filepath"
)

var state_file = flag.String("state", "", "state file")
var plugin_path = flag.String("plugin", "", "plugin path")

func main() {

	// path, err := exec.LookPath("./stated")
	// if err != nil {
	// 	log.Fatal("cannot find stated")
	// }
	// log.Printf("stated is available at %s\n", path)

	flag.Parse()
	state_file_path, _ := filepath.Abs(*state_file)
	state_file_path = filepath.Clean(state_file_path)
	//state_file_dir := filepath.Dir(state_file_path)

	// FIXME: Avoid parsing multiple times by each plugin
	stated, _ := parser.Parse(state_file_path)

	// for _, section := range stated.Start {
	// 	log.Printf("%s\n", section)
	// }
	// log.Println("-------------------------------------------------\n")
	// for _, s := range stated.File {
	// 	log.Printf("%s (%s) %s %d\n", s.Destination, s.Source, s.Mode, s.Mode)
	// }
	// log.Println("-------------------------------------------------\n")

	if *state_file != "" {
		if *plugin_path == "" {
			cmd := state.Serve(stated)
			state.Client(stated)
			cmd.Process.Kill()
		}
	} else {
		if *plugin_path != "" {
			state.Listen()
		}
	}
}
