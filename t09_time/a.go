package main

import ("fmt"
	"time"
)

func main(){
	test5()
	
//	test4()

//	test3()
	
//	test2()
}

func test5(){
currentTime := time.Now()
 
    fmt.Println("Current Time in String: ", currentTime.String())
     
    fmt.Println("MM-DD-YYYY : ", currentTime.Format("01-02-2006"))
     
    fmt.Println("YYYY-MM-DD : ", currentTime.Format("2006-01-02"))
     
    fmt.Println("YYYY.MM.DD : ", currentTime.Format("2006.01.02 15:04:05"))
     
    fmt.Println("YYYY#MM#DD {Special Character} : ", currentTime.Format("2006#01#02"))
     
    fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))
         
    fmt.Println("Time with MicroSeconds: ", currentTime.Format("2006-01-02 15:04:05.000000"))
     
    fmt.Println("Time with NanoSeconds: ", currentTime.Format("2006-01-02 15:04:05.000000000"))
     
    fmt.Println("ShortNum Month : ", currentTime.Format("2006-1-02"))
     
    fmt.Println("LongMonth : ", currentTime.Format("2006-January-02"))
     
    fmt.Println("ShortMonth : ", currentTime.Format("2006-Jan-02"))
     
    fmt.Println("ShortYear : ", currentTime.Format("06-Jan-02"))
     
    fmt.Println("LongWeekDay : ", currentTime.Format("2006-01-02 15:04:05 Monday"))
     
    fmt.Println("ShortWeek Day : ", currentTime.Format("2006-01-02 Mon"))   
     
    fmt.Println("ShortDay : ", currentTime.Format("Mon 2006-01-2"))
     
    fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5"))   
     
    fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5 PM"))    
     
    fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5 pm"))    
}

func test4(){
now := time.Now()
    fmt.Println("\nToday:", now)
     
    after := now.AddDate(1, 0, 0)
    fmt.Println("\nAdd 1 Year:", after)
     
    after = now.AddDate(0, 1, 0)
    fmt.Println("\nAdd 1 Month:", after)
     
    after = now.AddDate(0, 0, 1)
    fmt.Println("\nAdd 1 Day:", after)
     
    after = now.AddDate(2, 2, 5)
    fmt.Println("\nAdd multiple values:", after)
     
    after = now.Add(10*time.Minute)
    fmt.Println("\nAdd 10 Minutes:", after)
     
    after = now.Add(10*time.Second)
    fmt.Println("\nAdd 10 Second:", after)
     
    after = now.Add(10*time.Hour)
    fmt.Println("\nAdd 10 Hour:", after)
     
    after = now.Add(10*time.Millisecond)
    fmt.Println("\nAdd 10 Millisecond:", after)
     
    after = now.Add(10*time.Microsecond)
    fmt.Println("\nAdd 10 Microsecond:", after)
     
    after = now.Add(10*time.Nanosecond)
    fmt.Println("\nAdd 10 Nanosecond:", after)

}

func test3(){
loc, _ := time.LoadLocation("UTC")
now := time.Now().In(loc)
fmt.Println("\nToday : ", loc, " Time : ", now)
 
pastDate := time.Date(2015, time.May, 21, 23, 10, 52, 211, time.UTC)
fmt.Println("\n\nPast  : ", loc, " Time : ", pastDate) // 
fmt.Printf("###############################################################\n")
diff := now.Sub(pastDate)
 
hrs := int(diff.Hours())
fmt.Printf("Diffrence in Hours : %d Hours\n", hrs)
 
mins := int(diff.Minutes())
fmt.Printf("Diffrence in Minutes : %d Minutes\n", mins)
 
second := int(diff.Seconds())
fmt.Printf("Diffrence in Seconds : %d Seconds\n", second)
 
days := int(diff.Hours() / 24)
fmt.Printf("Diffrence in days : %d days\n", days)
 
fmt.Printf("###############################################################\n\n\n")
 
futureDate := time.Date(2019, time.May, 21, 23, 10, 52, 211, time.UTC)
fmt.Println("Future  : ", loc, " Time : ", futureDate) // 
fmt.Printf("###############################################################\n")
diff = futureDate.Sub(now)
 
hrs = int(diff.Hours())
fmt.Printf("Diffrence in Hours : %d Hours\n", hrs)
 
mins = int(diff.Minutes())
fmt.Printf("Diffrence in Minutes : %d Minutes\n", mins)
 
second = int(diff.Seconds())
fmt.Printf("Diffrence in Seconds : %d Seconds\n", second)
 
days = int(diff.Hours() / 24)
fmt.Printf("Diffrence in days : %d days\n", days)
 

}

func test2(){
 t := time.Now() 
    z, _ := t.Zone()
    fmt.Println("ZONE : ", z, " Time : ", t) // local time
      
    location, err := time.LoadLocation("EST")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("ZONE : ", location, " Time : ", t.In(location)) // EST
  
    loc, _ := time.LoadLocation("UTC")
    now := time.Now().In(loc)
    fmt.Println("ZONE : ", loc, " Time : ", now) // UTC 
     
    loc, _ = time.LoadLocation("MST")
    now = time.Now().In(loc)
    fmt.Println("ZONE : ", loc, " Time : ", now) // MST
}

func test1(){
	
	t := time.Now()
	//local time	
	fmt.Println("Location : ", t.Location(), " Time : ", t)
	
	location, err := time.LoadLocation("America/New_York")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Location : ", location, " Time : ", t.In(location)) // America/New_York
 
    loc, _ := time.LoadLocation("Asia/Shanghai")
    now := time.Now().In(loc)
    fmt.Println("Location : ", loc, " Time : ", now) // Asia/Shanghai
	

}
