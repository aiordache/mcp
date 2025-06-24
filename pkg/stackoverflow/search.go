package stackoverflow

import (
	"context"
	"fmt"
	"os"

	"github.com/algolia/mcp/pkg/mcputil"
	"github.com/algolia/mcp/pkg/stackoverflow/client"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterAll Stackoverflow Search tool with the MCP server.
func RegisterAll(mcps *server.MCPServer) {
	stackoverflowSearchTool := mcp.NewTool(
		"stackoverflow_search",
		mcp.WithDescription("Retrieve stackoverflow questions and answers"),
		mcp.WithString(
			"query",
			mcp.Description("The string to search for in Stackoverflow questions and answers"),
			mcp.Required(),
		),
		mcp.WithNumber(
			"page",
			mcp.Description("The requested page of results (default is 1)"),
			mcp.DefaultNumber(1),
		),
	)

	mcps.AddTool(stackoverflowSearchTool, func(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		apiKey := os.Getenv("STACKOVERFLOW_API_KEY")
		if apiKey == "" {
			return nil, fmt.Errorf("STACKOVERFLOW_API_KEY environment variables is required")
		}
		// Initialize Stackoverflow client.
		client := client.NewClient(apiKey)

		// Extract parameters
		query, ok := req.Params.Arguments["query"].(string)
		if !ok || query == "" {
			return nil, fmt.Errorf("query parameter is required")
		}
		page := 1
		pageNum, ok := req.Params.Arguments["page"]
		if ok {
			page = pageNum.(int)
		}

		response, err := client.Search(query, page)
		if err != nil {
			return nil, fmt.Errorf("failed search request: %w", err)
		}

		return mcputil.JSONToolResult("Stackoverflow Search", response)
	})
}
