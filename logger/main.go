package main

import (
	"learngo-pockets/logger/pocketlog"
	"os"
	"time"
)

func main() {
	lgr := pocketlog.New(pocketlog.LevelDebug, pocketlog.WithOutput(os.Stdout),
			pocketlog.AddLogMessageOptions(
			pocketlog.AddPrefixBasedOnLogLevel(),
			// pocketlog.AddDate(),
		))
	lgr.Infof("A little copying is better than a little dependency.")
	lgr.Errorf("Errors are values. Documentation is for %s.", "users")
	lgr.Debugf("Make the zero (%d) value useful.", 0)
	lgr.Infof("Hallo, %d %v", 2022, time.Now())

}
