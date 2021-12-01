package watch

import (
	algorand "agent/algorand/api/v1/model"
	"agent/api/v1/model"
	. "agent/pkg/watch"
	"agent/publisher"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type AlgorandBlockWatchConf struct {
	Path string
}

type AlgorandBlockWatch struct {
	AlgorandBlockWatchConf
	Watch

	JsonLogWatch Watcher
	logCh        chan interface{}
}

func NewAlgorandBlockWatch(conf AlgorandBlockWatchConf, jsonLogWatch Watcher) *AlgorandBlockWatch {
	w := new(AlgorandBlockWatch)
	w.AlgorandBlockWatchConf = conf
	w.Watch = NewWatch()
	w.JsonLogWatch = jsonLogWatch
	w.logCh = make(chan interface{}, 1)
	return w
}

func (w *AlgorandBlockWatch) StartUnsafe() {
	w.Watch.StartUnsafe()

	if w.JsonLogWatch == nil {
		w.JsonLogWatch = NewJsonLogWatch(JsonLogWatchConf{Path: w.Path}, nil)
	}
	w.JsonLogWatch.Subscribe(w.logCh)
	go w.handleLogMessage()

	Start(w.JsonLogWatch)
}

func (w *AlgorandBlockWatch) Stop() {
	w.Watch.Stop()

	w.JsonLogWatch.Stop()
}

func (w *AlgorandBlockWatch) handleLogMessage() {
	for {
		select {
		case message := <-w.logCh:
			jsonMessage := message.(map[string]interface{})

			if val, ok := jsonMessage["Type"]; ok && val == "RoundConcluded" {
				round, ok := jsonMessage["Round"]
				if !ok {
					log.Errorln("[AlgorandBlockWatch] detected corrupt log message: missing field 'Round' on 'Type=RoundConcluded'")
					continue
				}

				newBlockMetric := algorand.NewBlockMetric{
					Metric: model.NewMetric(true),
					Round:  round.(uint64),
				}
				newBlockMetricJson, err := json.Marshal(newBlockMetric)
				if err != nil {
					log.Errorln("[AlgorandBlockWatch] failed to marshal new block metric: ", err)
					return
				}

				metric := publisher.Metric{
					Type: "protocol.new_block",
					Body: newBlockMetricJson,
				}

				w.Emit(metric)
			}

		case <-w.StopKey:
			return

		}
	}
}