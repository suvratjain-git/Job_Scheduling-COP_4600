/*
I Suvrat Jain (su956999) affirm that this program is entirely my own work and that 
I have neither developed my code together with any another person, 
nor copied any code from any other person, nor permitted my code to be copied or otherwise used by any other person, 
nor have I copied, modified, or otherwise used programs created by others. 
I acknowledge that any violation of the above terms will be treated as academic dishonesty.
*/
package main

import (
     "fmt"
     "os"
     "bufio"
     "strconv"
     "sort"
)

//This struct holds the Process Information such as Name, Arrival Time, and Burst Time
type ProcessInfo struct {
   	name string
   	arrivalTime int
   	burstTime int
}

//list of variables
var (
	processCount, totalTime, quantum int
	schedulingAlgorithm string

	process []ProcessInfo	
	waitTime []int
	turnAroundTime []int
	completionTime []int
)

//returns the processCount, totalTime, schedulingAlgorithm, and quantum from the input file
func getProcessInfo (filename string) (pc int, tT int, sa string, q int) {

	//create a file pointer
	file, _ := os.Open(filename)

	//make a scan object to iterate through the file using the file pointer
	fileScanner := bufio.NewScanner(file)

	//split the file into words
	fileScanner.Split(bufio.ScanWords)

	
	for fileScanner.Scan() {

		//get the first word as a string and set it to word variable 
		word := fileScanner.Text()

		//if the string encountered is "processcount" then increment the pointer to point to next word,
		// which is the process count value and set it equal to pc variable
		// else if the string encountered is "runfor" then increment the pointer to point to next word,
		// which is the total time to run the algorithm for and set it equal to tT variable
		// else if the string encountered is "use" then increment the pointer to point to next word,
		// which tells what algorithm to use and set it equal to sa variable
		// else if the string encountered is "quantum" then increment the pointer to point to next word,
		// which is the quantum number and set it equal to q variable
		if(word == "processcount") {

			//increment scanner
			fileScanner.Scan()
			//convert the string to integer
		 	pc,_ = strconv.Atoi(fileScanner.Text())

		} else if(word == "runfor") {

        	fileScanner.Scan()
        	tT,_ = strconv.Atoi(fileScanner.Text())

        } else if(word == "use") {

        	fileScanner.Scan()
        	sa = fileScanner.Text()

        	//break the scanning if the algorithm is not round robin since because then it will not have a quantum value
        	if (sa != "rr") {
        		break
        	}

        } else if(word == "quantum"){

        	fileScanner.Scan()

        	//check if the scanned string is an integer to prevent quantum = -
        	temp, err := strconv.Atoi(fileScanner.Text()); 

        	if (err == nil) {
    			q = temp
			}

			//break the scanning after the quantum number is found
			break
        } 

	}

	return pc, tT, sa, q 
}

func getListOfProcesses (filename string, pc int) (p []ProcessInfo) {

	//create a file pointer
	file, _ := os.Open(filename)

	//make a scan object to iterate through the file using the file pointer
	fileScanner := bufio.NewScanner(file)

	//split the file into words
	fileScanner.Split(bufio.ScanWords)

	//make a slice of processes
	p = make([]ProcessInfo, pc)

	//struct variable
	var process ProcessInfo
	var count int = 0

	for fileScanner.Scan() {

		word := fileScanner.Text()

		if(word == "end") {

			break

		} else if(word == "name") {

        	fileScanner.Scan()
        	process.name = fileScanner.Text()
        	count++

        } else if(word == "arrival") {

        	fileScanner.Scan()
        	process.arrivalTime,_ = strconv.Atoi(fileScanner.Text())
        	count++
        	
        } else if(word == "burst") {

        	fileScanner.Scan()
        	process.burstTime,_ = strconv.Atoi(fileScanner.Text())
        	count++

        }

        //when all three entries have been added to the struct then put it in the slice
        if(count == 3) {
        	p = append(p, process)
        	count = 0
        }
	}

	//truncate the slice
	p = p[pc:]
	return p
}



func fcfs (process []ProcessInfo, processCount int)  {
	
	//Completion Time
	var completionT int = 0
		
	//sort the slice with respect to arrival time
	sort.Slice(process, func (i, j int) bool { return process[i].arrivalTime < process[j].arrivalTime })
	
	//allocate memory to completionTime, waitTime, and turnAroundTime slices
	completionTime = make([]int, processCount)
	turnAroundTime = make([] int, processCount)
	waitTime = make([]int, processCount)

	for i := 0; i < processCount; i++ {

		completionT += process[i].burstTime
		completionTime[i] = completionT
		
		turnAroundTime[i] =  completionTime[i] - process[i].arrivalTime

		waitTime[i] = turnAroundTime[i] - process[i].burstTime
	}

	fmt.Printf("Processes = %v\n", process)
	fmt.Printf("Completion Time = %v\n", completionTime)
	fmt.Printf("Turn Around Time = %v\n", turnAroundTime)
	fmt.Printf("Wait Time = %v\n", waitTime)
}

func sjf (process []ProcessInfo, processCount int)  {
	
	//sort the slice with respect to first arrival and then burst time
	sort.Slice(process, func(i, j int) bool { return process[i].burstTime < process[j].burstTime })
}

func rr (process []ProcessInfo, processCount int, q int)  {
	fmt.Println("RR = ")
	fmt.Println(process)
}

func main() {

	//Read the file name from the CLI arguements and convert it into an array of bytes
	inputFile := os.Args[1];

	processCount, totalTime, schedulingAlgorithm, quantum = getProcessInfo(inputFile)
	process = getListOfProcesses(inputFile, processCount)

	if(schedulingAlgorithm == "fcfs") {
		fcfs(process, processCount)
	} else if (schedulingAlgorithm == "sjf") {
		sjf(process, processCount)
	} else if (schedulingAlgorithm == "rr") {
		rr(process, processCount, quantum)
	}


    // fmt.Printf("\nProcess Count = %d\n", processCount)
    // fmt.Printf("Time to run the algorithm for = %d\n", totalTime)
    // fmt.Printf("Type of Scheduling Algorithm = %s\n", schedulingAlgorithm)
    // fmt.Printf("Quantum (if any) = %d\n\n", quantum)
    // fmt.Println("List of Processes: ")
    // fmt.Println(process)
    // fmt.Println()
   
 }













