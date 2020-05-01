package parse

import (
	"bufio"
	"strings"

	"github.com/src-d/go-mysql-server/sql"
	"github.com/src-d/go-mysql-server/sql/plan"
)

func parseShowStatus(ctx *sql.Context, s string) (sql.Node, error) {
	var pattern string

	r := bufio.NewReader(strings.NewReader(s))
	for _, fn := range []parseFunc{
		expect("show"),
		skipSpaces,
		func(in *bufio.Reader) error {
			var s string
			if err := readIdent(&s)(in); err != nil {
				return err
			}

			switch s {
			case "global", "session":
				if err := skipSpaces(in); err != nil {
					return err
				}

				return expect("status")(in)
			case "status":
				return nil
			}
			return errUnexpectedSyntax.New("show [global | session] status", s)
		},
		skipSpaces,
		func(in *bufio.Reader) error {
			if expect("like")(in) == nil {
				if err := skipSpaces(in); err != nil {
					return err
				}

				if err := readValue(&pattern)(in); err != nil {
					return err
				}
			}
			return nil
		},
		skipSpaces,
		checkEOF,
	} {
		if err := fn(r); err != nil {
			return nil, err
		}
	}

	return plan.NewShowStatus(pattern), nil
}
