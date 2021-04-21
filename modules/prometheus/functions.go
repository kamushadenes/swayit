package prometheus

import (
	"context"
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	pconfig "github.com/prometheus/common/config"
	"github.com/prometheus/common/model"
	"strconv"
	"strings"
	"time"
)

const (
	okColor  = "#2ECC40"
	nokColor = "#FF4136"
)

func hasIssue(m *Metric) bool {
	if sensitivity, ok := m.Labels["alert_sensitivity"]; ok {
		switch sensitivity {
		case "high":
			if m.Value < 95 {
				return true
			}
		case "medium":
			if m.Value < 90 {
				return true
			}
		case "low":
			if m.Value < 75 {
				return true
			}
		}
	}

	return false
}

func getMetrics(value model.Value) []*Metric {
	var metrics []*Metric
	lines := strings.Split(value.String(), "\n")

	for _, l := range lines {
		var m Metric
		m.Labels = make(map[string]string)
		m.Metric = strings.Split(l, "{")[0]
		m.Value, _ = strconv.ParseInt(strings.Fields(strings.Split(l, "=>")[1])[0], 10, 64)

		lblString := strings.Split(strings.Split(l, "{")[1], "}")[0]
		lvs := strings.Split(lblString, ",")
		for _, l := range lvs {
			l = strings.TrimSpace(l)
			fields := strings.Split(l, "=")
			m.Labels[fields[0]] = fields[1]
		}

		metrics = append(metrics, &m)
	}

	return metrics
}

func run(w *common.WaybarOutput) error {
	client, err := api.NewClient(api.Config{
		Address:      config.SwayItConfig.Prometheus.URL,
		RoundTripper: pconfig.NewBasicAuthRoundTripper(config.SwayItConfig.Prometheus.Username, pconfig.Secret(config.SwayItConfig.Prometheus.Password), "", api.DefaultRoundTripper),
	})
	if err != nil {
		return err
	}
	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _, err := v1api.Query(ctx, config.SwayItConfig.Prometheus.Query, time.Now())
	if err != nil {
		return err
	}

	metrics := getMetrics(result)

	results := make(map[string]map[string]*Metric)

	w.Text = fmt.Sprintf("<span color='%s'>\uf058</span>", okColor)

	for _, v := range metrics {
		if job, ok := v.Labels["job"]; ok {
			if instance, ok := v.Labels["instance"]; ok {
				if _, ok := results[job]; !ok {
					results[job] = make(map[string]*Metric)
				}
				results[job][instance] = v
				if hasIssue(v) {
					w.Text = fmt.Sprintf("<span color='%s'>\uf071</span>", nokColor)
				}
			}
		}
	}

	for job := range results {
		w.Tooltip += fmt.Sprintf("<b>%s</b>\n", job)

		for instance := range results[job] {
			if hasIssue(results[job][instance]) {
				w.Tooltip += fmt.Sprintf("\n%s: <b><span color='%s'>%d%%</span></b>", instance, nokColor, results[job][instance].Value)
			} else {
				w.Tooltip += fmt.Sprintf("\n%s: <span color='%s'>%d%%</span>", instance, okColor, results[job][instance].Value)
			}
		}

		w.Tooltip += "\n\n"
	}

	return nil
}
