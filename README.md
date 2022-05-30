# aiven-grafana-string-replacer

## Reason

`aiven-grafana-string-replacer` is a small tool that replaces strings in marshalled dashboards.
It is handy if you want to replace a metric expression that occurs more than once or if a metric was renamed upstream.

## Installing

```bash
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
    	What to replace (key<=>value, multiple entries allowed, required)
  -retries int
    	Retries when using the grafana api (default 3)
  -uid string
    	Dashboard uid to process (required)
  -url string
    	Grafana url (required)
```


For example, if you wish to rename metrics that start with `elasticsearch_` to metrics that start with `opensearch_`:

```bash
aiven-grafana-string-replacer -apikey [...] -url https://my-grafana.org/ -replace "elasticsearch_<=>opensearch_" -uid [...]
```
