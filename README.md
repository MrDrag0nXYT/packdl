# packdl

Download utility for Bukkit/Spigot/Paper/etc-based server packs

## Warning

Currently project in **early development state**, so it could be unstable for production. **Always make backups and use at your own risc!**

## Features

- Simple CLI
- Pack metadata in JSON format
- Downloading by direct link with file SHA-1 hashsum validation
- Generating pack config file with parsing plugins metadata

## Future roadmap

- [ ] Replacing SpigotMC links with spiget API?
- [ ] Command-based pack management (add/remove mods)?
- [ ] ???

## Usage

```shell
packdl
  </path/to/pack> - download pack from config
  gen(generate) - generate config for pack
    -core core.jar - set server core jar file
    </path/to/pack>
  ver(version) - show version
```

Default pack config file name is `packdl-config.json`, packdl will search it when specified only directory path. If there's no path specified, will search config in current directory

# For developers

## Pack format

### Config overview:

```json
{
    "name": "Pack name",
    "version": "1.0.0",
    "author": "Author",
    "description": "Description",
    "website": "https://example.tld/pack",

    "core": {CoreObject},
    "plugins": [{ModificationObject}],
    "mods": [{ModificationObject}],
}
```

Fields `name`, `author`, `version` and one of `plugins` or `mods` are required!

### CoreObject

```json
{
    "name": "Leaf",
    "gameVersion": "1.21.4",
    "build": "525",
    "file": {FileObject}
}
```

### ModificationObject

```json
{
    "name": "LuckPerms",
    "version": "v5.5.55",
    "website": "https://luckperms.net/",
    "file": {FileObject}
}
```

### FileObject

```json
{
  "name": "leaf-1.21.4-525.jar",
  "downloadUrl": "https://api.leafmc.one/v2/projects/leaf/versions/1.21.4/builds/525/downloads/leaf-1.21.4-525.jar",
  "sha1": "74AA0767310589E588B32AD9E53A723A61A6A0C0"
}
```

## Building

For building you need installed Go (recommend latest version). Then run `build.sh` for Linux or `build.bat` for Windows

If you wand build it manually, download dependencies with:

```shell
go mod tidy
```

After install run:

```shell
go build -ldflags "-s -w"
```

For specify version add to ldflags `-X main.Version=REPLACEWITHVERSION`

## For website admins

HTTP client uses user agent `packdl/VERSION`
