package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/layou233/ZBProxy/common"
	"github.com/layou233/ZBProxy/config"
	"github.com/layou233/ZBProxy/console"
	"github.com/layou233/ZBProxy/service"
	"github.com/layou233/ZBProxy/version"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
)

func main() {
	log.SetOutput(color.Output)
	//rand.Seed(time.Now().UnixNano())
	console.SetTitle(fmt.Sprintf("ZBProxy %v | Running...", version.Version))
	console.Println(color.HiRedString(` ______  _____   _____   _____    _____  __    __ __    __
|___  / |  _  \ |  _  \ |  _  \  /  _  \ \ \  / / \ \  / /
   / /  | |_| | | |_| | | |_| |  | | | |  \ \/ /   \ \/ /`), color.HiWhiteString(`
  / /   |  _  { |  ___/ |  _  /  | | | |   }  {     \  /
 / /__  | |_| | | |     | | \ \  | |_| |  / /\ \    / /
/_____| |_____/ |_|     |_|  \_\ \_____/ /_/  \_\  /_/`))
	color.HiGreen("Welcome to ZBProxy %s (%s)!\n", version.Version, version.CommitHash)
	color.HiBlack("Build Information: %s, %s/%s, CGO %s\n",
		runtime.Version(), runtime.GOOS, runtime.GOARCH, common.CGOHint)
	// go version.CheckUpdate()

	config.LoadConfig()
	service.Listeners = make([]net.Listener, 0, len(config.Config.Services))

	for _, s := range config.Config.Services {
		go service.StartNewService(s)
	}

	// hot reload
	// use inotify on Linux
	// use Win32 ReadDirectoryChangesW on Windows
	{
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Panic(err)
		}
		defer watcher.Close()
		err = config.MonitorConfig(watcher)
		if err != nil {
			log.Panic("Config Reload Error : ", err)
		}
	}

	{
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
		<-osSignals
		// stop the program
		// sometimes after the program exits on Windows, the ports are still occupied and "listening".
		// so manually closes these listeners when the program exits.
		for _, listener := range service.Listeners {
			if listener != nil { // avoid null pointers
				listener.Close()
			}
		}
	}
}
