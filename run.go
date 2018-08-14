package main

import (
	"./container"
	"./cgroups/subsystems"
	"./cgroups/"
	log "github.com/Sirupsen/logrus"
	"os"
	"strings"
)


func Run(tty bool, comArray []string, res *subsystems.ResourceConfig, volume string) {
	parent, writePipe := container.NewParentProcess(tty, volume)
	if parent == nil {
		log.Errorf("New parent process error")
		return
	}
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	//cgroup-mananger 
	cgroupManager := cgroups.NewCgroupManager("donkey-cgroup")
	defer cgroupManager.Destroy()
	cgroupManager.Set(res)
	cgroupManager.Apply(parent.Process.Pid)

	
	sendInitCommand(comArray, writePipe)
	if tty {
		parent.Wait()
	}
	mntURL := "./images/mnt/"
	rootURL := "./images/"
	container.DeleteWorkSpace(rootURL, mntURL, volume)
	os.Exit(0)
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	log.Infof("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}