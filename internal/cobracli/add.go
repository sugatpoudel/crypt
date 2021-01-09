package cobracli

import (
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/sugatpoudel/crypt/internal/asker"
	"github.com/sugatpoudel/crypt/internal/creds"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [service]",
	Short: "Add a service to crypt",
	Long: `Add a service along with any associated information
to the crypt getStore().

Expects a single argument, however multi word services
can be espaced using quotes.`,
	Args:    serviceIsNew,
	Example: "add 'Amazon Web Services'",
	RunE:    add,
}

func add(cmd *cobra.Command, args []string) error {
	service := args[0]
	asker := asker.DefaultAsker()

	email, err := asker.Ask("Email")
	if err != nil {
		return err
	}

	user, err := asker.Ask("Username")
	if err != nil {
		return err
	}

	pwd, err := asker.AskSecret("Password", true)
	if err != nil {
		return err
	}

	desc, err := asker.Ask("Description")
	if err != nil {
		return err
	}

	cred := creds.Credential{
		Service:     service,
		Email:       email,
		Username:    user,
		Password:    pwd,
		Description: desc,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}

	cred.PrintCredential()

	ok, err := asker.AskConfirm("Does this look right?")
	if ok {
		st, err := getStore()
		if err != nil {
			return err
		}
		st.Crypt.SetCredential(cred)
		color.Green("\nAdded service '%s'", service)
		return saveStore()
	}
	return nil
}