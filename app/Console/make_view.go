package console

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	Register(Command{
		Name:        "make:view",
		Description: "Create a new HTML view file",
		Execute:     handleMakeView,
	})
}

func handleMakeView(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("view name is required. Usage: go askar make:view <view_name>")
	}

	name := args[0]
	if !strings.HasSuffix(name, ".html") {
		name = name + ".html"
	}

	// Create view file
	if err := createViewFile(name); err != nil {
		return err
	}

	fmt.Printf("\n✅ View created successfully!\n")
	fmt.Printf("   📄 resources/views/%s\n\n", name)

	return nil
}

func createViewFile(name string) error {
	viewsDir := "resources/views/pages"
	fileName := filepath.Join(viewsDir, name)

	// Check if file already exists
	if _, err := os.Stat(fileName); err == nil {
		return fmt.Errorf("view file already exists: %s", fileName)
	}

	// Ensure directory exists
	if err := os.MkdirAll(viewsDir, 0755); err != nil {
		return fmt.Errorf("failed to create views directory: %w", err)
	}

	template := `{{ define "pages/` + strings.TrimSuffix(name, ".html") + `" }}
{{ template "layouts/header" . }}

<main class="max-w-7xl mx-auto px-4 py-20">
    <h1 class="text-4xl font-bold gradient-text">` + strings.TrimSuffix(capitalizeFirst(name), ".html") + `</h1>
    <p class="mt-6 text-slate-400 text-lg">New page created in Askar.</p>
    
    <div class="mt-12 p-8 glass rounded-2xl">
        <h2 class="text-xl font-bold mb-4">Interactivity Demo</h2>
        <div x-data="{ count: 0 }">
            <button @click="count++" class="bg-blue-600 hover:bg-blue-500 text-white px-4 py-2 rounded-lg transition">
                Count is: <span x-text="count"></span>
            </button>
        </div>
    </div>
</main>

{{ template "layouts/footer" . }}
{{ end }}
`

	if err := os.WriteFile(fileName, []byte(template), 0644); err != nil {
		return fmt.Errorf("failed to create view file: %w", err)
	}

	return nil
}
