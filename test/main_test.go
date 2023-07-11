package test
import(
	"testing"
	"strings"
	"github.com/liujiage/restapi/service"
)

func TestStart_test(t *testing.T) {
	s := service.Say("jiage liu")
	if !strings.Contains(s, "Hello") {
		t.Fatalf("fail! %s", s)
	}
}