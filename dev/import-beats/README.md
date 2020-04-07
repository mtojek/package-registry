# import-beats

The script is responsible for importing existing beats modules and transforming
them into integration packages compatible with Elastic Package Registry (EPR).

The `import-beats` script depends on active Kibana instance, which is used to
migrate existing dashboards to a newer version.

## Usage

```bash
$ mage ImportBeats
```

... or using `go run` (no need to install `mage`):

```bash
$ go run dev/import-beats/*.go -help
  Usage of /var/folders/gz/dht4sjdx5w9f72knybys10zw0000gn/T/go-build249388773/b001/exe/agent:
    -beatsDir string
       Path to the beats repository (default "../beats")
    -ecsDir string
       Path to the Elastic Common Schema repository (default "../ecs")
    -euiDir string
       Path to the Elastic UI framework repository (default "../eui")
    -kibanaDir string
       Path to the kibana repository (default "../kibana")
    -kibanaHostPort string
       Kibana host and port (default "http://localhost:5601")
    -outputDir string
       Path to the output directory (default "dev/packages/beats")
    -skipKibana
       Skip storing Kibana objects
```

## Import all packages

1. Make sure that the following repositories have been fetched locally:
https://github.com/elastic/beats
https://github.com/elastic/ecs
https://github.com/elastic/eui
https://github.com/elastic/kibana
2. Make sure you've the `mage` tool installed.
3. Start Kibana server (make sure the endpoint is accessible: http://localhost:5601/)
4. Run the importing procedure with the following command:

```bash
$ mage ImportBeats
```

## How does the import procedure work

This section describes next steps of the `import-beats` script that are performed to build integration packages in
the output directory.

Keep in mind that the script doesn't clean previously created artifacts, so you may encounter leftovers (detached
dashboards, renamed ingest pipeline, etc.). If you need to preserve a clean state in the output directory (which is
versioned), remove its content before executing the script.

The script requires few repositories (Kibana, EUI, etc.) to be present, but doesn't require to execute any of build
targets. It depends only on the existing, version content, so simple `git clone` should be enough.

## Troubleshooting

*Importing process takes too long.*

While developeing, you can try to perform the migration with skipping migration of all Kibana objects,
as this is the most time consuming part of whole process:

```bash
$ SKIP_KIBANA=true mage ImportBeats
```
