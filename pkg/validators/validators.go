package validators

import (
	"net/mail"
	"net/url"
	"regexp"

	"github.com/google/uuid"
)

var phoneRegex = regexp.MustCompile(`^\+?[1-9]\d{7,14}$`)

func Email(v string) bool {
	_, err := mail.ParseAddress(v)
	return err == nil
}

func URL(v string) bool {
	u, err := url.ParseRequestURI(v)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func UUID(v string) bool {
	_, err := uuid.Parse(v)
	return err == nil
}

func Phone(v string) bool { return phoneRegex.MatchString(v) }
