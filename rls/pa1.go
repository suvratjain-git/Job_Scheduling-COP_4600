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
	ID int
   	name string
   	arrivalTime int
   	burstTime int
   	waitTime int
   	turnAroundTime int
   	completionTime int
   	selected bool
   	completed bool
   	selectionTime int
}

//list of variables
var (
	processCount, totalTime, quantum int
	schedulingAlgorithm string
	completionT int

	
	process []ProcessInfo	
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
        	process.ID++
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



func fcfs (process []ProcessInfo, processCount int, usefor int)  {
		
	//sort the slice with respect to arrival time
	sort.Slice(process, func (i, j int) bool { return process[i].arrivalTime < process[j].arrivalTime })

	fmt.Printf("%3d processes\n", processCount)
	fmt.Println("Using First-Come First-Served")

	var arrivalQueue []ProcessInfo
 	arrivalQueue = make([]ProcessInfo, processCount)
 	arrivalQueueCapacity := 0

	for i:=0; i<processCount; i++ {

		process[i].completionTime = completionT + process[i].burstTime
		completionT = process[i].completionTime

		process[i].turnAroundTime = process[i].completionTime - process[i].arrivalTime

		process[i].waitTime = process[i].turnAroundTime - process[i].burstTime
		if process[i].waitTime < 0 {
			process[i].waitTime = 0
		}
	}
		
	time := 0
	index := 0
	for time < usefor {

		for i:=0; i < processCount; i++ {
			if(process[i].arrivalTime == time) {

				fmt.Printf("Time %3d : %s arrived\n", time, process[i].name)
				arrivalQueue[i] = process[i]
				arrivalQueueCapacity++
			}
		}

		 if (arrivalQueueCapacity == 0) {
		 	fmt.Printf("Time %3d : idle\n", time)
		 }

		if(arrivalQueueCapacity > 0 ) {


			if(arrivalQueue[index].selected && ((arrivalQueue[index].selectionTime + arrivalQueue[index].burstTime) == time)) {

				arrivalQueue[index].completed = true
				arrivalQueue[index].selected = false
				arrivalQueueCapacity--;

				fmt.Printf("Time %3d : %s finished\n", time, process[index].name)


				if(index < (processCount-1)) {
					index++
				}
				
			}

			if(!arrivalQueue[index].selected && !arrivalQueue[index].completed && arrivalQueueCapacity > 0) {

				arrivalQueue[index].selected = true
				arrivalQueue[index].selectionTime = time

				fmt.Printf("Time %3d : %s selected (burst %3d)\n", time, process[index].name, process[index].burstTime)

			} else if (arrivalQueueCapacity == 0) {
				fmt.Printf("Time %3d : idle\n", time)
			}


		} 

		time++
	}

 	//print how long was the system supposed to run for
 	fmt.Printf("Finished at time  %d\n\n", usefor)

 	//sort the slice with respect to process ID
 	sort.Slice(process, func (i, j int) bool { return process[i].ID < process[j].ID })

	//print the wait and turn around times
	for i:=0; i<processCount; i++ {
		fmt.Printf("%s wait   %3d turnaround   %3d\n", process[i].name, process[i].waitTime, process[i].turnAroundTime)
	}
	
	
}

func sjf (process []ProcessInfo, processCount int, usefor int)  {
	
	//sort the slice with respect to first arrival and then burst time
	//sort.Slice(process, func(i, j int) bool { return process[i].burstTime < process[j].burstTime })
}

func rr (process []ProcessInfo, processCount int, usefor int, q int)  {
	fmt.Println("RR = ")
	fmt.Println(process)
}

func main() {

	//Read the file name from the CLI arguements and convert it into an array of bytes
	inputFile := os.Args[1];

	processCount, totalTime, schedulingAlgorithm, quantum = getProcessInfo(inputFile)
	process = getListOfProcesses(inputFile, processCount)

	if(schedulingAlgorithm == "fcfs") {
		fcfs(process, processCount, totalTime)
	} else if (schedulingAlgorithm == "sjf") {
		sjf(process, processCount, totalTime)
	} else if (schedulingAlgorithm == "rr") {
		rr(process, processCount, totalTime, quantum)
	}

	
	
    // fmt.Printf("\nProcess Count = %d\n", processCount)
    // fmt.Printf("Time to run the algorithm for = %d\n", totalTime)
    // fmt.Printf("Type of Scheduling Algorithm = %s\n", schedulingAlgorithm)
    // fmt.Printf("Quantum (if any) = %d\n\n", quantum)
    // fmt.Println("List of Processes: ")
    // fmt.Println(process)
    // fmt.Println()
   
 }













