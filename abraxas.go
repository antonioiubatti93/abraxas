package abraxas

import "github.com/antonioiubatti93/abraxas/logger"

const answer = "I'm Abraxas, at your service!"

func Invoke(logger logger.Logger) {
	logger.Print(answer)
}
