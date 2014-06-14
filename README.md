# Stated

This is an _experiment_ to create a pluggable configuration management
system using [Golang](http://golang.org).  Stated works inside a
single machine, and there is no remote execution.  The project is in
**pre-alpha** stage.

Stated use [TOML](https://github.com/BurntSushi/toml) as the file
format for configuration with [Go
template](http://andlabs.lostsig.com/blog/2014/05/26/8/the-go-templates-post).

## Plugins

Stated provides a plugin mechanism built on top of RPC.  There will be
plugins available by default called _built-in plugins_.  Users can
create their own plugins using the API provided by the Stated.  User
created plugins is called _external plugins_.

These are the few plugin types planned for Stated:

- file
- package
- service
- exec
- user
- group
- schedule
- external

The section title defines the plugin type used.  For `external` plugin
types, a field named `plugin` points to the executable plugin.  The execuable
location could be an abosulte path, relative path or a command
available in the `PATH` environment variable.

## Example

`machine.state` content:

```
machine = "{{ env.STATED_MACHINE_DOMAIN }}"
start = ["etc_hosts", "etc_hosts_sample"]

[[file]]
name = "etc_hosts"
next = ["hello_world", "hello_stated"]
destination = "etc/hosts"
source = "files/hosts"
mode = 644

[[exec]]
name = "hello_world"
command = "echo 'Hello, World!' > out/hello-world.txt"

[[exec]]
name = "hello_stated"
command = "echo 'Hello, Stated!' > out/hello-stated.txt"

[[external]]
name = "etc_hosts_sample"
plugin = "stated-plugin-sample"
config = "etc_hosts_sample.state"
```

`etc_hosts_sample.state` content:

```
destination = "etc/hosts-sample"
template = "templates/hosts"
mode = 644
an_argument = "expected by the plugin"
```

`files/hosts` content:

```
127.0.0.1	localhost
127.0.1.1	example.org
```

`templates/hosts` content:

```
127.0.0.1	localhost
127.0.1.1	{{ env.STATED_MACHINE_DOMAIN }}
```

## Data Collectors

Stated provides a mechanism to retrieve data from various sources and
provide it to the state files.  The data will be available for
configuration templates as well as the plugins.  These are the few
planned data sources:

- environment variables (env)
- data derived from the underlying system (sys)
- External service discovery tools (etcd, consul)

## Features planned

- Section inheritance
- Concurrent execution
