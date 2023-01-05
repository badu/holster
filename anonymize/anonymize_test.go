package anonymize

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AnonymizeSuite struct {
	suite.Suite
}

func TestAnonymizeSuite(t *testing.T) {
	suite.Run(t, new(AnonymizeSuite))
}

func (s *AnonymizeSuite) TestAnonymizeWithNoSecretsReturnsUnchangedString() {
	anonimized, err := Anonymize("Hello Dear User", "")
	s.Nil(err)
	s.Equal("Hello Dear User", anonimized)
	anonimized, err = Anonymize("Hello Dear User")
	s.Nil(err)
	s.Equal("Hello Dear User", anonimized)
}

func (s *AnonymizeSuite) TestAnonymizeWithNoSecretsInStringReturnsUnchangedString() {
	anonimized, err := Anonymize("Hello Dear User", "John")
	s.Nil(err)
	s.Equal("Hello Dear User", anonimized)
}

func (s *AnonymizeSuite) TestAnonymizeEscapesSecrets() {
	anonimized, err := Anonymize(
		"Hello Dear User", `\s`)
	s.Nil(err)
	s.Equal("Hello Dear User", anonimized)
}

func (s *AnonymizeSuite) TestAnonymizeSquashesAdjacentSecrets() {
	anonimized, err := Anonymize(
		"Hello Иван Иванов ivan ivanov foo.bar",
		`"Иван Иванов" <ivan.ivanov@foo.bar>`)
	s.Nil(err)
	s.Equal("xxx", anonimized)
}

func (s *AnonymizeSuite) TestAnonymizeNames() {
	subjects := map[string]string{
		"ALEX MIA - Welcome to YORK":               "xxx MIA - Welcome to YORK",
		"MIA ALEX - Welcome to YORK":               "MIA xxx - Welcome to YORK",
		"MIA ALEX, Welcome to YORK":                "MIA xxx, Welcome to YORK",
		"MIA ALEX, welcome to YORK":                "MIA xxx, welcome to YORK",
		"Mia ALEX - Welcome to YORK":               "Mia xxx - Welcome to YORK",
		"ALEX Mia - Welcome to YORK":               "xxx Mia - Welcome to YORK",
		"Alex Mia - Welcome to YORK":               "xxx Mia - Welcome to YORK",
		"Mia Alex, Welcome to YORK":                "Mia xxx, Welcome to YORK",
		"ALEX MIA and BORNE ROY, Welcome to YORK":  "xxx MIA and BORNE xxx, Welcome to YORK",
		"Weekly Facebook Page":                     "Weekly Facebook Page",
		"Robert and ROY are coming, are you Alex?": "xxx and xxx are coming, are you xxx?",
	}
	for subject, expected := range subjects {
		anonimized, err := Anonymize(subject, `ivan.ivanov@foo.bar`)
		s.Nil(err)
		s.Equal(expected, anonimized)
		anonimized, err = Anonymize(subject)
		s.Nil(err)
		s.Equal(expected, anonimized)
	}
}
