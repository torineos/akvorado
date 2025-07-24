package hostname

// Create a mock of /etc/host or /etc/resolv.conf so that it doesn't create DNS requests when testing at every CI from the host.
// Make sure to read the file or the infos on here (or create a false resolv.conf with already all the informations in it) so
// that it fetches the infos locally
// import (
// 	"testing"

// 	"akvorado/common/daemon"
// 	"akvorado/common/helpers"
// 	"akvorado/common/reporter"
// 	//"akvorado/common/schema"
// )

// // NewMock instantiantes a new mocked hostname component.
// func NewMock(t *testing.T, r *reporter.Reporter, config Configuration) *Component {
// 	t.Helper()
// 	c, err := New(r, config, Dependencies{
// 		Daemon: daemon.NewMock(t),
// 		//Schema: schema.NewMock(t),
// 	})
// 	if err != nil {
// 		t.Fatalf("New() error:\n%+v", err)
// 	}
// 	helpers.StartStop(t, c)

// 	return c
// }
