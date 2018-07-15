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
)

//This struct holds a Process's Information such as Process ID, name, arrival time, burst time, wait time, turn arount time, completion time, selection time, boolean flag for when a process is selected and completed 
type ProcessInfo struct {
	ID int
   	name string

   	arrivalTime int
   	burstTime int
   	waitTime int
   	turnAroundTime int
   	completionTime int
   	selectionTime int

   	selected bool
   	completed bool
  
}

// type BurstTimesSFJ struct {
// 	ID int
// 	burstTime int
// }

//list of variables used in the program
//processCount = total number of processes in the input file
//totalTime = total time for which the algorithm is supposed to run for
//quantum (only for round robin algorithm)
//scheduling alogirhtm = what kind of scheduling algorithm to be used
//process = a slice of ProcessInfo struct
var (
	processCount, totalTime, quantum int
	schedulingAlgorithm string

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

func sort (process []ProcessInfo, typeOfSort string) (p []ProcessInfo) {

	p = make([]ProcessInfo, len(process))
	p = process

	/*
		 Sort processes with respect to following:
		 AT = Arrival Time
		 BT = Burst Time
		 ID = Process ID
	*/
	if(typeOfSort == "AT") {

		for i := 0; i < (len(p) - 1); i++ {

			minIndex := i

			for j := i+1; j < len(p); j++ {

				if(p[j].arrivalTime < p[minIndex].arrivalTime) {
					minIndex = j;
				}
			}

			//swap
			temp := p[minIndex]
			p[minIndex] = p[i]
			p[i] = temp
		}

	} else if (typeOfSort == "BT") {

		for i := 0; i < (len(p) - 1); i++ {

			minIndex := i

			for j := i+1; j < len(p); j++ {

				if(p[j].burstTime < p[minIndex].burstTime) {
					minIndex = j;
				}
			}

			//swap
			temp := p[minIndex]
			p[minIndex] = p[i]
			p[i] = temp
		}

	} else if (typeOfSort == "ID") {

		for i := 0; i < (len(p) - 1); i++ {

			minIndex := i

			for j := i+1; j < len(p); j++ {

				if(p[j].ID < p[minIndex].ID) {
					minIndex = j;
				}
			}

			//swap
			temp := p[minIndex]
			p[minIndex] = p[i]
			p[i] = temp
		}
	}

	return p
}

func fcfs (process []ProcessInfo, processCount int, usefor int, outputFile string)  {
	
	//output stream to print result to the output file
	output, _ := os.Create(outputFile)

	//sort processes with respect to arrival times
	process = sort(process,"AT")

	fmt.Fprintf(output, "%3d processes\n", processCount)
	fmt.Fprintf(output, "Using First-Come First-Served\n")

	var arrivalQueue []ProcessInfo
 	arrivalQueue = make([]ProcessInfo, processCount)
 	arrivalQueueCapacity := 0
		
	time := 0
	index := 0
	for time < usefor {

		//if time = arrival time of a process, then add it to arrival queue and increase queue capacity
		for i:=0; i < processCount; i++ {
			if(process[i].arrivalTime == time) {

				fmt.Fprintf(output, "Time %3d : %s arrived\n", time, process[i].name)
				arrivalQueue[i] = process[i]
				arrivalQueueCapacity++
			}
		}

		//if there is nothing in the arrival queue, then the CPU is idle
		 if (arrivalQueueCapacity == 0) {
		 	fmt.Fprintf(output, "Time %3d : Idle\n", time)
		 }

		 //if there is something in the arrival queue, then select a process and run it
		if(arrivalQueueCapacity > 0 ) {


			//if the selelcted process is completed, then flag it as completed and decerease queue capacity
			if(arrivalQueue[index].selected && ((arrivalQueue[index].selectionTime + arrivalQueue[index].burstTime) == time)) {

				arrivalQueue[index].completed = true
				arrivalQueue[index].completionTime = time

				arrivalQueue[index].selected = false
				arrivalQueueCapacity--;

				fmt.Fprintf(output, "Time %3d : %s finished\n", time, process[index].name)

				//increment the index to move on to the next process in the arrival queue 
				if(index < (processCount-1)) {
					index++
				}
				
			}

			//if none of the processes is selected in the queue then select one and mark it as selected
			if(!arrivalQueue[index].selected && !arrivalQueue[index].completed && arrivalQueueCapacity > 0) {

				arrivalQueue[index].selected = true
				arrivalQueue[index].selectionTime = time

				fmt.Fprintf(output, "Time %3d : %s selected (burst %3d)\n", time, process[index].name, process[index].burstTime)

			} else if (arrivalQueueCapacity == 0) {
				fmt.Fprintf(output, "Time %3d : Idle\n", time)
			}

		} 

		time++
	}

	//set the arrival queue to process to update the wait and turn around times
	process = arrivalQueue

	for i:=0; i<processCount; i++ {

		process[i].turnAroundTime = process[i].completionTime - process[i].arrivalTime
		process[i].waitTime = process[i].turnAroundTime - process[i].burstTime
		
	}

 	//print how long was the system supposed to run for
 	fmt.Fprintf(output, "Finished at time  %d\n\n", usefor)

 	//sort processes with respect to process IDs
 	process = sort(process,"ID")

	//print the wait and turn around times
	for i:=0; i<processCount; i++ {
		fmt.Fprintf(output, "%s wait %3d turnaround %3d\n", process[i].name, process[i].waitTime, process[i].turnAroundTime)
	}
		
}

func sjf (process []ProcessInfo, processCount int, usefor int, outputFile string)  {

	//output stream to print result to the output file
	output, _ := os.Create(outputFile)
	
	//sort processes with respect to arrival times
	process = sort(process,"AT")

	fmt.Fprintf(output, "%3d processes\n", processCount)
	fmt.Fprintf(output, "Using preemptive Shortest Job First\n")

	//create an arrival queue with capacity of processCount and length 0
 	var arrivalQueue []ProcessInfo = make([]ProcessInfo, 0, processCount)
 	
	time := 0
	arrivalQueueCapacity := 0
	
	for time < usefor {	

		//add processes to arrival queue as they arrive
		for i:=0; i < processCount; i++ {
			if(process[i].arrivalTime == time) {

				//print the processes that have arrived along with at what time did they arrive
				fmt.Fprintf(output, "Time %3d : %s arrived\n", time, process[i].name)

				//add the process into arrival queue
				arrivalQueue = append(arrivalQueue, process[i])
				arrivalQueueCapacity++
				 
			}
		}

		if (arrivalQueueCapacity == 0){
			fmt.Fprintf(output, "Time %3d : Idle\n", time)
		}

		index := 0
		previousProcessID := 0

		if(arrivalQueueCapacity > 0){

			for (arrivalQueue[index].completed) {
				index++
			}

			if((index < processCount) &&arrivalQueue[index].selected) {
				previousProcessID = arrivalQueue[index].ID
			}

			for i:=0; i < len(arrivalQueue); i++ {
				arrivalQueue[i].selected = false
			}

			arrivalQueue = sort(arrivalQueue,"BT")

			for ((index < processCount) && arrivalQueue[index].completed) {
				index++
			}

			if((index < processCount) && (arrivalQueue[index].ID == previousProcessID)) {
				arrivalQueue[index].selected = true
			}

			//if the selelcted process is completed, then flag it as completed and decerease queue capacity
			if((index < processCount) && arrivalQueue[index].selected && (arrivalQueue[index].burstTime == 0) && !arrivalQueue[index].completed) {

				//unselect the process, mark it completed and note its completion time
				arrivalQueue[index].selected = false
				arrivalQueue[index].completed = true
				arrivalQueue[index].completionTime = time
				arrivalQueueCapacity--;

				fmt.Fprintf(output, "Time %3d : %s finished\n", time, arrivalQueue[index].name)
				
			}


			if ((index < processCount) && arrivalQueue[index].completed && arrivalQueueCapacity > 0) {
				index++
			}

			
			//if none of the processes is selected in the queue then select one and mark it as selected
			if((index < processCount) && !arrivalQueue[index].selected && !arrivalQueue[index].completed && arrivalQueueCapacity > 0) {

				arrivalQueue[index].selected = true
				arrivalQueue[index].selectionTime = time

				fmt.Fprintf(output, "Time %3d : %s selected (burst %3d)\n", time, arrivalQueue[index].name, arrivalQueue[index].burstTime)

			} 

			if (arrivalQueueCapacity == 0) {
				fmt.Fprintf(output, "Time %3d : Idle\n", time)
			}

		}  

		 time++


		if((index < processCount) && arrivalQueue[index].burstTime > 0){
			arrivalQueue[index].burstTime--
		}

	}

	//calculate the turn around time of the processes
	for i:=0; i<processCount; i++ {
		arrivalQueue[i].turnAroundTime = arrivalQueue[i].completionTime - arrivalQueue[i].arrivalTime
	}

	//sort processes with respect to process IDs
 	arrivalQueue = sort(arrivalQueue,"ID")
 	process = sort(process,"ID")

 	//calculate the wait time of the processes
	for i:=0; i<processCount; i++ {
		arrivalQueue[i].waitTime = arrivalQueue[i].turnAroundTime - process[i].burstTime
	}

 	//print how long was the system supposed to run for
 	fmt.Fprintf(output, "Finished at time  %d\n\n", usefor)

	//print the wait and turn around times
	for i:=0; i<processCount; i++ {
		fmt.Fprintf(output,"%s wait %3d turnaround %3d\n", arrivalQueue[i].name, arrivalQueue[i].waitTime, arrivalQueue[i].turnAroundTime)
	}

}

func rr (process []ProcessInfo, processCount int, usefor int, q int, outputFile string)  {
	
	//output stream to print result to the output file
	output, _ := os.Create(outputFile)

	process = sort(process,"AT")

	fmt.Fprintf(output,"%3d processes\n", processCount)
	fmt.Fprintf(output,"Using Round-Robin\n")
	fmt.Fprintf(output,"Quantum %3d\n\n", q)

	//create an arrival queue with capacity of processCount and length 0
 	var arrivalQueue []ProcessInfo = make([]ProcessInfo, 0, processCount)
 	var completedQueue []ProcessInfo = make([]ProcessInfo, 0, processCount)
 	var scheduledProcess ProcessInfo
 	
 	
	time := 0
	arrivalQueueCapacity := 0
	index := 0

	for time < usefor {	

		quantum := 0

		if((index < processCount) && (process[index].arrivalTime == time)) {

			fmt.Fprintf(output,"Time %3d : %s arrived\n", time, process[index].name)
			arrivalQueue = append(arrivalQueue, process[index])
			arrivalQueueCapacity++

			//increment the pointer
			index++
		}
		

		if (arrivalQueueCapacity == 0){
			fmt.Fprintf(output, "Time %3d : Idle\n", time)
			time++
		}

		if(arrivalQueueCapacity > 0){

			//if arrival queue has elements
			// 1. Put it in scheduled process var and mark it selected to show it is running
			// 2. Remove it from arrival queue and decrease queue capacity
			// 3. Print the selected process
			// 4. Run it through quantum:- decrement burst time, increment quantum, and increment time. 
			//	  At. each time increment check if a process has arrived
			// 5. If a process has a arrived then add it to the arrival queue
			// 6. Add the selected process back to arrival queue at the end and run the process and unselect it

			scheduledProcess = arrivalQueue[0]
			scheduledProcess.selected = true

			//remove first element from arrival queue
			arrivalQueueCapacity--
			if(arrivalQueueCapacity > 0) {
				arrivalQueue = arrivalQueue[1:]
			} else {
				arrivalQueue = nil
			}

			if(scheduledProcess.selected && !scheduledProcess.completed) {

				fmt.Fprintf(output, "Time %3d : %s selected (burst %3d)\n", time, scheduledProcess.name, scheduledProcess.burstTime)

				for quantum < q {

					// fmt.Printf("Goes through quantum %d times\n", quantum)

					scheduledProcess.burstTime--
					quantum++
					time++

					if((index < processCount) && (process[index].arrivalTime == time)) {

						fmt.Fprintf(output, "Time %3d : %s arrived\n", time, process[index].name)
						arrivalQueue = append(arrivalQueue, process[index])
						arrivalQueueCapacity++

						//increment the pointer
						index++
					}

					if(scheduledProcess.burstTime == 0) {

						scheduledProcess.completed = true
						scheduledProcess.selected = false
						scheduledProcess.completionTime = time
						arrivalQueueCapacity--
						fmt.Fprintf(output, "Time %3d : %s finished\n", time, scheduledProcess.name)

						completedQueue = append(completedQueue, scheduledProcess)

						break
					}

					
				}

				//add the process to back of arrival queue after running it upto quantum
				scheduledProcess.selected = false
				arrivalQueue = append(arrivalQueue, scheduledProcess)
				arrivalQueueCapacity++
			}

			if (arrivalQueueCapacity == 0){
				fmt.Fprintf(output, "Time %3d : Idle\n", time)
				time++
			}

		}	

	}


 	//print how long was the system supposed to run for
 	fmt.Fprintf(output, "Finished at time  %d\n\n", usefor)

 	//sort processes with respect to process IDs
 	completedQueue = sort(completedQueue,"ID")
 	process = sort(process,"ID")

 	//calculate the turn around times and wait times of the processes
	for i:=0; i<processCount; i++ {
		completedQueue[i].turnAroundTime = completedQueue[i].completionTime - completedQueue[i].arrivalTime
		completedQueue[i].waitTime = completedQueue[i].turnAroundTime - process[i].burstTime
	}

 	//print the wait and turn around times
	for i:=0; i<processCount; i++ {
		fmt.Fprintf(output, "%s wait %3d turnaround %3d\n", completedQueue[i].name, completedQueue[i].waitTime, completedQueue[i].turnAroundTime)
	}

	
}

func main() {

	//Read the file name from the CLI arguements and convert it into an array of bytes
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	processCount, totalTime, schedulingAlgorithm, quantum = getProcessInfo(inputFile)
	process = getListOfProcesses(inputFile, processCount)

	if(schedulingAlgorithm == "fcfs") {
		fcfs(process, processCount, totalTime, outputFile)
	} else if (schedulingAlgorithm == "sjf") {
		sjf(process, processCount, totalTime, outputFile)
	} else if (schedulingAlgorithm == "rr") {
		rr(process, processCount, totalTime, quantum, outputFile)
	}

	
   
 }
