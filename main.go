package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "strconv"
)

// template structure for Postmark
type Template struct {
	TemplateID int64  `json:"templateID,omitempty"`
	Name       string `json:"name"`
	Subject    string `json:"subject"`
	HtmlBody   string `json:"htmlBody"`
	TextBody   string `json:"textBody"`
	Alias      string `json:"alias,omitempty"`
	Active     bool   `json:"active,omitempty"`
}

// TemplatesResponse represents the response from Postmark API for listing templates
type TemplatesResponse struct {
	Templates  []Template `json:"templates"`
	TotalCount int        `json:"totalCount"`
}

// PostmarkResponse represents the response from Postmark API for create/update actions
type PostmarkResponse struct {
	TemplateID int64  `json:"templateID"`
	ErrorCode  int    `json:"errorCode"`
	Message    string `json:"message"`
}

// CreateTemplate creates a new template in Postmark
func CreateTemplate(apiToken string, template Template) (int64, error) {
	url := "https://api.postmarkapp.com/templates"
	jsonData, err := json.Marshal(template)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal template: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Postmark-Server-Token", apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		return 0, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, bodyString)
	}

	var postmarkResponse PostmarkResponse
	if err := json.NewDecoder(resp.Body).Decode(&postmarkResponse); err != nil {
		return 0, fmt.Errorf("failed to decode response: %w", err)
	}

	if postmarkResponse.ErrorCode != 0 {
		return 0, fmt.Errorf("failed to create template: %s", postmarkResponse.Message)
	}

	return postmarkResponse.TemplateID, nil
}

// GetTemplate retrieves a template by its ID
func GetTemplate(apiToken string, templateID int64) (*Template, error) {
	url := fmt.Sprintf("https://api.postmarkapp.com/templates/%d", templateID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Postmark-Server-Token", apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		return nil, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, bodyString)
	}

	var template Template
	if err := json.NewDecoder(resp.Body).Decode(&template); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &template, nil
}

// UpdateTemplate updates an existing template in Postmark
func UpdateTemplate(apiToken string, template Template) error {
	url := fmt.Sprintf("https://api.postmarkapp.com/templates/%d", template.TemplateID)
	jsonData, err := json.Marshal(template)
	if err != nil {
		return fmt.Errorf("failed to marshal template: %w", err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Postmark-Server-Token", apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		return fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, bodyString)
	}

	var postmarkResponse PostmarkResponse
	if err := json.NewDecoder(resp.Body).Decode(&postmarkResponse); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if postmarkResponse.ErrorCode != 0 {
		return fmt.Errorf("failed to update template: %s", postmarkResponse.Message)
	}

	return nil
}

// DeleteTemplate deletes a template in Postmark
func DeleteTemplate(apiToken string, templateID int64) error {
	url := fmt.Sprintf("https://api.postmarkapp.com/templates/%d", templateID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Postmark-Server-Token", apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		return fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, bodyString)
	}

	return nil
}

// GetTemplates retrieves a list of all templates in Postmark
func GetTemplates(apiToken string, offset, count int) ([]Template, error) {
	url := fmt.Sprintf("https://api.postmarkapp.com/templates?offset=%d&count=%d", offset, count)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Postmark-Server-Token", apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		return nil, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, bodyString)
	}

	var templatesResponse TemplatesResponse
	if err := json.NewDecoder(resp.Body).Decode(&templatesResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return templatesResponse.Templates, nil
}

// ValidateTemplate validates a template in Postmark
func ValidateTemplate(apiToken string, template Template) error {
	url := "https://api.postmarkapp.com/templates/validate"
	jsonData, err := json.Marshal(template)
	if err != nil {
		return fmt.Errorf("failed to marshal template: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Postmark-Server-Token", apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		return fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, bodyString)
	}

	return nil
}

func main() {
	apiToken := 'YOUR_SEREVE_KEY'

	// Creating a template
	// template := Template{
	// 	Name:     "Password Reset",
	// 	Subject:  "Reset your password",
	// 	HtmlBody: "<html><body><p>Click <a href='{{reset_link}}'>here</a> to reset your password.</p></body></html>",
	// 	TextBody: "Click the following link to reset your password: {{reset_link}}",
	// }

	// templateID, err := CreateTemplate(apiToken, template)
	// if err != nil {
	//     log.Fatalf("Error creating template: %v\n", err)
	// }
	// fmt.Printf("Template created with ID: %d\n", templateID)

	// Getting a template
	// retrievedTemplate, err := GetTemplate(apiToken, templateID)
	// if err != nil {
	// 	log.Fatalf("Error getting template: %v\n", err)
	// }
	// fmt.Printf("Retrieved Template: %+v\n", retrievedTemplate)

	// Updating a template
	// retrievedTemplate.Subject = "Updated Subject"
	// err = UpdateTemplate(apiToken, *retrievedTemplate)
	// if err != nil {
	//     log.Fatalf("Error updating template: %v\n", err)
	// }
	// fmt.Println("Template updated successfully")

	// Listing templates
	fmt.Println("before deleting template id 36087615")
	offset := 0
	count := 10
	templates, err := GetTemplates(apiToken, offset, count)
	if err != nil {
		log.Fatalf("Error listing templates: %v\n", err)
	}
	fmt.Printf("Total Templates: %d\n", len(templates))
	for _, t := range templates {
		fmt.Printf("ID: %d, Name: %s, Active: %t, Subject: %s\n", t.TemplateID, t.Name, t.Active, t.Subject)
	}

	err = DeleteTemplate(apiToken, int64(36087634))
	if err != nil {
		log.Fatalf("Error deleting template: %v\n", err)
	}
	fmt.Println("Template deleted successfully")

	fmt.Println("After deleting template id 36087615")
	list, err := GetTemplates(apiToken, offset, count)
	if err != nil {
		log.Fatalf("Error listing templates: %v\n", err)
	}
	fmt.Printf("Total Templates: %d\n", len(list))
	for _, t := range list {
		fmt.Printf("ID: %d, Name: %s, Active: %t, Subject: %s\n", t.TemplateID, t.Name, t.Active, t.Subject)
	}

	// validating a template
	// err = ValidateTemplate(apiToken, template)
	// if err != nil {
	//     log.Fatalf("Error validating template: %v\n", err)
	// }
	// fmt.Println("Template validated successfully")
}
