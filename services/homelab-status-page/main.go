package main

//https://www.youtube.com/watch?v=Rp3A5q9L_bg
//
// CSV - output on the level of reddit.com/r/place



import (

	"bytes"
	"io"
	///"log"
	"os"
	"path"
	"time"
	"context"

	//"github.com/gosimple/slug"
	//"github.com/yuin/goldmark"
)
import (
	"fmt"
	"strings"
	"html/template"


	"path/filepath"

	"github.com/labstack/echo/v4"
		"net/http"
	//"github.com/labstack/echo/v4/middleware"
	//
	//
	"github.com/kkdai/youtube/v2"
	"github.com/a-h/templ"
	"os/exec"
)
type Post struct {
	Date    time.Time
	Title   string
	Content string
}


func Unsafe(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, html)
		return
	})
}

func renderTemplate(templateName string) echo.HandlerFunc {
	// sweaters := Inventory{"wool",
	// 	"17"}
	return func(c echo.Context) error {

		if templateName == "/index" {
			templateName = "index.temp"
		}
		rootPath := os.ExpandEnv("$HOME/hashirama/services/homelab-status-page/views")
		fmt.Printf("Received route on /%s\n", templateName)
		name := path.Join(rootPath, templateName + ".html")
		fmt.Printf("ROUTING TO %s\n", name)
		b, err := os.ReadFile(name) // just pass the file name
		if err != nil {
			fmt.Printf("oh shit dog %s\n", err)
			return nil
		}
		str := string(b) // convert content to a 'string'
		//content := Unsafe(str)
		c.Response().Header().Set("Content-Type", "text/html")
		//headerComponent("fuck").Render(context.Background(), os.Stdout)
		//

		//func ToGoHTML(ctx context.Context, c Component) (s template.HTML, err error)
		return headerComponent(str).Render(context.Background(), c.Response().Writer)
		var buf bytes.Buffer

		err = headerComponent(str).Render(context.Background(), &buf)
		_, err = buf.WriteTo(c.Response().Writer)


		return err
		//return
	}
}


func trimMyDick(filePath string) string {
	    baseName := filepath.Base(filePath)

    // Remove the file extension
    fileName := strings.TrimSuffix(baseName, filepath.Ext(baseName))
	//fmt.Printf("route\n", filepath.Base(fileName), "\n")
	return fileName
}



type Template struct {
	templates *template.Template
}
func custom_endpont (c echo.Context) error {
		base := filepath.Join(os.ExpandEnv("$HOME/hashirama/services/homelab-status-page/"), "readme.org")
		dat, err := os.ReadFile(base)
		if err != nil {
			fmt.Print("base \n\n\n", err)
			return err
		}
		return c.String(http.StatusOK, string(dat))
	}

//https://github.com/danawoodman/go-echo-htmx-templ/blob/main/handlers/render.go
// func ExampleClient(videoID string) (*youtube.Video, io.ReadCloser, *youtube.Format, error) {
//     client := youtube.Client{}

//     video, err := client.GetVideo(videoID)
//     if err != nil {
//         return nil, nil, nil, err
//     }

//     formats := video.Formats.WithAudioChannels() // only get videos with audio
//     stream, _, err := client.GetStream(video, &formats[0])
//     if err != nil {
//         return nil, nil, nil, err
//     }

//     return video, stream, &formats[0], nil
// }
//
func ExampleClient (videoID string) string {
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	defer stream.Close()
	vid_loc := fmt.Sprintf("static/video/%s.mp4", videoID)
	file, err := os.Create(vid_loc)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
	return vid_loc
}

type YM struct {
    URL string `form:"url"`
}

func handleConvertVideoToPDF(c echo.Context) error {
	fmt.Println("error", "my niggar")
    formData := new(YM)
    if err := c.Bind(formData); err != nil {
		fmt.Println("error", err)
        return err
    }


	fmt.Println("error", "shti")

    youtubeURL := formData.URL
    //video, stream, format, err := ExampleClient(youtubeURL)
    videoLoc := ExampleClient(youtubeURL)
fmt.Println("finish DL", youtubeURL)
    // Set the headers to serve the video directly
    // c.Response().Header().Set("Content-Type", format.MimeType)
    // c.Response().Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=\"%s.mp4\"", video.Title))

    // // Copy the video stream to the response writer
    // _, err = io.Copy(c.Response().Writer, stream)

        return c.JSON(http.StatusOK, map[string]string{
            "videoUrl": videoLoc,

        })

	//	//ai whisper
	//	 turn transcript into summary ?
	//
	//
	//	 on youtube page -> right click = summary
    return nil
}

//err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
func fuck() {
	cmd := exec.Command("bash", "~/go/bin/templ generate")


	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	err := cmd.Start()
	if err != nil { return }
	fmt.Printf("Output: %s\n", outBuf.String())
    fmt.Printf("Error: %s\n", errBuf.String())
	//exec a binary in go ~/go/bin/templ generate from golang
}


//external validation vs respect - no one can trully know you - but also maybe, you cant really know yourself except through others?
func main() {
	fuck()
	e := echo.New()
	e.POST("/video-to-pdf", handleConvertVideoToPDF)
	expandMyPuss := os.ExpandEnv("$HOME/hashirama/services/homelab-status-page/views/*.html")
	allMyRoutes, err := filepath.Glob(expandMyPuss)
	if err != nil {return }
	for i := 0; i < len(allMyRoutes); i++ {
		filePath := string(allMyRoutes[i])
		if filePath == "/" {continue}
		trimmed := trimMyDick(filePath)
		e.GET(trimmed, renderTemplate(trimmed))
    }
	e.GET("/form", renderTemplate("form"))
	e.GET("/", renderTemplate("index"))
	e.GET("/custom_endpoint", custom_endpont)


	e.Static("/static", "static")
	e.Static("/assets", "static")
	e.Logger.Fatal(e.Start("0.0.0.0:1323"))
}
		//https://github.com/explore
// https://github.com/2b-t/docker-realtime
// https://arc.net/l/quote/afqsckfz
// docker-ros-desktop-vnc
// Packer for building VM OS (Windows & Linux) templates using Proxmox builder
// https://github.com/netdata/netdata
// https://dashy.to/
// https://github.com/badele/nix-homelab
// https://github.com/codespaces
// https://github.blog/2024-06-03-arm64-on-github-actions-powering-faster-more-efficient-build-systems/

		//https://pkg.go.dev/text/template
	   // 	filePath := filepath.Join(os.ExpandEnv("$HOME/hashirama/services/homelab-status-page/views"), templateName+".html")

       // data := map[string]interface{}{
       //      "Title":   "Hello, World!",
       //      "Content": "This is a templ template rendered on the fly.",
       //  }
//https://echo.labstack.com/docs/category/middleware
// google drive
// phone
// find my friends - realtime dispatch
// validate projects quick
