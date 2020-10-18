package main

import (
	"fmt"
	"time"
)	


func main(){

	fmt.Println("Welcome to AGILITY DIGITAL BALL CLOCK!")
	fmt.Println("Enter a number between 27 - 127. (Enter -1 to Exit)")
	fmt.Print("-> ")
	isStillAlive := true
	for isStillAlive{
		var ballCount int
		_, _= fmt.Scan(&ballCount)
  		//Checks if user wants to exit the program
		if ballCount == -1{
			break;
		}
		//Verifies the user input is correct
		if ballCount >= 27 && ballCount <= 127{
			//initialize slice of balls
			var balls []int
			
			for i := 0; i < ballCount; i++{
				balls = append(balls, i)
			}
			//Initialize minutes track, 5-minute track, and hour track
			minuteTrack := make([]int,5)
			minuteTrackCount := 0
			fiveMinuteTrack := make([]int,12)
			fiveMinuteTrackCount := 0
			hourTrack := make([]int, 12)
			hourTrackCount := 0

			//Starts stopwatch
			start := time.Now()

			minutesPassed := 0

			for true{
				//minute passed
				minutesPassed++
				//Get the first ball in queue
				ball := balls[0]
				//removes the ball from queue
				balls = balls[1:]
				//Put ball in the minute track
				minuteTrack[minuteTrackCount] = ball
				minuteTrackCount++
				//Checks if minute track is full
				if minuteTrackCount == 5{
					//Put 4 balls in the minute track to the queue in the revers order.
					for i := minuteTrackCount -2; i >= 0; i--{
						balls = append(balls, minuteTrack[i])
					}
					minuteTrackCount = 0
					//Put the last ball to the 5-minute track
					fiveMinuteTrack[fiveMinuteTrackCount] = ball
					fiveMinuteTrackCount++
					//Check if 5-minute track is full
					if fiveMinuteTrackCount == 12{
						//Put 11 balls in the 5-minute track to the queue in the reverse oder
						for i := fiveMinuteTrackCount-2; i >= 0; i--{
							balls = append(balls, fiveMinuteTrack[i])
						}
						fiveMinuteTrackCount =0
						//Put the last ball to the hour track
						hourTrack[hourTrackCount] = ball
						hourTrackCount++
						//Check if hour track is full
						if hourTrackCount == 12{
							//Pull 11 balls in the hours track to the queue in the reverse order
							for i := hourTrackCount -2; i >=0; i--{
								balls = append(balls, hourTrack[i])
							}
							hourTrackCount = 0
							//Put the last ball in the queue
							balls = append(balls, ball)
						}

					}
				}
				//Checks if the balls are in the original order
				if minuteTrackCount == 0 && fiveMinuteTrackCount ==0 &&hourTrackCount ==0{
					ordered := true

					if minutesPassed == 24*60*15{
						ordered = true
					}
					index := 0
					for _, item := range balls{
						if item != index{
							ordered = false
							break
						}
						index++
					}
					//Stop if all the balls are in the original order
					if ordered{
						break
					}
				}
			}
			elapsed := time.Since(start)
			totalDays := minutesPassed/60/24
			fmt.Println(ballCount, "balls take", totalDays, "days, computed in", elapsed)
			fmt.Println("Enter a number between 27 - 127:")
			fmt.Print("-> ")
		}else{
			fmt.Println("Incorrect input! Please enter a number between 27 - 127." )
			fmt.Print("-> ")
		}
	}
}

