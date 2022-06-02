# gh-members
A small tool to find members of a orgs in Github

Pre-requisites:

* You need to config an env called GH_AUTH_TOKEN with your personal access token, to do the requests

How to use:

```bash
cat orgs.txt | gh-members
```

Installation
First, you'll need to install go.

Then run this command to download + compile gh-members:

go install github.com/edivangalindo/gh-members@latest
You can now run ~/go/bin/gh-members. If you'd like to just run gh-members without the full path, you'll need to export PATH="/go/bin/:$PATH". You can also add this line to your ~/.bashrc file if you'd like this to persist.
