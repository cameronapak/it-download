package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/heroku/x/hmetrics/onload"
)

var tpl *template.Template

// var g Get

// g is the struct that stores the json from Images.json - this file is located in dropbox where you can update
// the contents of to download different files.
// type Get struct {
// 	Images struct {
// 		Lighting struct {
// 			Name []string `json:"Name"`
// 			URL  []string `json:"URL"`
// 		} `json:"Lighting"`
// 		Propresenter struct {
// 			Name []string `json:"Name"`
// 			URL  []string `json:"URL"`
// 		} `json:"Propresenter"`
// 		Protools struct {
// 			Name []string `json:"Name"`
// 			URL  []string `json:"URL"`
// 		} `json:"Protools"`
// 		Usermac struct {
// 			Name []string `json:"Name"`
// 			URL  []string `json:"URL"`
// 		} `json:"Usermac"`
// 		Checkin struct {
// 			Name []string `json:"Name"`
// 			URL  []string `json:"URL"`
// 		} `json:"Checkin"`
// 		Smaart struct {
// 			Name []string `json:"Name"`
// 			URL  []string `json:"URL"`
// 		} `json:"Smaart"`
// 	} `json:"Images"`
// }

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

func main() {
	server := http.Server{
		Addr:        ":" + "8080",
		IdleTimeout: time.Hour,
	}

	http.HandleFunc("/", Index)
	// http.HandleFunc("/image.json", DownloadFile)

	server.ListenAndServe()
}

func Index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.tmpl.html", nil)
	HandleError(w, err)
}

// func DownloadFile(w http.ResponseWriter, r *http.Request) {
// 	url := "https://www.dropbox.com/s/bngfnoxe5kwv8lc/images.json?dl=1"

// 	timeout := time.Duration(5) * time.Second
// 	transport := &http.Transport{
// 		ResponseHeaderTimeout: timeout,
// 		Dial: func(network, addr string) (net.Conn, error) {
// 			return net.DialTimeout(network, addr, timeout)
// 		},
// 		DisableKeepAlives: true,
// 	}
// 	client := &http.Client{
// 		Transport: transport,
// 	}
// 	resp, err := client.Get(url)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer resp.Body.Close()

// 	//copy the relevant headers. If you want to preserve the downloaded file name, extract it with go's url parser.
// 	w.Header().Set("Content-Disposition", "attachment: filename=images.json")
// 	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
// 	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))

// 	//stream the body to the client without fully loading it into memory
// 	io.Copy(w, resp.Body)
// }

// func downloadScript(y1 []string, y2 []string) {
// 	start := time.Now()
// 	for x := range y1 {
// 		user, err := user.Current()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		os.Chdir(user.HomeDir + "/Downloads")

// 		dst, err := os.Create(y1[x])
// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 		resp, _ := http.Get(y2[x])

// 		src, err := io.Copy(dst, resp.Body)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(y1[x], src, "Bytes", "Downloaded Successful!.")
// 	}
// 	elapsed := time.Since(start)
// 	log.Printf("Downloaded all files in %s", elapsed)
// }

// func JsonImages() {
// 	url := "https://www.dropbox.com/s/bngfnoxe5kwv8lc/images.json?dl=1"

// 	dst, err := os.Create("images.json")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	res, err := http.Get(url)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	src, err := io.Copy(dst, res.Body)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Println(src, "file created!")

// 	body, err := ioutil.ReadFile("./images.json")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	json.Unmarshal(body, &g)
// }

// func Install(images string) {
// 	switch images {
// 	case "Lighting":
// 		fmt.Println("Downloading Lighting Files...")
// 		downloadScript(g.Images.Lighting.Name, g.Images.Lighting.URL)
// 		openDropbox := exec.Command("hdiutil", "attach", "Dropbox.dmg")
// 		openDropbox.Run()
// 		installDropbox := exec.Command("hdiutil", "attach", "/Volumes/Dropbox Installer/Dropbox.app")
// 		installDropbox.Run()

// 		openVista := exec.Command("hdiutil", "attach", "Vista.dmg")
// 		openVista.Run()

// 		openKaseya := exec.Command("unzip", "-a", "Kaseya.zip")
// 		openKaseya.Run()
// 	case "Checkin":
// 		fmt.Println("Downloading Check-In Files...")
// 		downloadScript(g.Images.Checkin.Name, g.Images.Checkin.URL)
// 		openDropbox := exec.Command("Dropbox.exe")
// 		openDropbox.Run()

