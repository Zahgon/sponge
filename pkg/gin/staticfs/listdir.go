package staticfs

import (
	"time"

	"github.com/gin-gonic/gin"
)

// ListDirOption set options.
type ListDirOption func(*listDirOptions)

type listDirOptions struct {
	prefixPath     string
	enableDownload bool // default: false
	enableFilter   bool // default: true
	middlewares    []gin.HandlerFunc
}

func (o *listDirOptions) apply(opts ...ListDirOption) { _ = "STUB: not implemented"; return }

func defaultListDirOptions() *listDirOptions { _ = "STUB: not implemented"; return nil }

// WithListDirPrefixPath sets prefix path.
func WithListDirPrefixPath(prefixPath string) ListDirOption {
	_ = "STUB: not implemented"
	return *new(ListDirOption)
}

// WithListDirDownload enables download feature.
func WithListDirDownload() ListDirOption { _ = "STUB: not implemented"; return *new(ListDirOption) }

// WithListDirFilter enables file filter feature.
func WithListDirFilter(enable bool) ListDirOption {
	_ = "STUB: not implemented"
	return *new(ListDirOption)
}

// WithListDirFilesFilter sets file name filter.
func WithListDirFilesFilter(filters ...string) ListDirOption {
	_ = "STUB: not implemented"
	return *new(ListDirOption)
}

// WithListDirDirsFilter sets directory name filter.
func WithListDirDirsFilter(filters ...string) ListDirOption {
	_ = "STUB: not implemented"
	return *new(ListDirOption)
}

// WithListDirMiddlewares sets middlewares.
func WithListDirMiddlewares(middlewares ...gin.HandlerFunc) ListDirOption {
	_ = "STUB: not implemented"
	return *new(ListDirOption)
}

// -------------------------------------------------------------------------------------------

// FileInfo is a struct that represents a file or directory in the file system.
type FileInfo struct {
	Name    string    `json:"name"`
	Path    string    `json:"path"`
	IsDir   bool      `json:"is_dir"`
	Size    int64     `json:"size,omitempty"`
	ModTime time.Time `json:"mod_time,omitempty"`
}

// default file filters
var sensitiveDirs = []string{"/proc", "/sys", "/dev", "/run", "/boot", "/root", "/etc"}
var sensitiveFiles = []string{".git", ".env", ".DS_Store"}

func isAllowedPath(p string, enableFilter bool) bool { _ = "STUB: not implemented"; return false }

// nolint
func formatSize(size int64) string { _ = "STUB: not implemented"; return "" }

func listDirectory(dir string, enableFilter bool) ([]FileInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sortFiles(files []FileInfo, sortBy, order string) { _ = "STUB: not implemented"; return }

// default: desc

// name

func toggleOrder(current string) string { _ = "STUB: not implemented"; return "" }

func badRequestData(data any) gin.H { _ = "STUB: not implemented"; return *new(gin.H) }

func handleList(prefixPath string, o *listDirOptions) gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}

// Default display of 20 files per page.

// Calculate pagination information

// Pagination

// Retrieve the file of the current page

// Template FuncMap

func handleDownload(c *gin.Context) { _ = "STUB: not implemented"; return }

func handleAPIList(enableFilter bool) gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}

// ListDir registers the routes for serving static files.
func ListDir(r *gin.Engine, opts ...ListDirOption) { _ = "STUB: not implemented"; return }

