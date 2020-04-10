// Copyright © 2018 Sugat Poudel <taguspoudel@gmail.com>

package cmds

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/sugatpoudel/crypt/internal/utils"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List stored services",
	Long:    `Lists the name of all stored service credentials.`,
	Run:     ls,
	Aliases: []string{"list"},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

func ls(cmd *cobra.Command, args []string) {
	creds := getStore().Crypt.Credentials

	data := make([][]string, len(creds))
	counter := 0
	for _, v := range creds {
		createdAt := v.GetCreatedAt().Format("Jan _2 2006")
		data = append(data, []string{strconv.Itoa(counter), v.Service, createdAt})
		counter++
	}

	utils.PrintTable(data, []string{"index", "name", "created at"}, "credentials")
	fmt.Printf("%d credential(s).\n", len(creds))
}
