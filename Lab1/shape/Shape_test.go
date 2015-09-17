package shapepk
import "testing"

func TestShape(t *testing.T) {
r := Rectangle{2, 4}
var v float64
v = r.Area()
if v != 8 {
t.Error("Expected 8, got ", v)
}
}