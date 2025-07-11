# MCP

This repository contains experimental [Model Context Protocol (or MCP)](https://modelcontextprotocol.io/introduction) servers for interacting with Algolia APIs. This MCP repo is something we're sharing for you to explore and experiment with. Feel free to use it, fork it, or build on top of it — but just know that it's not officially supported by Algolia and isn't covered under our SLA. We might update it, break it, or remove it entirely at any time. If you customize or configure things here, there's a chance that work could be lost. Also, using MCP in production could affect your Algolia usage.

We also have a Node.js MCP Server that you can find here: https://github.com/algolia/mcp-node.

If you have feedback or ideas (even code!), we'd love to hear it. Just know that we might use it to help improve our products. This project is provided "as is" and "as available," with no guarantees or warranties. To be super clear: MCP isn't considered an "API Client" for SLA purposes.

## Installation

First follow the [quick start](https://modelcontextprotocol.io/quickstart/user), which will install Claude Desktop and setup a sample Fileserver MCP server.  This is a great introduction to using MCP and will let you debug things using the official guide if there are issues.

## Setup the prototype Algolia MCP server

Requirements:

* Go (https://go.dev/doc/install)

### Clone the repo and build the server

Clone the repo, amd build the mcp server:

```shell
$ git clone git@github.com:algolia/mcp.git
$ cd mcp/cmd/mcp
$ go build
```
We need to have the full path of the built server binary:
```shell
$ pwd
/path/to/the/repo/cmd/mcp
```

__NOTE:__  When adding this command to your configuration, you must specify the binary along with the path (`/path/to/the/repo/cmd/mcp/mcp`)

### Update the settings to point to the new server

In Claude desktop edit the settings as per https://modelcontextprotocol.io/quickstart/user#2-add-the-filesystem-mcp-server and this time add the server definition for algolia (using the server path that you found earlier).

```json
{
   "mcpServers": {
      "algolia": {
         "command": "/path/to/the/repo/cmd/mcp/mcp",
         "env": {
            "ALGOLIA_APP_ID": "<APP_ID>",
            "ALGOLIA_INDEX_NAME": "<INDEX_NAME>",
            "ALGOLIA_API_KEY": "<API_KEY>",
            "ALGOLIA_WRITE_API_KEY": "<ADMIN_API_KEY>",  /* if you want to allow write operations, use your ADMIN key here */
            "MCP_ENABLED_TOOLS": "",  /* optional: specify which tools to enable (e.g., "search,collections") */
            "MCP_SERVER_TYPE": "stdio",  /* optional: server type, either "stdio" (default) or "sse". If not set, defaults to "stdio" */
            "MCP_SSE_PORT": "8080"  /* optional: port for SSE server, default is 8080 (only used when MCP_SERVER_TYPE is "sse") */
         }
      }
   }
}
```

By default, all available tools are enabled when MCP_ENABLED_TOOLS is empty or not set. If you want to enable only specific tools, you can set this variable to a comma-separated list of tool names. Available tools are: abtesting, analytics, collections, monitoring, querysuggestions, recommend, search, search_read, search_write, usage.

- `search`: Enables all search operations (both read and write)
- `search_read`: Enables only read operations (list indices, get settings, run queries, get objects)
- `search_write`: Enables only write operations (clear, copy, delete, move, set settings, delete objects, insert objects)

Restart Claude desktop, and you should see a new `"algolia"` tool is available.

## Debugging

You can run the Inspector (see https://modelcontextprotocol.io/docs/tools/inspector) to check the MCP features and run them manually.

From the repo root, setup the environment

```shell
$ export ALGOLIA_APP_ID=""
$ export ALGOLIA_INDEX_NAME=""
$ export ALGOLIA_API_KEY=""
$ export ALGOLIA_WRITE_API_KEY=""  # if you want to allow write operations, use your ADMIN key here
$ export MCP_ENABLED_TOOLS=""  # if you want to restrict the tools activated you can optionally specify a list
$ export MCP_SERVER_TYPE="stdio"  # optional: server type, either "stdio" (default) or "sse". If not set, defaults to "stdio"
$ export MCP_SSE_PORT="8080"  # optional: port for SSE server, default is 8080 (only used when MCP_SERVER_TYPE is "sse")
```
Move into the server directory, and rebuild (if necessary):
```shell
$ cd cmd/mcp
$ go build # might already be up-to-date
```
Run the MCP inspector on the server:
```shell
$ npx @modelcontextprotocol/inspector ./mcp
```

## Using with Ollama

You can actually run a local mcphost (which orchestrates the MCP servers for you), and then use them with other models locally via Ollama.

We are using https://github.com/mark3labs/mcphost for this.

As per the [README](https://github.com/mark3labs/mcphost?tab=readme-ov-file#installation-) you need a a config file, so you can copy the Claude one, and put it somewhere sensible so you can use it on the command line (for example `~/mcp.json`)

```json filename="~/mcp.json"
{
   "mcpServers": {
      "algolia": {
         "command": "/path/to/the/repo/cmd/mcp/mcp",
         "env": {
            "ALGOLIA_APP_ID": "<APP_ID>",
            "ALGOLIA_INDEX_NAME": "<INDEX_NAME>",
            "ALGOLIA_API_KEY": "<API_KEY>",
            "MCP_ENABLED_TOOLS": "",  /* optional: specify which tools to enable (e.g., "search,collections") */
            "MCP_SERVER_TYPE": "stdio",  /* optional: server type, either "stdio" (default) or "sse". If not set, defaults to "stdio" */
            "MCP_SSE_PORT": "8080"  /* optional: port for SSE server, default is 8080 (only used when MCP_SERVER_TYPE is "sse") */
         }
      }
   }
}
```

By default, all available tools are enabled when MCP_ENABLED_TOOLS is empty or not set. If you want to enable only specific tools, you can set this variable to a comma-separated list of tool names. Available tools are: abtesting, analytics, collections, monitoring, querysuggestions, recommend, search, search_read, search_write, usage.
You can now run it directly (no need to check out the repo):
```shell
$ go run github.com/mark3labs/mcphost@latest --config ~/mcp.json -m ollama:qwen2.5:3b
```

# FAQ
### What sort of things can I do once I install this MCP?
Here are some sample prompts to seed your imagination:
   * “Search all products in the index where brand = ‘Nike’ and price < 100.”
   * “Add this JSON object to the blog_posts index.”
   * “Update the searchable attributes for the recipes index to include ingredients.”
   * “Configure my index to rank nebula_award winners higher”
### Where can I read more about the Algolia MCP server?
Check out [this blog post](https://www.algolia.com/blog/engineering/algolia-mcp-server) including an embedded demo video.
### Resource templates and root are not supported by Claude desktop right now

[This is a weird one](https://github.com/orgs/modelcontextprotocol/discussions/136), since there is a bunch of content online showing the templates, maybe it's just not GA yet.



# Run in a Docker container

### Build docker image
In the root directory of the repository, run:
 ```
 $ docker build -t mcp .
 ```
 Check the image has been created:
```
 $ docker images | grep mcp
  mcp        latest          6bb2b1324b68   2 hours ago     1GB
```

Run the MCP container with the 8080 mapped on the host and auth credentials passed as environment variables (transport must be set to sse):
```
$ docker run -it -p8080:8080 -e MCP_SERVER_TYPE=sse -e MCP_SSE_PORT=8080 -e ALGOLIA_APP_ID=<APP-ID> -e ALGOLIA_API_KEY=<API-KEY> -e ALGOLIA_INDEX_NAME=<index> -e 'STACKOVERFLOW_API_KEY=<SO_API_KEY>' mcp
```

The URL of the MCP server to be set in Claude desktop: `http:127.0.0.1:8080/sse`

Claude config sample:
```
{
  "mcpServers": {
    "Algolia MCP": {
      "command": "npx",
      "args": [
        "mcp-remote",
        "http://127.0.0.1:8080/sse",
        "--header",
        "Authorization: Bearer <APP-ID>:API-KEY>"
      ]
    }
  }
}
```



