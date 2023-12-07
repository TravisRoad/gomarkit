package setup

import (
	"log/slog"

	"github.com/sqids/sqids-go"
)

func initSqids() *sqids.Sqids {
	s, err := sqids.New(sqids.Options{
		Alphabet: "0123456789AbcDefghIjklmnoPqrStuVwXyZ",
	})
	if err != nil {
		slog.Error("sqids init error", err)
		panic(err)
	}
	return s
}
