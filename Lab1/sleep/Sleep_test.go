//package main


package mySleep

import ("testing";"time")

func TestCustomSleep(t *testing.T) {
	currentTime:= time.Now()
    CustomSleep(time.Duration(time.Minute *1))
    sleepTime := time.Now()
    if currentTime == sleepTime {
      t.Error("failed !")
    }
}

