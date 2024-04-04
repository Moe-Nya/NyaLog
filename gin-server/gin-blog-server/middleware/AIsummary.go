package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type AIData struct {
	Status       string
	Text         string
	TotalTokens  int64
	OutputTokens int64
	InputTokens  int64
}

func GPT(url string, apiKey string, model string, question string, language string) (data AIData) {
	// 构建返回数据
	var AIData AIData
	requestBody := strings.NewReader(`{
		    "model": "` + model + `",
		   "messages": [{"role": "user", "content": "这是一篇文章:` + question + `。请你使用` + language + `写一则简短的文章摘要。"}]
		  }`)
	// 创建POST请求
	req, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return AIData
	}

	// 设置请求的header
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return AIData
	}
	defer resp.Body.Close()

	// 获取响应的状态码
	AIData.Status = resp.Status

	// 读取响应的内容
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return AIData
	}
	// 读取AI的回复
	var respContent map[string]interface{}
	err = json.Unmarshal([]byte(string(responseBody)), &respContent)
	if err != nil {
		return AIData
	}
	respText := ((respContent["choices"]).([]interface{}))[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	AIData.Text = respText

	// 读取token状态
	respTotalTokens := (respContent["usage"]).(map[string]interface{})["total_tokens"].(float64)
	respOutputTokens := (respContent["usage"]).(map[string]interface{})["completion_tokens"].(float64)
	respInputTokens := (respContent["usage"]).(map[string]interface{})["prompt_tokens"].(float64)
	AIData.TotalTokens = int64(respTotalTokens)
	AIData.OutputTokens = int64(respOutputTokens)
	AIData.InputTokens = int64(respInputTokens)
	return AIData
}

func QianWen(url string, apiKey string, model string, question string, language string) (data AIData) {
	// 构建返回数据
	var AIData AIData

	// 构建请求的body数据
	requestBody := map[string]interface{}{
		"model": model,
		"input": map[string]interface{}{
			"prompt": "这是一篇文章:" + question + "。请你使用" + language + "写一则简短的文章摘要。",
		},
		"parameters": map[string]interface{}{
			"enable_search": true,
		},
	}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling request body:", err)
		return AIData
	}
	// 创建POST请求
	req, err := http.NewRequest("POST", url, strings.NewReader(string(requestBodyBytes)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return AIData
	}

	// 设置请求的header
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return AIData
	}
	defer resp.Body.Close()

	// 获取响应的状态码
	AIData.Status = resp.Status

	// 读取响应的内容
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return AIData
	}
	// 读取AI的回复
	var respContent map[string]interface{}
	err = json.Unmarshal([]byte(string(responseBody)), &respContent)
	if err != nil {
		return AIData
	}
	respText := respContent["output"].(map[string]interface{})["text"].(string)
	AIData.Text = respText

	// 读取token状态
	respTotalTokens := respContent["usage"].(map[string]interface{})["total_tokens"].(float64)
	respOutputTokens := respContent["usage"].(map[string]interface{})["output_tokens"].(float64)
	respInputTokens := respContent["usage"].(map[string]interface{})["input_tokens"].(float64)
	AIData.TotalTokens = int64(respTotalTokens)
	AIData.OutputTokens = int64(respOutputTokens)
	AIData.InputTokens = int64(respInputTokens)
	return AIData
}
