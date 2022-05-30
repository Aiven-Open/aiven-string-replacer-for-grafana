# grafana-converter

## Installing

```
go install github.com/aiven/grafana-converter
```

## Usage
```bash
Usage of grafana-converter:
  -apikey string
    	Grafana api key (required)
  -overwrite
    	Overwrite dashboard on conflict (default true)
  -replace value
    	What to replace (key:value, multiple entries allowed, required)
  -retries int
    	Retries when grafana using the grafana api (default 3)
  -uid string
    	Dashboard uid to process (required)
  -url string
    	Grafana url (required)
```
