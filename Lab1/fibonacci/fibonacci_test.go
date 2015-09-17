package fibpackage
import "testing"

func TestFib(t *testing.T) {
var v int64
v = ComputeFibSum(5)
if v != 5 {
t.Error("Expected 5, got ", v)
}
}