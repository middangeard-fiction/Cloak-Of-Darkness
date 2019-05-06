package main

import (
	"fmt"

	mid "github.com/middangeard-fiction/middangeard"
)

var game mid.Game

func init() {
	game = mid.Game{
		Title:  "Cloak of Darkness",
		Author: "Roger Firth (implemented by Mark Bauermeister)",
		Intro: `Hurrying through the rainswept November night, you're glad to see the
		bright lights of the Opera House.
		
		It's surprising that there aren't more
		people about but, hey, what do you expect in a cheap demo game...?`,
		Goodbye:  "See you later!",
		MaxScore: 2,
		Verbs: map[string]func(...string){
			"drop": drop,
		},
		Synonyms: map[string][]string{
			"drop":  {"throw"},
			"north": {"south"},
		},
	}

	game.Player = mid.Player{
		Name:        "Self",
		Description: "Just an average individual.",
		Location:    "foyer",
		Score:       0,
	}

	game.Rooms = map[string]*mid.Room{
		"foyer": {
			Name: "Foyer of the Opera House",
			Description: `You are standing in a spacious hall, splendidly decorated in red
            and gold, with glittering chandeliers overhead. The entrance from
			the street is to the north, and there are doorways south and west.`,
			Lit: true,
			OnEnter: func(r *mid.Room) {
				if r.Visited {
					r.Description = "Nothing has changed since you last entered here."
				}
			},
			Directions: mid.Directions{
				West:  "cloakroom",
				South: "bar",
			},
		},
		"cloakroom": {
			Name: "Cloakroom",
			Description: `The walls of this small room were clearly once lined with hooks, though now only one remains.
			The exit is a door to the east.`,
			Lit: true,
			Directions: mid.Directions{
				East: "foyer",
			},
		},
		"bar": {
			Name: "Foyer Bar",
			Description: `The bar, much rougher than you'd have guessed
			after the opulence of the foyer to the north, is completely empty.
			There seems to be some sort of message scrawled in the sawdust on the floor.`,
			Directions: mid.Directions{
				North: "foyer",
			},
		},
	}
}

// b, err := json.Marshal(g)
// if err != nil {
// 	fmt.Println(err)
// 	return
// }
// fmt.Println(string(b))

func drop(args ...string) {
	fmt.Println("Dropped something")
}
