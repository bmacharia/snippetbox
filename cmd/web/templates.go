// In a real world application multiple pieces of dynamic data
// tht I want to display in the same page
// a lightweight way to pass data to the template is to use a dynamic struct which
// acts like a single structure for the data

package main

import "snippetbox.bmacharia/internal/models"

// define a templateData type to act as a holding structure for any dynamic data that I want to pass to my HTML templates
type templateData struct {
	Snippet models.Snippet
}
