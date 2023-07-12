package common
import(
	"github.com/google/uuid"
)

/***
Get uuid
***/
func GetUUID() string{
	return uuid.New().String()
}