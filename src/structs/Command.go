package structs
import "velk/src/interfaces"

type Command struct {
	Player        interfaces.PlayerInterface
	CommandName   string
	CommandSuffix string
}
