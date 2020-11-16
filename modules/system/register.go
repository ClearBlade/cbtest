package system

import (
	"strings"
	"testing"

	cb "github.com/clearblade/Go-SDK"
)

// cbRegisterDeveloper registers a new developer in the system if it doesn't
// exists already.
func cbRegisterDeveloper(t *testing.T, system *EphemeralSystem) error {

	devClient := cb.NewDevClientWithAddrs(system.PlatformURL(), system.MessagingURL(), "", "")

	email := system.config.Developer.Email
	password := system.config.Developer.Password
	firstname := "cbtest"
	lastname := "cbtest"
	org := "ClearBlade, Inc."
	regkey := system.config.RegistrationKey

	_, err := devClient.RegisterDevUserWithKey(email, password, firstname, lastname, org, regkey)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return err
	}

	return nil
}

// cbRegisterUser registers a new user in the system if it doesn't exists already.
func cbRegisterUser(t *testing.T, system *EphemeralSystem) error {

	devClient, err := LoginAsDevE(t, system)
	if err != nil {
		return err
	}

	email := system.config.User.Email
	password := system.config.User.Password

	_, err = devClient.RegisterNewUser(email, password, system.SystemKey(), system.SystemSecret())
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return err
	}

	user, err := devClient.GetUserInfo(system.SystemKey(), email)
	if err != nil {
		return err
	}

	err = devClient.AddUserToRoles(system.SystemKey(), user["user_id"].(string), []string{"Administrator"})
	if err != nil {
		return err
	}

	return nil
}
