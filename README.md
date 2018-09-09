1. Enable Google Tasks API.
  - Select + Create a new project.
  - Download the configuration file.
  - Move the downloaded file to your working directory and ensure it is named `credentials.json`

2. Grab go dependencies:

        go get -u google.golang.org/api/tasks/v1
        go get -u golang.org/x/oauth2/...
        go get -u golang.org/x/net/context
