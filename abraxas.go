package abraxas

import (
	"github.com/antonioiubatti93/abraxas/logger"
	"github.com/antonioiubatti93/baphomet"
	baphometlog "github.com/antonioiubatti93/baphomet/logger"
)

const answer = "I'm Abraxas, at your service!"

func Invoke(logger logger.Logger) {
	logger.Print(answer)
}

func Mediate(logger baphometlog.Logger) {
	baphomet.Summon(logger)
	logger.Print("inovking Abraxas now")
	Invoke(logger)
}
