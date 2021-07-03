package main

import (
	
		"fmt"
		"log"
		"encoding/csv"
		"os"
		"strconv"

		"github.com/gocolly/colly"

 )
func main(){
	 fname:="data.csv"
	 file,err:= os.Create(fname)
	 if err!=nil{
		 log.Fatal("fuck u ,err :%q",err)
		 return
	 }
	 defer file.Close()

	 writer:=csv.NewWriter(file)
	 defer writer.Flush()

	 c:= colly.NewCollector(
		 colly.AllowedDomains("internshala.com"),
	 )
	 c.OnHTML(".internship_meta", func(e *colly.HTMLElement){
		 writer.Write([]string{
			 e.ChildText("a"),
			 e.ChildText("span"),

		 })
		
	 })
	 for i:=0;i<312;i++{
		 fmt.Printf("scraping page: %d\n",i)
		c.Visit("https://internshala.com/internships/page-" + strconv.Itoa(i))	 
	}
	log.Printf("scraping complete \n")
	log.Println(c)
}