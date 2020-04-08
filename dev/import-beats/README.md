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

## Package import procedure

This section describes next steps of the `import-beats` script that are performed to build integration packages in
the output directory.

Keep in mind that the script doesn't clean previously created artifacts, so you may encounter leftovers (detached
dashboards, renamed ingest pipeline, etc.). If you need to preserve a clean state in the output directory (which is
versioned), remove its content before executing the script.

The script requires few repositories (Kibana, EUI, etc.) to be present, but doesn't require to execute any of build
targets. It depends only on the existing, version content, so simple `git clone` should be enough.

### Package Repository

The package repository is responsible for building packages - loading package data from sources (Beats modules, Kibana
resources, etc.) and writing them to disk. It supports two types of beats - logs and metrics.

#### Load input data from sources

The script needs to visit and process input data from [beats](https://github.com/elastic/beats), generally logs and
metrics modules.

Starting with modules, it collects and processes information about module fields, release type, icons, screenshots,
Kibana dashboards and docs. While browsing datasets content, it focuses on fields specific for the dataset, release
type, ingestion pipeline, stream and agent configuration.

##### Fields

Fields are extracted from `fields.yml` files and divided into 3 buckets - ECS fields, module fields
and package fields.

##### Integration title

The correct spelling makes better impression on users, so the scripts uses `title` property in the module fields
as the proper form. Remember to adjust this value if working on the migration from Beats.

##### Release type

Values: _beta, experimental, ga_

The value depends on definitions in module and dataset fields. The scripts determines the correct release type
for dataset, depending on overall release status for module (e.g. dataset can't be annotated as GA if the entire module
is in beta).

##### Images

The script supports two kinds of images - icons and screenshots. Even though they're stored in different media formats,
they analyzed to prepare a metadata information (title, size, media type).

###### Icons

The icons are loaded from the following sources: Kibana home tutorials and Elastic UI. Icons must be in SVG format and
have defined dimensions (information stored in manifest, used by Kibana). Keep in mind that only icon files referenced
in tutorials are processed.

###### Screenshots

The script parses module docs to find and collect all references to screenshots presenting Kibana dashboards.

##### Kibana dashboards

TODO

##### Documentation

TODO

#### Write package content to disk

TODO

## Troubleshooting

### Importing process takes too long

While developeing, you can try to perform the migration with skipping migration of all Kibana objects,
as this is the most time consuming part of whole process:

```bash
$ SKIP_KIBANA=true mage ImportBeats
```
