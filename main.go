package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"regexp"

	gapi "github.com/grafana/grafana-api-golang-client"
)

type config struct {
	url      string
	apikey   string
	title    string
	replace  string
	by       string
	override bool
	retries  int
}

var cfg config

func main() {
	flag.StringVar(&cfg.url, "url", "", "Grafana url (required)")
	flag.StringVar(&cfg.apikey, "apikey", "", "Grafana api key (required)")
	flag.StringVar(&cfg.title, "title", "", "Title expression to process (required)")
	flag.StringVar(&cfg.replace, "replace", "", "String to replace (required)")
	flag.StringVar(&cfg.by, "by", "", "String to replace by (required)")
	flag.BoolVar(&cfg.override, "override", true, "Override dashboard on conflict")
	flag.IntVar(&cfg.retries, "retries", 3, "Retries when grafana using the grafana api")

	flag.Parse()

	if err := processDashboards(cfg); err != nil {
		log.Fatal("unable to process dashboards: ", err)
	}
	log.Println("processed all dashboards")
}

func processDashboards(cfg config) error {
	p, err := newProcessor(cfg)
	if err != nil {
		return fmt.Errorf("unable to create processor: %w", err)
	}

	if err := p.processDashboards(); err != nil {
		return fmt.Errorf("unable to process dashboards: %w", err)
	}
	return nil
}

type processor struct {
	cfg    config
	client *gapi.Client
}

func newProcessor(cfg config) (processor, error) {
	p := processor{}
	if cfg.url == "" {
		return p, errors.New("'url' is required")
	}
	if cfg.apikey == "" {
		return p, errors.New("'apikey' is required")
	}
	if cfg.title == "" {
		return p, errors.New("'title' is required")
	}
	if cfg.replace == "" {
		return p, errors.New("'replace' is required")
	}
	if cfg.by == "" {
		return processor{}, errors.New("'by' is required")
	}
	if _, err := regexp.Compile(cfg.title); err != nil {
		return processor{}, fmt.Errorf("'title' is a bad regex: %w", err)
	}
	client, err := gapi.New(cfg.url, gapi.Config{
		APIKey:     cfg.apikey,
		NumRetries: cfg.retries,
	})
	if err != nil {
		return processor{}, fmt.Errorf("unable to create grafana client: %w", err)
	}

	return processor{cfg: cfg, client: client}, nil
}

func (p processor) processDashboards() error {
	allDashboards, err := p.client.Dashboards()
	if err != nil {
		return fmt.Errorf("unable to fetch dashboards: %w", err)
	}

	for i := range allDashboards {
		if matches, _ := regexp.MatchString(cfg.title, allDashboards[i].Title); !matches {
			continue
		}
		log.Println("processing: ", allDashboards[i].Title)

		if err := p.processDashboardWithUID(allDashboards[i].UID); err != nil {
			return fmt.Errorf("unable to process dashboard %q: %w", allDashboards[i].Title, err)
		}
	}
	return nil
}

func (p processor) processDashboardWithUID(uid string) error {
	dashboard, err := p.client.DashboardByUID(uid)
	if err != nil {
		return fmt.Errorf("unable to fetch dashboard with uid %q: %w", uid, err)
	}
	dbytes, err := json.Marshal(dashboard.Model)
	if err != nil {
		return fmt.Errorf("unable to marshal model: %w", err)
	}
	dbytes = bytes.ReplaceAll(dbytes, []byte(cfg.replace), []byte(cfg.by))

	model := make(map[string]interface{})
	if err := json.Unmarshal(dbytes, &model); err != nil {
		return fmt.Errorf("unable to marshal processed model: %w", err)
	}
	dashboard.Model = model

	if _, err := p.client.SaveDashboard(dashboard.Model, p.cfg.override); err != nil {
		return fmt.Errorf("unable to save dashboard with uid %q: %w", uid, err)
	}

	return nil
}
