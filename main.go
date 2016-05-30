package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/oliamb/cutter"
	"github.com/tebeka/selenium"
)

func main() {

	uriPtr := flag.String("url", "http://redash.com", "a redash url")
	xpath := flag.String("xpath", `//*[@id="1"]/div[1]/div[3]/page-header/div/div[1]/h3`, "an xpath to graph")
	flag.Parse()

	// Chrome driver without specific version
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, "")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer wd.Quit()

	username := os.Getenv("REDASH_USER")
	password := os.Getenv("REDASH_PASSWORD")
	url := *uriPtr

	// Get simple playground interface
	wd.Get(url)

	// Enter username in textarea
	elem, _ := wd.FindElement(selenium.ById, "inputEmail")
	elem.Clear()
	elem.SendKeys(username)

	// Enter password in textarea
	elempass, _ := wd.FindElement(selenium.ById, "inputPassword")
	elempass.Clear()
	elempass.SendKeys(password)

	// Click the run button
	btn, _ := wd.FindElement(selenium.ByXPATH, "/html/body/div/div[2]/div/div/form/button")
	btn.Click()
	fmt.Println("Logged in...")

	time.Sleep(1 * time.Second)
	// Get the result
	h3, err := wd.FindElement(selenium.ByXPATH, *xpath)
	if err != nil {
		fmt.Println(err, "dashboard_title")
		os.Exit(1)
	}

	output := ""
	// Wait for run to finish
	for {
		output, _ = h3.Text()
		fmt.Println("Checking if output eq Player Overall =", output)
		if output != "Adserver Admin Report" {
			time.Sleep(time.Millisecond * 100)
		} else {
			break
		}
	}

	fmt.Printf("Got: %s\n", output)
	time.Sleep(5 * time.Second)
	pic, err := wd.Screenshot()
	r := bytes.NewReader(pic)
	//err  = ioutil.WriteFile("/vagrant/image.jpg", pic, 0644)
	if err != nil {
		fmt.Println(err)
	}

	img, _, err := image.Decode(r)
	if err != nil {
		fmt.Println("Cannot decode image:", err)
	}
	fmt.Println("Screenshot taken about to crop...")

	//overallChart, _ := wd.FindElement(selenium.ByXPATH, `//*[@id="dashboard"]/div[1]`)
	overallChart, _ := wd.FindElement(selenium.ByXPATH, `//*[@id="1"]/div[1]/div[3]/div[1]/div/div/visualization-renderer/div/chart-renderer/div/plotly-chart/div/div/div`)
	size, _ := overallChart.Size()
	location, _ := overallChart.Location()
	cImg, err := cutter.Crop(img, cutter.Config{
		Height:  size.Height,                         // height in pixel or Y ratio(see Ratio Option below)
		Width:   size.Width,                          // width in pixel or X ratio
		Mode:    cutter.TopLeft,                      // Accepted Mode: TopLeft, Centered
		Anchor:  image.Point{location.X, location.Y}, // Position of the top left point
		Options: 0,                                   // Accepted Option: Ratio
	})
	fmt.Println("Image Cropped, running img2png...")
	image2png(cImg, "result.jpg")
}

func image2png(img image.Image, name string) {
	outfilename := name
	outfile, err := os.Create(outfilename)
	if err != nil {
		panic(err.Error())
	}
	defer outfile.Close()
	png.Encode(outfile, img)

}
