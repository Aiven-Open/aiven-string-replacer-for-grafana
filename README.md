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
  -dry
    	Just show diffs without saving
  -from string
    	Replace from (required)
  -overwrite
    	Overwrite dashboard on conflict (default true)
  -retries int
    	Retries when grafana the grafana api (default 3)
  -to string
    	Replace to (required)
  -uid string
    	Dashboard uid to process (required)
  -url string
    	Grafana url (required)
```


For example, if you wish to rename metrics that start with `elasticsearch_` to metrics that start with `opensearch_`:

```bash
aiven-grafana-string-replacer -apikey [...] -url https://my-grafana.org/ -from elasticsearch_ -to opensearch_ -uid [...]
```

## License
`aiven-grafana-string-replacer` is licensed under the Apache license, version 2.0. Full license text is available in the LICENSE file.

Please note that the project explicitly does not require a CLA (Contributor License Agreement) from its contributors.

## Contact
Bug reports and patches are very welcome, please post them as GitHub issues and pull requests at https://github.com/aiven/aiven-grafana-string-replacer . To report any possible vulnerabilities or other serious issues please see our security policy.