// 		openF1 := exec.Command("msiexec", "/a", "F1.msi")
// 		openF1.Run()

// 		openKaseya := exec.Command("Kaseya.exe")
// 		openKaseya.Run()
// 	case "Protools":
// 		fmt.Println("Downloading Protools Files...")
// 		downloadScript(g.Images.Protools.Name, g.Images.Protools.URL)
// 		openProtools := exec.Command("open", "Protools.dmg")
// 		openProtools.Run()

// 		openDVS := exec.Command("open", "DVS.dmg")
// 		openDVS.Run()

// 		openMidi := exec.Command("open", "Midi.pkg")
// 		openMidi.Run()

// 		openDropbox := exec.Command("open", "Dropbox.dmg")
// 		openDropbox.Run()

// 		installDropbox := exec.Command("hdiutil", "attach", "/Volumes/Dropbox Installer/Dropbox.app")
// 		installDropbox.Run()

// 		openPitch := exec.Command("open", "Pitchle3.pkg")
// 		openPitch.Run()

// 		openDanteController := exec.Command("open", "DanteController.dmg")
// 		openDanteController.Run()

// 		openSpotify := exec.Command("unzip", "-a", "SpotifyInstaller.zip")
// 		openSpotify.Run()

// 		openKaseya := exec.Command("unzip", "-a", "Kaseya.zip")
// 		openKaseya.Run()
// 	case "Propresenter":
// 		fmt.Println("Downloading Propresenter Files...")
// 		downloadScript(g.Images.Propresenter.Name, g.Images.Propresenter.URL)
// 		openPropresenter := exec.Command("open", "Propresenter.dmg")
// 		openPropresenter.Run()

// 		openVideohub := exec.Command("open", "Videohub.dmg")
// 		openVideohub.Run()

// 		openVlc := exec.Command("open", "Vlc.dmg")
// 		openVlc.Run()

// 		openDropbox := exec.Command("open", "Dropbox.dmg")
// 		openDropbox.Run()

// 		installDropbox := exec.Command("hdiutil", "attach", "/Volumes/Dropbox Installer/Dropbox.app")
// 		installDropbox.Run()

// 		openHipchat := exec.Command("open", "Hipchat.dmg")
// 		openHipchat.Run()

// 		openKaseya := exec.Command("unzip", "-a", "Kaseya.zip")
// 		openKaseya.Run()
// 	case "Usermac":
// 		fmt.Println("Downloading User Mac Files...")
// 		downloadScript(g.Images.Usermac.Name, g.Images.Usermac.URL)
// 		openOffice := exec.Command("open", "Office.pkg")
// 		openOffice.Run()

// 		deleteDock := exec.Command("defaults", "delete", "com.apple.dock", "persistent-apps")
// 		deleteDock.Run()

// 		deleteDock2 := exec.Command("killall", "Dock")
// 		deleteDock2.Run()

// 		openDropbox := exec.Command("open", "Dropbox.dmg")
// 		openDropbox.Run()

// 		installDropbox := exec.Command("hdiutil", "attach", "/Volumes/Dropbox Installer/Dropbox.app")
// 		installDropbox.Run()

// 		openBria := exec.Command("open", "Bria.dmg")
// 		openBria.Run()

// 		openFirefox := exec.Command("open", "Firefox.dmg")
// 		openFirefox.Run()

// 		openChrome := exec.Command("open", "Chrome.dmg")
// 		openChrome.Run()

// 		openVidyo := exec.Command("open", "Vidyo.dmg")
// 		openVidyo.Run()

// 		openCisco := exec.Command("open", "Cisco.dmg")
// 		openCisco.Run()

// 		openKaseya := exec.Command("unzip", "-a", "Kaseya.zip")
// 		openKaseya.Run()
// 	case "Smaart":
// 		fmt.Println("Downloading Smaart Files...")
// 		downloadScript(g.Images.Smaart.Name, g.Images.Smaart.URL)
// 		openDropbox := exec.Command("open", "Dropbox.dmg")
// 		openDropbox.Run()

// 		installDropbox := exec.Command("hdiutil", "attach", "/Volumes/Dropbox Installer/Dropbox.app")
// 		installDropbox.Run()
// 		openKaseya := exec.Command("unzip", "-a", "Kaseya.zip")
// 		openKaseya.Run()
// 	}
// }
