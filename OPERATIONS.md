# Reviewing an SLO before publishing its rules

A valid SLO specification can still measure the wrong user experience. This
checklist separates generator validation from operational validation.

## Start with the checked-in example

Read [getting-started.yml](examples/getting-started.yml). Its availability SLO
counts HTTP 5xx and 429 responses as errors and uses a 99.9% objective.
Decide whether those choices match your service before reusing the example.

With the Go toolchain required by [go.mod](go.mod), run from the repository root:

```sh
go run ./cmd/sloth validate -i examples/getting-started.yml
go run ./cmd/sloth generate -i examples/getting-started.yml
```

Generation prints rules to standard output; it does not install them in
Prometheus or Kubernetes. Go may download the toolchain and dependencies.
Do not edit files in examples/_gen directly.

## Review the measurement

- Confirm the metric name and job selector exist in your Prometheus.
- Confirm numerator and denominator cover the same traffic population.
- Decide how retries, cancellations and rate limits count toward user success.
- Preserve the window template in rate queries so generated windows remain valid.
- Treat zero traffic or missing telemetry separately from a healthy SLO.
- Verify ownership, routing labels and both page and ticket destinations.

## Common surprises

| Observation | Likely investigation |
| --- | --- |
| Validation succeeds, but rules return no data | Metric or selector does not match deployed instrumentation |
| Error ratio exceeds one | Numerator and denominator include different event populations |
| Alerts appear but nobody receives them | Alertmanager routing labels and receivers |
| Only one alert window has data | Retention, scrape gaps or recent metric creation |
| Generated output changes unexpectedly | Input, Sloth version and configured plugins |

Before promotion, compare generated rules with the previous version and evaluate
the actual queries in a lab Prometheus. A successful generator run is not an
end-to-end paging test. Use [Makefile](Makefile) for the project's test targets.

## Development note

This review guide was added with AI assistance. Upstream code, licenses and
contributor attribution remain unchanged.
