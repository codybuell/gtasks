package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/tasks/v1"
)

var helpBool bool
var versionBool bool

func getClient(config *oauth2.Config) *http.Client {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	json.NewEncoder(f).Encode(token)
}

func parseFlags(flag *flag.Flag) {
	fmt.Println(">", flag.Name, "value=", flag.Value)
	switch flag.Name {
	case "v":
		fmt.Print("show version\n")
	case "h":
		fmt.Print("show help\n")
	default:
		fmt.Printf("The %s option has not been configured yet\n", flag.Name)
		printUsage()
	}
}

func printUsage() {
	/**
	 * t [-vh] [arguments]
	 *
	 * flags:
	 *   -v                                       show tool version
	 *   -h                                       print out this help text
	 *
	 * arguments:
	 *                                            no arguments prints all tasks
	 *   @context                                 print all tasks of context
	 *   +project                                 print all tasks of project (google task list)
	 *   <id>                                     print details of task with id
	 *   mit                                      print out mits
	 *   mit <day>                                print out mits of specified day
	 *                                            <day> can be:
	 *                                                  day of week  (mon, monday)
	 *                                                  relative day (tomorrow)
	 *                                                  hard date    (YYYY.MM.DD)
	 *   a <task> [@context] [+project]           add task (defaults to ?? project and no context)
	 *   mit <day> <task> [@context] [+project]   add a mit (defaults to ?? project and no context)
	 *                                            <day> can be:
	 *                                                  day of week  (mon, monday)
	 *                                                  relative day (tomorrow)
	 *                                                  hard date    (YYYY.MM.DD)
	 *   do <id>                                  complete task <id>
	 *   rm <id>                                  delete task <id>
	 *
	 */
	fmt.Print("print some usage here\n")
}

func getTasksByContext(context string) {
	fmt.Printf("context: " + context + "\n")
}

func getTasksByProject(srv *tasks.Service, tasklist *tasks.TaskLists, project string) {
	// if just + list projects, else list tasks in queried project
	if project == "+" {
		fmt.Println("Task Lists:")
		if len(tasklist.Items) > 0 {
			for _, i := range tasklist.Items {
				fmt.Printf("%s (%s)\n", i.Title, i.Id)
			}
		} else {
			fmt.Println("No task lists found.")
		}
	} else {
		// strip + from project name
		tasklistname := strings.Replace(project, "+", "", 1)

		// grab tasklist id
		for _, i := range tasklist.Items {
			if i.Title == tasklistname {
				r, err := srv.Tasks.List(i.Id).Do()
				if err != nil {
					log.Fatalf("Unable to retrieve task lists.", err)
				}

				fmt.Println("Tasks in " + i.Title + ":")
				if len(r.Items) > 0 {
					for _, t := range r.Items {
						fmt.Printf("%s (%s)\n", t.Title, t.Id)
					}
				} else {
					fmt.Println("No projects found.")
				}
				break
			}
		}
	}

}

func processMITS(arguments []string) {
	// case statements
	fmt.Print("mit stuffs\n")
}

func processAdd(arguments []string) {
	fmt.Print("add tasks\n")
}

func processCompletion(arguments []string) {
	fmt.Print("complete tasks\n")
}

func processDeletion(arguments []string) {
	fmt.Print("delete tasks\n")
}

func getAllTasks() {
	fmt.Printf("print all tasks\n")
}

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

func main() {

	// read the credentials file
	creds, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// if modifying scopes, delete the previously saved token.json
	config, err := google.ConfigFromJSON(creds, tasks.TasksReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	// initialize new tasks client
	srv, err := tasks.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve tasks Client %v", err)
	}

	// query tasks lists
	tasklist, err := srv.Tasklists.List().Do()
	if err != nil {
		log.Fatalf("Unable to retrieve task lists.", err)
	}

	// process our flagged arguments
	flag.Visit(parseFlags)

	//fmt.Printf("versionBool: %t, helpBool: %t\n", versionBool, helpBool)

	// gather non flagged arguments (array of all args)
	arguments := flag.Args()
	//fmt.Printf("%s\n", arguments)    // print out args

	// bail if no aguments
	if len(arguments) == 0 {
		return
	}

	// process our non-flagged arguments
	switch {
	case strings.HasPrefix(arguments[0], "@"):
		getTasksByContext(arguments[0])
	case strings.HasPrefix(arguments[0], "+"):
		getTasksByProject(srv, tasklist, arguments[0])
	case arguments[0] == "mit":
		processMITS(arguments)
	case arguments[0] == "a":
		processAdd(arguments)
	case arguments[0] == "do":
		processCompletion(arguments)
	case arguments[0] == "rm":
		processDeletion(arguments)
	default:
		// if arg0 is a task id then show it, else usage
		getAllTasks()
	}

}
