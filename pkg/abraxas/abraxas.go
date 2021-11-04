package abraxas

import (
	"github.com/antonioiubatti93/abraxas/logger"
	baphometlog "github.com/antonioiubatti93/baphomet/logger"
)

const answer = "I'm Abraxas, at your service!"

func Invoke(logger logger.Logger) {
	logger.Print(answer)
}

func Mediate(logger baphometlog.Logger) {
	logger.Print("inovking Abraxas now")
	Invoke(logger)
}
