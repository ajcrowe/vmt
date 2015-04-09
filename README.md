# Vagrant Metadata Tool (VMT)

This tool allows easy creation, modification and deletion of versions from vagrant box metadata files. 


## Install

Run the build script

```
# ./build
```

Then copy the `vmt` binary to somewhere in your `$PATH`

##  Usage

You can configure global parameters 

### Box URL 

config: `box_url`

envvar: `VMT_BOX_URL`

### Box File Root

config: `box_file_root`

envvar: `VMT_BOX_FILE_ROOT`

These can live in a `.vmtrc` file in your home directory (see `vmtrc.sample`) 

### `vmt generate`

Description: Creates an initial metadata file

Params:

```
--description, -d       description
--shortdescription, -s  short box description
--boxname, -b           name of the box
--output, -o            file to write metadata to
```

### `vmt version list`

Description: Lists the versions present in the specified metadata

Params:

```
--input, -i         file to read metadata from
```

### `vmt version add`

Description: Adds a new version to an existing json file

Params:

```
--quiet, -q         suppress output to stdout 
--noop, -n          run in no-op mode
--version, -v       box version
--description, -d   description
--input, -i         file to read metadata from
--output, -o        file to write metadata to
--boxfile, -f       path to the box file
```

### `vmt version remove`

Description: Removes a version from the specified metadata

Params:

```
--quiet, -q         suppress output to stdout
--version, -v       box version
--input, -i         file to read metadata from
--output, -o        file to write metadata to
--remove, -r        remove box file when deleting version
```

## Contributing

Please feel free to modify and submit PRs.