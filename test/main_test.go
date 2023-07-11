package test
import(
	"testing"
	"strings"
	"github.com/liujiage/restapi/service"
)

func TestMain_main(t *testing.T) {
	s := service.Say("jiage liu")
	if !strings.Contains(s, "Hello") {
		t.Fatalf("fail! %s", s)
	}
}