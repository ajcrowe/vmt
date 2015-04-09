# Vagrant Metadata Tool (VMT)

This tool allows easy creation, modification and deletion of versions from vagrant box metadata files. 

This tool is not complete and requires additional functionality in relation to providers and modification of existing metadata.

## Install

Run the build script

```
./build
```

Then copy the `bin/vmt` binary to somewhere in your `$PATH`

##  Configuration

You can configure a number of global parameters. These can live in a `.vmtrc` file in your home directory (see `vmtrc.sample`) or can be set with environment variables

#### Box URL 

config: `box_url`

envvar: `VMT_BOX_URL`

#### Box File Root

config: `box_file_root`

envvar: `VMT_BOX_FILE_ROOT`

#### Default Provider

config: `default_provider`

envvar: `VMT_DEFAULT_PROVIDER`

## Commands

### `vmt generate`

Description: Creates an initial metadata file

Flags:

```
--description, -d       description
--shortdescription, -s  short box description
--boxname, -b           name of the box
--output, -o            file to write metadata to
```

### `vmt version list`

Description: Lists the versions present in the specified metadata

Flags:

```
--input, -i         file to read metadata from
```

### `vmt version add`

Description: Adds a new version to an existing json file

Flags:

```
--quiet, -q         suppress output to stdout 
--noop, -n          run in no-op mode
--version, -v       box version
--description, -d   description
--input, -i         file to read metadata from
--output, -o        file to write metadata to
--boxfile, -f       path to the box file
--provider, -p      provider of the version
```

### `vmt version remove`

Description: Removes a version from the specified metadata

Flags:

```
--quiet, -q         suppress output to stdout
--version, -v       box version
--input, -i         file to read metadata from
--output, -o        file to write metadata to
--remove, -r        remove box file when deleting version
```

## Example

Lets generate a new json metadata file

```
$ vmt generate -o example.json -d "Vagrant Example Box" -s "examplebox vagrant" -b "examplebox"
$ cat example.json
{
  "name": "examplebox",
  "description": "Vagrant Example Box",
  "short_description": "examplebox vagrant",
  "versions": null
}
```

Now lets add our box as a version to this metdata. We'll create a fake box file to use

```
$ touch examplebox-0.1.box
$ vmt version add -i example.json -v 0.1 -f examplebox-0.1.box -d "Initial Version" -q
cat example.json
{
  "name": "examplebox",
  "description": "Vagrant Example Box",
  "short_description": "examplebox vagrant",
  "versions": [
    {
      "version": "0.1",
      "status": "active",
      "description_html": "\u003cp\u003eInitial Version\u003c/p\u003e",
      "description_markdown": "Initial Version",
      "providers": [
        {
          "name": "virtualbox",
          "url": "http://vagrantbox.example.com/examplebox-0.1.box",
          "checksum_type": "sha1",
          "checksum": "da39a3ee5e6b4b0d3255bfef95601890afd80709"
        }
      ]
    }
  ]
}
```

We can view a summary of the versions in the metadata with `version list`

```
$ vmt version list -i example.json
Version   Description       Status
0.1       Initial Version   active
```

Finally we can remove the version

```
$ vmt version remove -i example.json -v 0.1 -q
$ cat example.json
{
  "name": "examplebox",
  "description": "Vagrant Example Box",
  "short_description": "examplebox vagrant",
  "versions": null
}
```

## Contributing

Please feel free to modify and submit PRs.

## ToDo

* Allow for removal of specific providers in version
* Add error checking for all input flags
* Add ability to modify existing versions with addition providers
