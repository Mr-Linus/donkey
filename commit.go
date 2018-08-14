package main

import (
	log "github.com/Sirupsen/logrus"
	"fmt"
	"os/exec"
)

func commitContainer(imageName string){
	mntURL := "images/mnt"
	imageTar := "images/" + imageName + ".tar"
	fmt.Printf("%s",imageTar)
	if _, err := exec.Command("tar", "-czf", imageTar, "-C", mntURL, ".").CombinedOutput(); err != nil {
		log.Errorf("Tar folder %s error %v", mntURL, err)
	}
}