package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

const (
	PathToCards  = "./data/cards.json"
	CardsPerDeck = 75
	RandSeed     = 34384239482
)

// no need to generate identities within the pool. totally optional though.
var ExcludeTypeCode = [...]string{"identity"}

// removing special /a lternative art cards
var ExcludeSetCode = [...]string{"special", "alt"}

// removing 6 since lunar cycle isn't out yet
var ExcludeCycleNumber = [...]int{6}

type Card struct {
	LastModified    string `json:"last-modified"`
	Code            string `json:"code"`
	Title           string `json:"title"`
	Type            string `json:"type"`
	TypeCode        string `json:"type_code"`
	Subtype         string `json:"subtype"`
	SubtypeCode     string `json:"subtype_code"`
	Text            string `json:"text"`
	BaseLink        int    `json:"baselink,omitempty"`
	Faction         string `json:"faction"`
	FactionCode     string `json:"faction_code"`
	FactionLetter   string `json:"faction_letter"`
	Flavor          string `json:"flavor"`
	Illustrator     string `json:"illustrator"`
	InfluenceLimit  int    `json:"influencelimit,omitempty"`
	MinimumDeckSize int    `json:"minimumdecksize,omitempty"`
	Number          int    `json:"number"`
	Quantity        int    `json:'quantity"`
	SetName         string `json:"setname"`
	SetCode         string `json:"set_code"`
	Side            string `json:"side"`
	SideCode        string `json:"side_code"`
	Uniqueness      bool   `json:"uniqueness"`
	CycleNumber     int    `json:"cyclenumber"`
	Url             string `json:"url"`
	ImageSrc        string `json:"imagesrc"`
	LargeImageSrc   string `json:"largeimagesrc,omitempty"`
}

func ExcludeCard(card Card) (result bool) {

	for _, value := range ExcludeTypeCode {
		if card.TypeCode == value {
			return true
		}
	}

	for _, value := range ExcludeSetCode {
		if card.SetCode == value {
			return true
		}
	}

	for _, value := range ExcludeCycleNumber {
		if card.CycleNumber == value {
			return true
		}
	}
	return false
}

func ProcessFile(file []byte) (corp []Card, runner []Card) {
	// difference between make and new.
	raw := make([]json.RawMessage, 10)
	if err := json.Unmarshal(file, &raw); err != nil {
		log.Fatalf("error %v \n", err)
		os.Exit(1)
	}

	corp = make([]Card, 0)
	runner = make([]Card, 0)

	for i := 0; i < len(raw); i++ {
		card := Card{}
		if err := json.Unmarshal(raw[i], &card); err != nil {
			log.Fatalf("error %v \n", err)
			os.Exit(1)
		}

		// fmt.Printf("Card: %#v\n", card)

		// generate corp / runner lists
		if card.Side == "Corp" {
			if !ExcludeCard(card) {
				corp = append(corp, card)
			}
		}

		if card.Side == "Runner" {
			if !ExcludeCard(card) {
				runner = append(runner, card)
			}
		}
	}

	fmt.Printf("Number of corp cards: %d \n", len(corp))
	fmt.Printf("Number of runner cards: %d \n", len(runner))
	return corp, runner
}

// GeneratePool pseudo-randomly generates a new pool of cards of size
// cardPoolSize
func GeneratePool(cardPoolSize int, cards []Card) (pool map[string]int) {
	originalPoolSize := len(cards)

	pool = make(map[string]int)

	for i := 0; i < cardPoolSize; i++ {
		index := rand.Intn(originalPoolSize)
		card := cards[index]
		cardTitle := card.Title
		numItems := pool[cardTitle]

		pool[cardTitle] = numItems + 1
	}

	return pool
}

func GenerateText(cards map[string]int, isCorp bool, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("File error: %v \n", err)
		os.Exit(1)
	}

	defer f.Close()

	if isCorp {
		f.WriteString("The Shadow: Pulling the Strings\n")
	} else {
		f.WriteString("The Masque: Cyber General\n")
	}

	for title, value := range cards {
		cardLine := fmt.Sprintf("%s x%d\n", title, value)
		f.WriteString(cardLine)
	}
}

func main() {
	rand.Seed(RandSeed)

	file, err := ioutil.ReadFile(PathToCards)
	if err != nil {
		log.Fatalf("File error: %v \n", err)
		os.Exit(1)
	}

	corp, runner := ProcessFile(file)

	corpDeck := GeneratePool(CardsPerDeck, corp)
	runnerDeck := GeneratePool(CardsPerDeck, runner)

	GenerateText(corpDeck, true,
		fmt.Sprintf("pools/corp-%d.txt", RandSeed))

	GenerateText(runnerDeck, false,
		fmt.Sprintf("pools/runner-%d.txt", RandSeed))

	fmt.Printf("corp deck: %#v \n", corpDeck)
	fmt.Printf("runner deck: %#v \n", runnerDeck)

}
