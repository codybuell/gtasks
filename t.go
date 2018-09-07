/**
 *
 *  Use first line of notes field for custom meta, denote meta with a meta:
 *    meta: @context !A
 *
 *
 *
 *  LISTING:
 *  t @somthing               list all tasks with @something context
 *                            context name
 *                            ------------
 *                            [YYYY.MM.DD] this is a task +project
 *  t +project                list all tasks in +project project
 *                            project name
 *                            ------------
 *                            [YYYY.MM.DD] this is a task @context
 *  t mit                     list the next 7 days worth of tasks grouped by day
 *  t                         list all tasks
 *                            [YYYY.MM.DD] this is a task +project @context
 *  t [task id]               show task info (title, notes / meta, etc)
 *
 *  MANIPULATING:
 *  t a [task text]           add a task, from within task text parse any
 *                            contexts (@[^ ]*) or projects (+[^ ]*)
 *  t mit [dow] [task text]   create a dated task (up to 7 days out) by dow
 *                            (day of week: mon, tue, wed, etc), relative day
 *                            (tommorrow, +2, etc), or hard date (YYYY.MM.DD)
 *  t do [task id]            mark task as complete
 *  t rm [task id]            delete a task
 */

package main

import (
	"flag"
	"fmt"
	"strings"
)

var helpBool bool
var versionBool bool

/**
 * Initialization
 *
 * Define and grab our flagged arguments.
 *
 * @return {null}
 *
 */
func init() {
	// define usages
	const (
		helpUsage    = "show help text"
		versionUsage = "show the application version"
	)

	// define our options
	flag.BoolVar(&helpBool, "help", false, helpUsage)
	flag.BoolVar(&helpBool, "h", false, helpUsage+" (shorthand)")
	flag.BoolVar(&versionBool, "version", false, versionUsage)
	flag.BoolVar(&versionBool, "v", false, versionUsage+" (shorthand)")

	// grab and parse our options
	flag.Parse()
}

/**
 * Get Tasks by Context
 *
 * Print out all tasks of a given context.
 *
 * @param {string} context - context to query
 * @return {null}
 *
 */
func getTasksByContext(context string) {
	fmt.Printf("context: " + context + "\n")
}

/**
 * Get Tasks by Project
 *
 * Print out all tasks of a given project.
 *
 * @param {string} project - project to query
 * @return {null}
 *
 */
func getTasksByProject(project string) {
	fmt.Printf("project: " + project + "\n")
}

/**
 * Process MITS
 *
 * Process various MIT related functions.
 *
 * @param {string} arguments - all non-flagged cli arguments
 */
func processMITS(arguments []string) {
	// case statements
	fmt.Print("mit stuffs\n")
}

/**
 * Get All Tasks
 *
 * Print out all tasks.
 *
 * @return {null}
 *
 */
func getAllTasks() {
	fmt.Printf("print all tasks\n")
}

/**
 * Main
 *
 * @return {null}
 */
func main() {
	// process our flagged arguments
	fmt.Printf("versionBool: %t, helpBool: %t\n", versionBool, helpBool)

	// gather non flagged arguments (array of all args)
	arguments := flag.Args()

	// print our arguments
	fmt.Printf("%s\n", arguments)

	// process our non-flagged arguments
	switch {
	case strings.HasPrefix(arguments[0], "@"):
		getTasksByContext(arguments[0])
	case strings.HasPrefix(arguments[0], "+"):
		getTasksByProject(arguments[0])
	case arguments[0] == "mit":
		processMITS(arguments)
	case arguments[0] == "a":
		fmt.Print("add tasks\n")
	case arguments[0] == "do":
		fmt.Print("complete tasks\n")
	case arguments[0] == "rm":
		fmt.Print("delete tasks\n")
	default:
		// if arg0 is a task id then show it, else usage
		getAllTasks()
	}

}
