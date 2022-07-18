package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"

	gapi "github.com/grafana/grafana-api-golang-client"
	"github.com/wI2L/jsondiff"
)

type config struct {
	url       string
	apikey    string
	uid       string
	from      string
	to        string
	overwrite bool
	retries   int
	dry       bool
}

var cfg config

func main() {
	flag.StringVar(&cfg.url, "url", "", "Grafana url (required)")
	flag.StringVar(&cfg.apikey, "apikey", "", "Grafana api key (required)")
	flag.StringVar(&cfg.uid, "uid", "", "Dashboard uid to process (required)")
	flag.StringVar(&cfg.from, "from", "", "Replace from (required)")
	flag.StringVar(&cfg.to, "to", "", "Replace to (required)")
	flag.BoolVar(&cfg.overwrite, "overwrite", true, "Overwrite dashboard on conflict")
	flag.BoolVar(&cfg.dry, "dry", false, "Just show diffs without saving")
	flag.IntVar(&cfg.retries, "retries", 3, "Retries when grafana the grafana api")

	flag.Parse()

	if err := checkConfig(cfg); err != nil {
		log.Fatal("bad config: ", err)
	}
	if err := processDashboard(cfg); err != nil {
		log.Fatal("unable to process dashboard: ", err)
	}

	log.Println("processed dashboard")
}

func processDashboard(cfg config) error {
	client, err := gapi.New(cfg.url, gapi.Config{
		APIKey:     cfg.apikey,
		NumRetries: cfg.retries,
	})
	if err != nil {
		return fmt.Errorf("unable to create grafana client: %w", err)
	}

	dashboard, err := client.DashboardByUID(cfg.uid)
	if err != nil {
		return fmt.Errorf("unable to fetch dashboards: %w", err)
	}
	dbytes, err := json.Marshal(dashboard.Model)
	if err != nil {
		return fmt.Errorf("unable to marshal model: %w", err)
	}
	dbytes = bytes.ReplaceAll(dbytes, []byte(cfg.from), []byte(cfg.to))

	replaced := make(map[string]interface{})
	if err := json.Unmarshal(dbytes, &replaced); err != nil {
		return fmt.Errorf("unable to marshal processed model: %w", err)
	}
	orig := dashboard.Model

	if cfg.dry {
		ops, _ := jsondiff.Compare(orig, replaced)
		for i := range ops {
			fmt.Println(ops[i].String())
		}
		return nil
	} else {
		dashboard.Model = replaced
		dashboard.Overwrite = cfg.overwrite
		dashboard.Message = replacerMessage(cfg)
		if _, err := client.NewDashboard(*dashboard); err != nil {
			return fmt.Errorf("unable to save dashboard: %w", err)
		}
	}

	return nil
}

func replacerMessage(cfg config) string {
	const (
		prefix = "aiven-string-replacer-for-grafana"
	)
	return fmt.Sprintf("%s: %s<=>%s", prefix, cfg.from, cfg.to)
}

func checkConfig(cfg config) error {
	if cfg.url == "" {
		return errors.New("'url' is required")
	}
	if cfg.apikey == "" {
		return errors.New("'apikey' is required")
	}
	if cfg.uid == "" {
		return errors.New("'uid' is required")
	}
	if cfg.from == "" {
		return errors.New("'from' is required")
	}
	if cfg.to == "" {
		return errors.New("'to' is required")
	}
	return nil
}
