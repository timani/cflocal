# CF Local - cf CLI Plugin

[![Linux Build Status](https://travis-ci.org/sclevine/cflocal.svg?branch=master)](https://travis-ci.org/sclevine/cflocal)
[![Windows Build Status](https://ci.appveyor.com/api/projects/status/tbaf399k1d60q78j/branch/master?svg=true)](https://ci.appveyor.com/project/sclevine/cflocal/branch/master)
[![GoDoc](https://godoc.org/github.com/sclevine/cflocal?status.svg)](https://godoc.org/github.com/sclevine/cflocal)

![CF Local Demo](https://raw.githubusercontent.com/sclevine/cflocal/master/assets/cflocal-demo.gif) \
*Note: Image download/build only occurs when a new rootfs is available.*

CF Local is a Cloud Foundry CLI plugin that enables you to:

* Stage and run Cloud Foundry apps in Docker
* Download pre-staged apps from Cloud Foundry and run them in Docker
* Stage apps in Docker and push them to Cloud Foundry
* Make Cloud Foundry service instances instantly available to Docker apps
* Share a local directory with an app running in Docker, and monitor it for changes
* Rapidly iterate on Cloud Foundry app development with only Docker installed
* Convert Cloud Foundry apps to Docker images that do not require Cloud Foundry or CF Local to run

Notably, CF Local:

* Does not require the Docker CLI
* Can run against a remote Docker daemon
* Uses the latest official Cloud Foundry buildpack releases by default
* Always uses the latest Cloud Foundry rootfs (cflinuxfs2) release

```
USAGE:
   cf local stage   <name> [ (-b <name> | -b <URL>) (-p <dir> | -p <zip>) ]
                           [ -m (-s <app> | -f <app>) ]
   cf local run     <name> [ (-i <ip>) (-p <port>) (-d <dir> [-w]) ]
                           [ (-s <app>) (-f <app>) ]
   cf local export  <name> [ (-r <ref>) ]
   cf local pull    <name>
   cf local push    <name> [-e -k]
   cf local help
   cf local version

STAGE OPTIONS:
   stage <name>   Build a droplet using the app in the current directory and
                     the environment variables and service bindings specified
                     in local.yml.
                     Droplet filename: <name>.droplet

   -b <name>      Use an official CF buildpack, specified by name.
                     Default: (uses detection)
   -b <url>       Use a buildpack specified by URL (git or zip-over-HTTP).
                     Default: (uses detection)
   -p <dir>       Use the specified directory as the app directory.
                     Default: current working directory
   -p <zip>       Use the specified zip file contents as the app directory.
                     Default: current working directory
   -d <dir>       Mount the provided directory into the app root before
                     staging. Should be used with a VCS to track changes.
                     Ideal for use with: cf local run <name> -d <dir>
                     Use of -r is STRONGLY RECOMMENDED to avoid staging with
                     files that should be ignored.
                     Default: (not mounted)
   -r             When used with -d, rsync any files that were created,
                     modified, or moved during staging into the specified
                     directory. The directory is mounted elsewhere in the
                     container. No files are deleted.
                     Default: false
   -s <app>       Use the service bindings from the specified remote CF app
                     instead of the service bindings in local.yml.
                     Default: (uses local.yml)
   -f <app>       Same as -s, but re-writes the service bindings to match
                     what they would be if they were tunneled through the app
                     with: cf local run <name> -f <app>
                     Default: (uses local.yml)

RUN OPTIONS:
   run <name>     Run a droplet with the configuration specified in local.yml.
                     Droplet filename: <name>.droplet

   -i <ip>        Listen on the specified interface IP
                     Default: localhost
   -p <port>      Listen on the specified port
                     Default: (arbitrary free port)
   -d <dir>       Mount the specified directory into the app root.
                     If empty, the app root is copied into the directory.
                     If not empty, the app root is replaced by the directory.
                     Default: (not mounted)
   -r             When used with -d, rsync the contents of the specified
                     directory into the container app root. The directory is
                     mounted elsewhere in the container. No files are deleted.
                     Default: false
   -w             When used with -d, restart the app when the contents of the
                     specified directory are changed.
                     Default: false
   -s <app>       Use the service bindings from the specified remote CF app
                     instead of the service bindings in local.yml.
                     Default: (uses local.yml or app provided by -f)
   -f <app>       Tunnel service connections through the specified remote CF
                     app. This re-writes the service bindings in the container
                     environment in order to use the tunnel. The service
                     bindings from the specified app will be used if -s is not
                     also passed.
                     Default: (uses local.yml)

EXPORT OPTIONS:
   export <name>  Export a standalone Docker image using the specified droplet
                     and configuration from local.yml.
                     Droplet filename: <name>.droplet

   -r <ref>       Tag the exported image with the provided reference.
                     Default: none

PULL OPTIONS:
   pull <name>    Download the droplet, environment variables, environment
                     variable groups, and start command of the named remote
                     CF app. The local.yml file is updated with the downloaded
                     configuration.
                     Droplet filename: <name>.droplet

PUSH OPTIONS:
   push <name>    Push a droplet to a remote CF app and restart the app.
                     Droplet filename: <name>.droplet

   -e             Additionally replace the remote app environment variables
                     with the environment variables from local.yml. This does
                     not read or replace environment variable groups.
                     Default: false
   -k             Do not restart the application after pushing the droplet.
                     The current droplet will continue to run until the next
                     restart.
                     Default: false

ENVIRONMENT:
   DOCKER_HOST    Docker daemon address
                     Default: /var/run/docker.sock

SAMPLE: local.yml

applications:
- name: first-app
  command: "some start command"
  memory: 2G
  disk_quota: 4G
  staging_env:
    SOME_STAGING_VAR: "some staging value"
  running_env:
    SOME_RUNNING_VAR: "some running value"
  env:
    SOME_VAR: "some value"
  services:
    (( VCAP_SERVICES object in YAML ))
```

## Install

```bash
$ ./cflocal-v0.13.0-macos
Plugin successfully installed. Current version: 0.13.0
```
***Or***
```bash
$ cf install-plugin cflocal-0.13.0-macos

**Attention: Plugins are binaries written by potentially untrusted authors. Install and use plugins at your own risk.**

Do you want to install the plugin cflocal-0.13.0-macos?> y

Installing plugin cflocal-0.13.0-macos...
OK
Plugin cflocal v0.13.0 successfully installed.
```
***Or***
```bash
$ cf install-plugin -r CF-Community cflocal

**Attention: Plugins are binaries written by potentially untrusted authors. Install and use plugins at your own risk.**

Do you want to install the plugin cflocal?> y
Looking up 'cflocal' from repository 'CF-Community'
11354404 bytes downloaded...
Installing plugin cflocal-0.13.0-macos...
OK
Plugin cflocal v0.13.0 successfully installed.
```
Note: The version available in the 'CF-Community' plugin repo may not always be the latest available.

## Uninstall

```
$ cf uninstall-plugin cflocal
Uninstalling plugin cflocal...
OK
Plugin cflocal successfully uninstalled.
```

## Security Notes

* Forwarded services (`-f`) are not reachable during staging
* Images are never exported with remote service credentials
* Service credentials from remote apps are never stored in local.yml
* CF Local should not be used to download untrusted Cloud Foundry applications
* CF Local is not intended for production use and is offered without warranty

## TODO

* Allow specification of resource limits
* Re-establish SSH tunnel if it closes
* Allow local buildpacks to be specified
* Permit specification of cflinuxfs2 version
* Add warnings about mismatched Docker client / server versions
