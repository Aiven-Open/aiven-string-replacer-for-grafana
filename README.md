# aiven-grafana-string-replacer

## Installing

```
go install github.com/aiven/aiven-grafana-string-replacer
```

## Usage
```bash
Usage of aiven-grafana-string-replacer:
  -apikey string
    	Grafana api key (required)
  -overwrite
    	Overwrite dashboard on conflict (default true)
  -replace value
    	What to replace (key:value, multiple entries allowed, required)
  -retries int
    	Retries when using the grafana api (default 3)
  -uid string
    	Dashboard uid to process (required)
  -url string
    	Grafana url (required)
```
