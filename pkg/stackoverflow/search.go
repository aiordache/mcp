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

		page := mcp.ParseInt(req, "page", 1)

		response, err := client.SearchAnswers(query, page)
		if err != nil {
			return nil, fmt.Errorf("failed search request: %w", err)
		}

		return mcputil.JSONToolResult("Stackoverflow Search", response)
	})

	stackOverflowGetAnswersTool := mcp.NewTool(
		"stackoverflow_get_answers_by_question_id",
		mcp.WithDescription("Retrieve contents of stackoverflow answers by question ID"),
		mcp.WithNumber(
			"questionID",
			mcp.Description("The question identifier to retrieve answers"),
			mcp.DefaultNumber(0),
		),
	)

	mcps.AddTool(stackOverflowGetAnswersTool, func(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		apiKey := os.Getenv("STACKOVERFLOW_API_KEY")
		if apiKey == "" {
			return nil, fmt.Errorf("STACKOVERFLOW_API_KEY environment variables is required")
		}
		// Initialize Stackoverflow client.
		client := client.NewClient(apiKey)

		// Extract parameters
		questionID := mcp.ParseInt(req, "questionID", 0)
		if questionID == 0 {
			return nil, fmt.Errorf("questionID parameter is required")
		}

		response, err := client.GetAnswers(questionID)
		if err != nil {
			return nil, fmt.Errorf("failed search request: %w", err)
		}

		return mcputil.JSONToolResult("Stackoverflow Search", response)
	})

	// getAnswerIDsByQuestionID := mcp.NewTool(
	// 	"stackoverflow_answer_ids_by_question_id",
	// 	mcp.WithDescription("Retrieve stackoverflow answers IDs given a question ID"),
	// 	mcp.WithNumber(
	// 		"questionID",
	// 		mcp.Description("The question_id (integer) to retrieve answer IDs"),
	// 		mcp.Required(),
	// 	),
	// )
	//
	// mcps.AddTool(getAnswerIDsByQuestionID, func(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 	apiKey := os.Getenv("STACKOVERFLOW_API_KEY")
	// 	if apiKey == "" {
	// 		return nil, fmt.Errorf("STACKOVERFLOW_API_KEY environment variables is required")
	// 	}
	// 	// Initialize Stackoverflow client.
	// 	client := client.NewClient(apiKey)
	//
	// 	// Extract parameters
	// 	questionID := mcp.ParseInt(req, "questionID", 0)
	// 	if questionID == 0 {
	// 		return nil, fmt.Errorf("question ID parameter is required")
	// 	}
	//
	// 	response, err := client.GetAnswerIDsByQuestionID(questionID)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed search request: %w", err)
	// 	}
	//
	// 	return mcputil.JSONToolResult("Stackoverflow Question's Answers", response)
	// })

	getAnswerTool := mcp.NewTool(
		"stackoverflow_get_answer_by_id",
		mcp.WithDescription("Retrieve stackoverflow answer body given an answer ID"),
		mcp.WithNumber(
			"answerID",
			mcp.Description("The answer_id (integer) to retrieve answer body"),
			mcp.Required(),
		),
	)

	mcps.AddTool(getAnswerTool, func(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		apiKey := os.Getenv("STACKOVERFLOW_API_KEY")
		if apiKey == "" {
			return nil, fmt.Errorf("STACKOVERFLOW_API_KEY environment variables is required")
		}
		// Initialize Stackoverflow client.
		client := client.NewClient(apiKey)

		// Extract parameters
		answerID, ok := req.Params.Arguments["answerID"].(int)
		if !ok || answerID == 0 {
			return nil, fmt.Errorf("query parameter is required")
		}

		response, err := client.GetAnswer(answerID)
		if err != nil {
			return nil, fmt.Errorf("failed search request: %w", err)
		}

		return mcputil.JSONToolResult("Stackoverflow Question's Answers", response)
	})
}
