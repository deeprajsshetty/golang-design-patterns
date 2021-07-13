package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)


var catelogCount = 0

/*
	Following SRP
	Catelog Type holds responsibility of Catelog Functionality
*/
type Catelog struct {
	catelogEntries []string
}

func (c *Catelog) String() string {
	return strings.Join(c.catelogEntries, "\n")
}

// Add Catelog
func (c *Catelog) AddCatelogEntry(catelog string) int {
	catelogCount++
	catelogEntry := fmt.Sprintf("%d: %s", catelogCount, catelog)
	c.catelogEntries = append(c.catelogEntries, catelogEntry)
	return catelogCount
}

// Remove Catelog
func (c *Catelog) RemoveCatelogEntry(catelog string) int {
	for i := 0; i < len(c.catelogEntries); i++ {
		if c.catelogEntries[i] == catelog {
			c.catelogEntries = append(c.catelogEntries[:i], c.catelogEntries[i+1:]...)
		}
	}
	catelogCount--
	return catelogCount
}


/*
	Breaking SRP
	As we included other functionality like Saving to File and Persistent related logics.
	Which can be handled in different Package as Generic Layer. Can be Infrastructure.
*/
// This Logic should be part of Seperate Type or Can be seperate Package. Preferred seperate package.
func (c *Catelog) Save(filename string) {
	_=ioutil.WriteFile(filename, []byte(c.String()), 0644)
}

func main() {
	// Following SRP
	catelogs := Catelog{}
	catelogs.AddCatelogEntry("ITEM 001")
	catelogs.AddCatelogEntry("ITEM 002")
	catelogs.AddCatelogEntry("ITEM 003")
	catelogs.AddCatelogEntry("ITEM 004")
	catelogs.AddCatelogEntry("ITEM 005")
	fmt.Println(catelogs.catelogEntries)
	catelogs.RemoveCatelogEntry("3: ITEM 003")
	catelogs.RemoveCatelogEntry("5: ITEM 005")
	fmt.Println(catelogs.catelogEntries)
}