// nolint
var htmlTextSrc = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Directory Listing: {{.Dir}}</title>
    <style>
        :root {
            --primary-color: #3498db;
            --secondary-color: #2980b9;
            --background-color: #f8f9fa;
            --card-color: #ffffff;
            --text-color: #333333;
            --border-color: #e0e0e0;
            --hover-color: #f1f7fc;
            --folder-color: #f39c12;
            --file-color: #7f8c8d;
            --header-bg: #f5f7fa;
            --sort-indicator-color: #3498db;
        }

        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
        }

        body {
            background-color: var(--background-color);
            color: var(--text-color);
            line-height: 1.1;
            padding: 20px;
            max-width: 1200px;
            margin: 0 auto;
        }

        .container {
            background-color: var(--card-color);
            border-radius: 8px;
            box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
            padding: 30px;
        }

        .header {
            display: flex;
            align-items: center;
            justify-content: space-between;
            margin-bottom: 20px;
            padding-bottom: 15px;
            border-bottom: 1px solid var(--border-color);
        }

        h1 {
            font-size: 24px;
            font-weight: 500;
            color: var(--primary-color);
        }

        .path-display {
            background-color: rgba(52, 152, 219, 0.1);
            padding: 10px 15px;
            border-radius: 6px;
            margin-bottom: 20px;
            overflow-x: auto;
            white-space: nowrap;
            font-family: monospace;
            font-size: 14px;
            border-left: 4px solid var(--primary-color);
        }

        .back-link {
            display: inline-flex;
            align-items: center;
            color: var(--primary-color);
            text-decoration: none;
            font-weight: 500;
            padding: 8px 16px;
            border-radius: 4px;
            transition: all 0.2s ease;
            margin-bottom: 20px;
            border: 1px solid var(--primary-color);
        }

        .back-link:hover {
            background-color: var(--primary-color);
            color: white;
        }

        .back-icon {
            margin-right: 8px;
        }

        table {
            width: 100%;
            border-collapse: collapse;
            margin-top: 10px;
            table-layout: fixed;
        }

        th {
            background-color: var(--header-bg);
            text-align: left;
            padding: 12px 15px;
            font-weight: 600;
            color: var(--text-color);
            border-bottom: 2px solid var(--border-color);
            position: sticky;
            top: 0;
        }

        td {
            padding: 12px 15px;
            border-bottom: 1px solid var(--border-color);
            white-space: nowrap;
            overflow: hidden;
            text-overflow: ellipsis;
        }

        tr:hover {
            background-color: var(--hover-color);
        }

        th a {
            color: var(--text-color);
            text-decoration: none;
            align-items: center;
            justify-content: space-between;
        }

        th a:hover {
            color: var(--primary-color);
        }

        .sort-indicator {
            color: var(--sort-indicator-color);
            font-weight: bold;
            margin-left: 5px;
        }

        .file-link {
            display: flex;
            align-items: center;
            text-decoration: none;
            color: var(--text-color);
        }

        .file-link:hover {
            color: var(--primary-color);
        }

        .file-text {
            display: flex;
            align-items: center;
            color: var(--text-color);
        }

        .file-icon {
            margin-right: 10px;
            font-size: 18px;
        }

        .folder-icon {
            color: var(--folder-color);
        }

        .file-icon-regular {
            color: var(--file-color);
        }

        .size-cell {
            color: #666;
            font-size: 0.9em;
        }

        .date-cell {
            color: #666;
            font-size: 0.9em;
        }

        .empty-message {
            text-align: center;
            padding: 30px;
            color: #7f8c8d;
            font-style: italic;
        }

        .pagination {
            display: flex;
            justify-content: center;
            align-items: center;
            margin-top: 30px;
            flex-wrap: wrap;
        }

        .pagination-item {
            margin: 0 5px;
            padding: 8px 15px;
            border-radius: 4px;
            background-color: var(--background-color);
            color: var(--text-color);
            text-decoration: none;
            border: 1px solid var(--border-color);
            transition: all 0.2s ease;
        }

        .pagination-item:hover {
            background-color: var(--hover-color);
            border-color: var(--primary-color);
        }

        .pagination-item.active {
            background-color: var(--primary-color);
            color: white;
            border-color: var(--primary-color);
        }

        .pagination-item.disabled {
            opacity: 0.5;
            cursor: not-allowed;
            pointer-events: none;
        }

        .pagination-info {
            margin: 0 15px;
            color: var(--text-color);
        }

        .pagination-form {
            display: flex;
            align-items: center;
            margin-left: 15px;
        }

        .pagination-input {
            width: 60px;
            padding: 6px 10px;
            border-radius: 4px;
            border: 1px solid var(--border-color);
            margin: 0 5px;
        }

        .pagination-button {
            padding: 6px 12px;
            border-radius: 4px;
            background-color: var(--primary-color);
            color: white;
            border: none;
            cursor: pointer;
        }

        .pagination-button:hover {
            background-color: var(--secondary-color);
        }

        @media (max-width: 768px) {
            .container {
                padding: 20px 10px;
            }

            h1 {
                font-size: 20px;
            }

            th, td {
                padding: 8px;
            }

            .date-cell {
                display: none;
            }
            
            .pagination {
                flex-direction: column;
                gap: 10px;
            }
            
            .pagination-form {
                margin-top: 10px;
                margin-left: 0;
            }
        }

        @media (max-width: 480px) {
            .size-cell {
                display: none;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Listing Directory</h1>
        </div>

        <div class="path-display">
            {{.Dir}}
        </div>

        {{if .ParentDir}}
        <a href="{{$.ListPath}}?dir={{.ParentDir}}&root={{.Root}}&sort={{.SortBy}}&order={{.Order}}" class="back-link">
            <span class="back-icon">⬅</span> Back to Parent
        </a>
        {{end}}

        <table>
            <thead>
                <tr>
                    <th style="width: 60%">
                        <a href="?dir={{.Dir}}&root={{.Root}}&sort=name&order={{.NextOrder}}&page={{.CurrentPage}}">
                            Name
                            {{if eq .SortBy "name"}}
                                <span class="sort-indicator">{{if eq .Order "desc"}}⬇️{{else}}⬆️{{end}}</span>
                            {{end}}
                        </a>
                    </th>
                    {{if $.EnableFileMeta}}
                        <th style="width: 15%">
                            <a href="?dir={{.Dir}}&root={{.Root}}&sort=size&order={{.NextOrder}}&page={{.CurrentPage}}">
                                Size
                                {{if eq .SortBy "size"}}
                                    <span class="sort-indicator">{{if eq .Order "desc"}}⬇️{{else}}⬆️{{end}}</span>
                                {{end}}
                            </a>
                        </th>
                        <th style="width: 25%">
                            <a href="?dir={{.Dir}}&root={{.Root}}&sort=time&order={{.NextOrder}}&page={{.CurrentPage}}">
                                Modified Time
                                {{if eq .SortBy "time"}}
                                    <span class="sort-indicator">{{if eq .Order "desc"}}⬇️{{else}}⬆️{{end}}</span>
                                {{end}}
                            </a>
                        </th>
                    {{end}}
                </tr>
            </thead>
            <tbody>
                {{if .Files}}
                    {{range .Files}}
                    <tr>
                        <td>
                            {{if .IsDir}}
                                <a href="{{$.ListPath}}?dir={{.Path}}&root={{$.Root}}&sort={{$.SortBy}}&order={{$.Order}}" class="file-link">
                                    <span class="file-icon folder-icon">📁</span> {{.Name}}
                                </a>
                            {{else if $.EnableDownload}}
                                <a href="{{$.DownloadPath}}?path={{.Path}}" class="file-link">
                                    <span class="file-icon file-icon-regular">📄</span> {{.Name}}
                                </a>
                            {{else}}
                                <div class="file-text">
                                    <span class="file-icon file-icon-regular">📄</span> {{.Name}}
                                </div>
                            {{end}}
                        </td>
                        {{if $.EnableFileMeta}}
                            <td class="size-cell">{{if not .IsDir}}{{.Size | FormatSize}}{{end}}</td>
                            <td class="date-cell">{{.ModTime.Format "2006-01-02 15:04:05"}}</td>
                        {{end}}
                    </tr>
                    {{end}}
                {{else}}
                    <tr>
                        <td colspan="{{if $.EnableFileMeta}}3{{else}}1{{end}}" class="empty-message">
                            This directory is empty.
                        </td>
                    </tr>
                {{end}}
            </tbody>
        </table>
        
        {{if gt .TotalPages 1}}
        <div class="pagination">
            {{if .HasPrevPage}}
            <a href="?dir={{.Dir}}&root={{.Root}}&sort={{.SortBy}}&order={{.Order}}&page={{.PrevPage}}" class="pagination-item">
                Previous
            </a>
            {{else}}
            <span class="pagination-item disabled">Previous</span>
            {{end}}
            
            <span class="pagination-info">
                Page {{.CurrentPage}} / {{.TotalPages}}
            </span>
            
            {{if .HasNextPage}}
            <a href="?dir={{.Dir}}&root={{.Root}}&sort={{.SortBy}}&order={{.Order}}&page={{.NextPage}}" class="pagination-item">
                Next
            </a>
            {{else}}
            <span class="pagination-item disabled">Next</span>
            {{end}}
            
            <form class="pagination-form" action="{{$.ListPath}}" method="get">
                <input type="hidden" name="dir" value="{{.Dir}}">
                <input type="hidden" name="root" value="{{.Root}}">
                <input type="hidden" name="sort" value="{{.SortBy}}">
                <input type="hidden" name="order" value="{{.Order}}">
                <label>Go to: </label>
                <input type="number" name="page" min="1" max="{{.TotalPages}}" value="{{.CurrentPage}}" class="pagination-input">
                <button type="submit" class="pagination-button">Go</button>
            </form>
        </div>
        {{end}}
    </div>
</body>
</html>
`
