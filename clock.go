package main
import (
	"fmt"
	"os"
)

func main(){

	var args []string = os.Args

	if len(args) < 2{
		fmt.Println("You need to provide the Hour:Minute as an argument!")
		return
	}

	var hour int
	var minute int

	_, err := fmt.Sscanf(args[1], "%d:%d", &hour, &minute)
	if err != nil{
		fmt.Println("Error parsing your argument. Make sure your argument is in the Hour:Minute format")
	}

	if hour < 0 || hour > 12{
		fmt.Println("Your hour hand must be between 0-12")
		return
	}
	if minute < 0 || minute > 60{
		fmt.Println("Your minute hand must be between 0-60")
		return
	}

	fmt.Println("Your angle between clock hands is", find_clock_hands_angle(hour,minute), "degrees")
}

func find_clock_hands_angle(hour int, minute int) int{

	const DEGREE_PER_HOUR int = 360 / 12
	const DEGREE_PER_MINUTE int = 360 / 60

	var hourAngle int = hour * DEGREE_PER_HOUR
	var minuteAngle int = minute * DEGREE_PER_MINUTE

	var result int = hourAngle - minuteAngle

	// We use absolute value here because direction does not matter
	// when we are finding the minimal angle as direction is just which
	// hand you evaluate first.
	if result < 0 {
		result *= -1
	}

	// We use min here to quickly determine if it's faster to move right directly
	// if if it's better to move left.
	return min(result, 360 - result)
}