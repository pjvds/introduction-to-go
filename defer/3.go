import (
	"time"

	"github.com/happypancake/fsd"
	"github.com/happypancake/nanobus/guid"
	"github.com/happypancake/nanobus/request"
)

// START OMIT
func (r *client) RequestSingle(timeout time.Duration, req []byte) (response []byte, err error) {
	defer logRequestMetrics(time.Now(), &err) // HL

	envelop := request.NewEnvelope(guid.New(), req)
	context := request.NewContext()

	r.requests.Add(envelop.CorrelationId, context)
	defer r.requests.Delete(envelop.CorrelationId) // HL

	if err = r.socket.Send(envelop.MustPack()); err != nil {
		return
	}
	return context.ReceiveWithTimeout(timeout)
}

func logRequestMetrics(started time.Time, err *error) {
	fsd.TimeSinceL("nanobus.request", started, 0.1)

	if *err != nil {
		fsd.CountL("nanobus.request.err", 1, 0.1)
	} else {
		fsd.CountL("nanobus.request.ok", 1, 0.1)
	}
}

// END OMIT