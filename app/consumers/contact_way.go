package consumers

import (
	"github.com/pkg/errors"
	"openscrm/app/models"
	"openscrm/common/delay_queue"
	"openscrm/common/util"
)

type contactWay struct {
}

func (o contactWay) Refresh(job delay_queue.Job) (err error) {
	defer util.FuncTracer("job", job)()
	_, err = models.ContactWay{}.Refresh(models.DB, job.Body)
	if err != nil {
		err = errors.Wrap(err, "refresh contactWay")
		return
	}

	return
}
