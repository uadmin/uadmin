package uadmin

import (
	"encoding/json"
	"github.com/uadmin/rrd"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func getRRDTemplate(name string) (*rrd.RRD, error) {
	tmpl := rrd.RRD{}
	buf, err := ioutil.ReadFile("templates/uadmin/rrd/" + name + ".json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &tmpl)
	if err != nil {
		return nil, err
	}
	return &tmpl, nil
}

func NewMetric(name string, template string) error {
	if strings.HasPrefix(name, "uadmin/") && !SystemMetrics {
		return nil
	}
	if !strings.HasPrefix(name, "uadmin/") && !UserMetrics {
		return nil
	}
	tmpl, err := getRRDTemplate(template)
	if err != nil {
		Trail(ERROR, "NewMetric.getRRDTemplate. %s", err.Error())
		return err
	}
	fName := "rrd/" + name + ".rrd"
	err = rrd.CreateRRD(fName, *tmpl)
	if err != nil {
		Trail(ERROR, "NetMetric.rrd.CreateRRD. %s", err.Error())
	}
	return err
}

func SetMetric(name string, value float64) {
	go setRRDValue(name, value, "gauge")
}

func IncrementMetric(name string) {
	go setRRDValue(name, 1, "absolute")
}

func TimeMetric(name string, div float64, f func()) {
	sTime := time.Now()
	f()
	SetMetric(name, float64(time.Now().Sub(sTime).Nanoseconds())/div)
}

func setRRDValue(name string, value float64, tmpl string) error {
	var err error
	if strings.HasPrefix(name, "uadmin/") && !SystemMetrics {
		return nil
	}
	if !strings.HasPrefix(name, "uadmin/") && !UserMetrics {
		return nil
	}
	fName := "rrd/" + name + ".rrd"
	if _, err = os.Stat(fName); err != nil {
		err = NewMetric(name, tmpl)
		if err != nil {
			return err
		}
	}
	err = rrd.UpdateRRD(fName, 1, value)
	if err != nil {
		Trail(ERROR, "setRRDValue.rrd.UpdateRRD. %s", err.Error())
	}
	return err
}
