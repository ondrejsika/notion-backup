package backup

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/ondrejsika/notion-backup/lib/client"
	"github.com/ondrejsika/notion-backup/lib/utils/download_utils"
	"github.com/ondrejsika/notion-backup/lib/utils/file_utils"
)

func BackupAll(token, spaceID, name, dir string) {
	date := time.Now().Format("2006-01-02_15-04-05")
	Backup(token, spaceID, "html",
		filepath.Join(dir, "notion-"+name+"-"+date+"-html.zip"))
	Backup(token, spaceID, "markdown",
		filepath.Join(dir, "notion-"+name+"-"+date+"-md.zip"))
}

func Backup(token, spaceID, exportType, path string) {
	url, _ := GetExportURL(token, spaceID, exportType)
	file_utils.EnsureDir(path)
	download_utils.DownloadFile(path, url)
}

func GetExportURL(token, spaceID, exportType string) (string, error) {
	cli := client.New(token)
	taskID, _ := cli.ExportSpace(spaceID, exportType)
	for {
		completed, pagesExported, exportURL, _ := cli.GetTasksExportSpace(taskID)
		if completed {
			return exportURL, nil
		}
		fmt.Printf("Exported %d pages (%s)\n", pagesExported, taskID)
		time.Sleep(5 * time.Second)
	}
}
