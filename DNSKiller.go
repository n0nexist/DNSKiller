/*
DNSKiller - github.com/n0nexist
DNS top and sub level domain bruteforcer
*/

package main

import (
    "fmt"
    "net"
	"os"
	"bufio"
	"strconv"
	"sync"
)

var (
	doesnotexist = ""
	currentloadingstatus = "-"
	currentdomain = ""
	currentcolor = ""
	currentfilename = ""
	alreadyfound = []string {}

	purple = "\033[0;35m"
	boldpurple = "\033[1;35m"

	cyan = "\033[0;36m"
	boldcyan = "\033[1;36m"
	
	blue = "\033[0;34m"
	boldblue = "\033[1;34m"

	reset = "\033[0;0m"
)

func checkInList(elemento string) bool {
	/*
	checks if an element is present
	on a list of strings
	*/
    for _, item := range alreadyfound {
        if item == elemento {
            return true
        }
    }
    return false
}

func tryDomainName(domain string) {
	/*
	checks if the given domain is valid
	*/
    ips, err := net.LookupIP(domain)
    if err == nil {
		currentcolor = cyan
		if checkInList(ips[0].String()) {
			currentcolor = boldblue+"(already found) "
		}else{
			alreadyfound = append(alreadyfound, ips[0].String())
		}
		if ips[0].String() != doesnotexist {
			currentstr := fmt.Sprintf("%s>%s Found %s%s%s -> %s%s%s%s\n",purple,reset,cyan,domain,reset,currentcolor,cyan,ips[0].String(),reset)
			fmt.Print(currentstr)
			appendToFile(fmt.Sprintf("> Found %s -> %s",domain,ips[0].String()))
		}
	}
}

func getInvalidDomain(domain string) string {
	/*
	gets the response of an invalid domain name
	that will be confronted with next responses
	*/
    ips, err := net.LookupIP(domain)
    if err != nil {
        return "*no*"
    }
    return ips[0].String()
}

func appendToFile(content string){
	/*
	writes to the file path
	indicated by the global variable currentfilename
	(in append mode) and writes the content inside of the content variable
	*/
	if currentfilename != "" {
		f, err := os.OpenFile(currentfilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {}
		defer f.Close()
		if _, err = f.WriteString(content+"\n"); err != nil {}
	}
}

func slideloading(){
	/*
	changes the loading animation
	*/
	switch currentloadingstatus {
		case "-":
			currentloadingstatus = "/"
			break;
		case "/":
			currentloadingstatus = "\\"
			break;
		case "\\":
			currentloadingstatus = "-"
			break;
	}
}

func logo(){
	/*
	prints the logo
	*/
	fmt.Printf(`%s	.__ .  . __..  . ..      
	|  \|\ |(__ |_/ *|| _ ._.
%s	|__/| \|.__)|  \|||(/,[  							
	%s%s`,purple,boldpurple,reset,"\n")
}

func showcredits(){
	/*
	prints the credits and quit
	*/
	fmt.Printf("%s>%s DNSKiller was developed by %sn0nexist %s(%shttps://www.n0nexist.github.io%s)%s\n\n",purple,reset,boldpurple,boldcyan,blue,boldcyan,reset)
	os.Exit(0)
}

func showhelp(){
	/*
	shows help menu
	*/
	fmt.Printf(`
%sUsage%s: %s%s %s(%sdomain%s) (%ssubdomain-wordlist-path%s) (%stopleveldomain-wordlist-path%s) (%sthreads%s) (%soutput file [OPTIONAL]%s)
%sExample%s: %s%s %sgoogle wordlists/sub.txt wordlist/top.txt 200

`,boldpurple,purple,blue,os.Args[0],boldblue,cyan,boldblue,cyan,boldblue,cyan,boldblue,cyan,boldblue,cyan,boldblue,boldpurple,purple,blue,os.Args[0],cyan)
	showcredits()
}

func getlines(filename string) int {
	/*
	gets the number of lines
	in a file
	*/
	count := 0
	readFile, err := os.Open(filename)
	if err != nil{
		showhelp()
	}
	fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    for fileScanner.Scan() {
        count += 1
    }
  
    readFile.Close()
	return count
}

func worker(mycurrentline chan string, wg *sync.WaitGroup) {
	/*
	a single task
	*/
	defer wg.Done()
	for line := range mycurrentline {
		tryDomainName(line)
	}
}

func main() {
	/*
	DNSKiller starts here
	*/

    doesnotexist = getInvalidDomain("domain_that_does_not_exist.blablabla")
    logo()
	if len(os.Args) < 5 {
		showhelp()
	}
	target := os.Args[1]
	subpath := os.Args[2]
	toppath := os.Args[3]
	tempthreadcount := os.Args[4]
	threadcount, convertingerror := strconv.Atoi(tempthreadcount)
	if convertingerror != nil {
		showhelp()
	}
	if len(os.Args) >= 6 {
		currentfilename = os.Args[5]
		appendToFile("[ github.com/n0nexist/DNSKiller ]")
		appendToFile(fmt.Sprintf("[ target = %s ]",target))
	}
	totaltries := getlines(subpath)*getlines(toppath)
	fmt.Printf("%s>%s Starting to resolve possible valid domains for %s'%s%s%s'\n",purple,reset,boldpurple,blue,target,boldpurple)
	fmt.Printf("%s>%s Calculated that the total possibilities are %d\n\n",purple,reset,totaltries)

	subdomain_file, err := os.Open(subpath)
	if err != nil {
		showhelp()
	}

	topdomain_file, errtwo := os.Open(toppath)
	if errtwo != nil {
		showhelp()
	}

	subdomain_file_scanner := bufio.NewScanner(subdomain_file)
	subdomain_file_scanner.Split(bufio.ScanLines)

	var wg sync.WaitGroup
	mycurrentline := make(chan string)
	for i := 0; i < threadcount; i++ {
		wg.Add(1)
		go worker(mycurrentline, &wg)
	}

	for subdomain_file_scanner.Scan() {
		topdomain_file, errtwo := os.Open(toppath)
		if errtwo != nil {
			showhelp()
		}
		topdomain_file_scanner := bufio.NewScanner(topdomain_file)
		topdomain_file_scanner.Split(bufio.ScanLines)

		for topdomain_file_scanner.Scan() {
			fmt.Printf("%s[%s%s%s]%s Trying %s.%s.%s          \r",boldpurple,blue,currentloadingstatus,boldpurple,reset,subdomain_file_scanner.Text(), target, topdomain_file_scanner.Text())
			slideloading()

			currentdomain = fmt.Sprintf("%s.%s.%s", subdomain_file_scanner.Text(), target, topdomain_file_scanner.Text())

			mycurrentline <- currentdomain
		}
		topdomain_file.Close()
	}

	close(mycurrentline)
	wg.Wait()

	subdomain_file.Close()
	topdomain_file.Close()	

}