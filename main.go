package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args

	if len(args) < 3 || args[1] != "create" {
		fmt.Println("❌ Penggunaan salah.")
		fmt.Println("Gunakan format:")
		fmt.Println("  generate-folder create <nama-proyek>")
		os.Exit(1)
	}

	projectName := args[2]
	baseDir, _ := filepath.Abs(projectName)

	folders := []string{
		"app",
		"app/assets",
		"app/css",
		"app/js",
	}

	files := map[string]string{
		"app/tes.html": fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
  <title>%s</title>
  <link rel="stylesheet" href="css/style.css">
</head>
<body>
  <h1>Welcome to %s</h1>
  <script src="js/script.js"></script>
</body>
</html>`, projectName, projectName),
		"app/css/style.css": "body { font-family: sans-serif; background: #fafafa; text-align: center; }",
		"app/js/script.js":  fmt.Sprintf("console.log('Script loaded for %s');", projectName),
	}

	// Buat folder proyek
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		os.MkdirAll(baseDir, 0755)
		fmt.Printf("📁 Folder proyek dibuat: %s\n", baseDir)
	} else {
		fmt.Printf("⚠️ Folder %s sudah ada. File akan ditimpa jika dibuat ulang.\n", projectName)
	}

	// Buat struktur folder
	for _, folder := range folders {
		folderPath := filepath.Join(baseDir, folder)
		os.MkdirAll(folderPath, 0755)
		fmt.Printf("✅ Folder dibuat: %s\n", folderPath)
	}

	// Buat file
	for path, content := range files {
		fullPath := filepath.Join(baseDir, path)
		os.WriteFile(fullPath, []byte(content), 0644)
		fmt.Printf("📝 File dibuat: %s\n", fullPath)
	}

	fmt.Printf("\n🎉 Proyek \"%s\" berhasil dibuat!\n", projectName)
}
