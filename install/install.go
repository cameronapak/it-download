package install

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"time"
)

var g Get

// g is the struct that stores the json from Images.json - this file is located in dropbox where you can update
// the contents of to download different files.
type Get struct {
	Images struct {
		Lighting struct {
			Name []string `json:"Name"`
			URL  []string `json:"URL"`
		} `json:"Lighting"`
		Propresenter struct {
			Name []string `json:"Name"`
			URL  []string `json:"URL"`
		} `json:"Propresenter"`
		Protools struct {
			Name []string `json:"Name"`
			URL  []string `json:"URL"`
		} `json:"Protools"`
		Usermac struct {
			Name []string `json:"Name"`
			URL  []string `json:"URL"`
		} `json:"Usermac"`
		Checkin struct {
			Name []string `json:"Name"`
			URL  []string `json:"URL"`
		} `json:"Checkin"`
		Smaart struct {
			Name []string `json:"Name"`
			URL  []string `json:"URL"`
		} `json:"Smaart"`
	} `json:"Images"`
}

func downloadScript(y1 []string, y2 []string) {
	start := time.Now()
	for x := range y1 {
		user, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		os.Chdir(user.HomeDir + "/Downloads")

		dst, err := os.Create(y1[x])
		if err != nil {
			fmt.Println(err)
		}

		resp, _ := http.Get(y2[x])

		src, err := io.Copy(dst, resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(y1[x], src, "Bytes", "Downloaded Successful!.")
	}
	elapsed := time.Since(start)
	log.Printf("Downloaded all files in %s", elapsed)
}

func JsonImages() {
	url := "https://www.dropbox.com/s/bngfnoxe5kwv8lc/images.json?dl=1"

	dst, err := os.Create("images.json")
	if err != nil {
		fmt.Println(err)
	}

	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	src, err := io.Copy(dst, res.Body)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(src, "file created!")

	body, err := ioutil.ReadFile("./images.json")
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal(body, &g)
}

func Install(images string) {
	switch images {
	case "Lighting":
		fmt.Println("Downloading Lighting Files...")
		JsonImages()
		downloadScript(g.Images.Lighting.Name, g.Images.Lighting.URL)
		openDropbox := exec.Command("hdiutil", "attach", "Dropbox.dmg")
		openDropbox.Run()

		installDropbox := exec.Command("hdiutil", "attach", "/Volumes/Dropbox Installer/Dropbox.app")
		installDropbox.Run()

		openVista := exec.Command("hdiutil", "attach", "Vista.dmg")
		openVista.Run()

		openKaseya := exec.Command("unzip", "-a", "Kaseya.zip")
		openKaseya.Run()
	case "Checkin":
		fmt.Println("Downloading Check-In Files...")
		JsonImages()
		downloadScript(g.Images.Checkin.Name, g.Images.Checkin.URL)
		openDropbox := exec.Command("Dropbox.exe")
		openDropbox.Run()

		openF1 := exec.Command("msiexec", "/a", "F1.msi")
		openF1.Run()

		openKaseya := exec.Command("Kaseya.exe")
		openKaseya.Run()
	case "Protools":
		fmt.Println("Downloading Protools Files...")
		JsonImages()
		downloadScript(g.Images.Protools.Name, g.Images.Protools.URL)
		openProtools := exec.Command("open", "Protools.dmg")
		openProtools.Run()

		openDVS := exec.Command("open", "DVS.dmg")
		openDVS.Run()

		openMidi := exec.Command("open", "Midi.pkg")
		openMidi.Run()

		openDropbox := exec.Command("open", "Dropbox.dmg")
		openDropbox.Run()

		installDropbox := exec.Command("hdiutil", "attach", "/Volumes/Dropbox Installer/Dropbox.app")
		installDropbox.Run()

		openPitch := exec.Command("open", "Pitchle3.pkg")
		openPitch.Run()

		openDanteController := exec.Command("open", "DanteController.dmg")
		openDanteController.Run()

		openSpotify := exec.Command("unzip", "-a", "SpotifyInstaller.zip")
		openSpotify.Run()

		openKaseya := exec.Command("unzip", "-a", "Kaseya.zip")
		openKaseya.Run()
	case "Propresenter":
		fmt.Println("Downloading Propresenter Files...")
		JsonImages()
		downloadScript(g.Images.Propresenter.Name, g.Images.Propresenter.URL)
		openPropresenter := exec.Command("open", "Propresenter.dmg")
		openPropresenter.Run()

		openVideohub := exec.Command("open", "Videohub.dmg")
		openVideohub.Run()

		openVlc := exec.Command("open", "Vlc.dmg")
		openVlc.Run()

		openDropbox := exec.Command("open", "Dropbox.dmg")
		openDropbox.Run()

		installDropbox := exec.Command("hdiutil", "attach", "/Volumes/Dropbox Installer/Dropbox.app")
		installDropbox.Run()

		openHipchat := exec.Command("open", "Hipchat.dmg")
		openHipchat.Run()

		openKaseya := exec.Command("unzip", "-a", "Kaseya.zip")
		openKaseya.Run()
	case "Usermac":
		fmt.Println("Downloading User Mac Files...")
		JsonImages()
		downloadScript(g.Images.Usermac.Name, g.Images.Usermac.URL)
		openOffice := exec.Command("open", "Office.pkg")
		openOffice.Run()

		deleteDock := exec.Command("defaults", "delete", "com.apple.dock", "persistent-apps")
		deleteDock.Run()

		deleteDock2 := exec.Command("killall", "Dock")
		deleteDock2.Run()

		openDropbox := exec.Command("open", "Dropbox.dmg")
		openDropbox.Run()

		installDropbox := exec.Command("hdiutil", "attach", "/Volumes/Dropbox Installer/Dropbox.app")
		installDropbox.Run()

		openBria := exec.Command("open", "Bria.dmg")
		openBria.Run()

		openFirefox := exec.Command("open", "Firefox.dmg")
		openFirefox.Run()

		openChrome := exec.Command("open", "Chrome.dmg")
		openChrome.Run()

		openVidyo := exec.Command("open", "Vidyo.dmg")
		openVidyo.Run()

		openCisco := exec.Command("open", "Cisco.dmg")
		openCisco.Run()

		openKaseya := exec.Command("unzip", "-a", "Kaseya.zip")
		openKaseya.Run()
	case "Smaart":
		fmt.Println("Downloading Smaart Files...")
		JsonImages()
		downloadScript(g.Images.Smaart.Name, g.Images.Smaart.URL)
		openDropbox := exec.Command("open", "Dropbox.dmg")
		openDropbox.Run()

		installDropbox := exec.Command("hdiutil", "attach", "/Volumes/Dropbox Installer/Dropbox.app")
		installDropbox.Run()
		openKaseya := exec.Command("unzip", "-a", "Kaseya.zip")
		openKaseya.Run()
	}
}
