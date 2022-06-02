package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {

	token := os.Getenv("GH_AUTH_TOKEN")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Fprintln(os.Stderr, "No orgs detected. Hint: cat orgs.txt | gh-members")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		org := scanner.Text()

		var allMembers []*github.User

		opt := &github.ListMembersOptions{
			ListOptions: github.ListOptions{PerPage: 100},
		}

		for {
			members, resp, err := client.Organizations.ListMembers(ctx, org, opt)

			if err != nil {
				fmt.Println(err)
			}

			allMembers = append(allMembers, members...)

			opt.Page = resp.NextPage

			if resp.NextPage == 0 {
				break
			}
		}

		for _, m := range allMembers {
			fmt.Println(*m.Login)
		}
	}
}
